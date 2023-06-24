package state

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sjpau/passage/src/graphics"
	"github.com/sjpau/passage/src/help"
	"github.com/sjpau/passage/src/lvl"
	"github.com/sjpau/vector"
	"github.com/tinne26/etxt"
	"github.com/tinne26/fonts/liberation/lbrtserif"
)

type Gameplay struct {
	changeState   bool
	changeStateTo int
	text          *etxt.Renderer
	LvlFS         *embed.FS
	lvls          []*lvl.Level
	assignLayout  bool
	currentLayout [][]int
	currentLvl    int
	currentChar   *vector.Vector2D
}

const (
	charLeft  = "˂"
	charRight = "˃"
	charDown  = "˅"
	charUp    = "˄"
	charExit  = "o"
	charStart = "*"
)

var lvlCharMap = map[int]string{
	0: charStart,
	1: charExit,
	2: charUp,
	3: charDown,
	4: charRight,
	5: charLeft,
}

func (g *Gameplay) Change() int {
	if g.changeState {
		g.changeState = false
		return g.changeStateTo
	} else {
		return -1
	}
}

func (g *Gameplay) Init() {
	renderer := etxt.NewStdRenderer()
	renderer.SetFont(lbrtserif.Font())
	cache := etxt.NewDefaultCache(1024 * 1024)
	renderer.SetCacheHandler(cache.NewHandler())
	renderer.SetColor(graphics.COLOR_RED)
	renderer.SetSizePx(int(fontSize * ebiten.DeviceScaleFactor()))
	renderer.SetAlign(etxt.YCenter, etxt.XCenter)
	g.text = renderer

	g.currentChar = &vector.Vector2D{X: 0, Y: 0}
	g.assignLayout = true

	basePath := "src/lvl/conf/"

	//TODO create slice of lvls to append to
	lvlEscape, e := lvl.LoadFromJSON(g.LvlFS, basePath+"escape.json")
	log.Print(lvlEscape)
	help.Check(e)
	g.lvls = append(g.lvls, lvlEscape)

}

func (g *Gameplay) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.changeStateTo = STATE_MENU
		g.changeState = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.currentChar.X += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.currentChar.X -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.currentChar.Y += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.currentChar.Y -= 1
	}
	rows := len(g.lvls[0].Layout)
	cols := len(g.lvls[0].Layout[0])

	if g.currentChar.Y < 0 {
		g.currentChar.Y = 0
	} else if int(g.currentChar.Y) >= rows {
		g.currentChar.Y = float64(rows - 1)
	}
	if g.currentChar.X < 0 {
		g.currentChar.X = 0
	} else if int(g.currentChar.X) >= cols {
		g.currentChar.X = float64(cols - 1)
	}
}

func (g *Gameplay) Draw(screen *ebiten.Image) {
	screen.Fill(graphics.COLOR_GREY)
	g.text.SetTarget(screen)
	bounds := screen.Bounds()
	x, y := bounds.Dx()/2, bounds.Dy()/2
	if g.assignLayout {
		g.currentLayout = g.lvls[0].Layout
		g.assignLayout = false
	}

	rows := len(g.currentLayout)
	cols := len(g.currentLayout[0])

	//Calculations so that matrix is drawn around the center of it's "middle" element
	startX := x - (cols/2)*horizontalSP
	startY := y - (rows/2)*verticalSP

	for i, row := range g.currentLayout {
		for j, value := range row {
			if i == int(g.currentChar.Y) && j == int(g.currentChar.X) {
				g.text.SetColor(graphics.COLOR_PINK)
			} else {
				g.text.SetColor(graphics.COLOR_RED)
			}
			g.text.Draw(lvlCharMap[value], startX+j*horizontalSP, startY+i*verticalSP)
		}
	}
}
