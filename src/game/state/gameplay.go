package state

import (
	"embed"
	"log"
	"os"

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
	lvlCompleted  bool
}

const (
	charLeft  = "˂"
	charRight = "˃"
	charDown  = "˅"
	charUp    = "˄"
	charExit  = "o"
	charStart = "*"
)

const (
	codeExit = iota
	codeStart
	codeUp
	codeDown
	codeRight
	codeLeft
)

var lvlCharMap = map[int]string{
	codeStart: charStart,
	codeExit:  charExit,
	codeUp:    charUp,
	codeDown:  charDown,
	codeRight: charRight,
	codeLeft:  charLeft,
}

var reverseMap = map[int]int{
	codeStart: codeExit,
	codeExit:  codeStart,
	codeLeft:  codeRight,
	codeRight: codeLeft,
	codeUp:    codeDown,
	codeDown:  codeUp,
}

var passageFinderMap = map[int][]int{
	codeUp:    {codeUp, codeLeft, codeRight},
	codeDown:  {codeDown, codeLeft, codeRight},
	codeLeft:  {codeLeft, codeUp, codeDown},
	codeRight: {codeRight, codeUp, codeDown},
}

var up = vector.Vector2D{X: 0, Y: -1}
var down = vector.Vector2D{X: 0, Y: 1}
var right = vector.Vector2D{X: 1, Y: 0}
var left = vector.Vector2D{X: -1, Y: 0}

var movePassageMap = map[int]vector.Vector2D{
	codeUp:    up,
	codeDown:  down,
	codeLeft:  left,
	codeRight: right,
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

func (g *Gameplay) FindInLayout(value int) [][]int {
	var positions [][]int

	for i, row := range g.lvls[g.currentLvl].Layout {
		for j, v := range row {
			if v == value {
				positions = append(positions, []int{i, j})
			}
		}
	}
	return positions
}

func (g *Gameplay) RestartCurrentLevel() {
	g.currentLayout = CopyLayout(g.lvls[g.currentLvl].Layout)
}

func (g *Gameplay) ReverseCurrentChar() {
	currentValue := g.currentLayout[int(g.currentChar.Y)][int(g.currentChar.X)]
	reversedValue, exists := reverseMap[currentValue]

	if currentValue == codeExit || currentValue == codeStart {
		for i := 0; i < len(g.currentLayout); i++ {
			for j := 0; j < len(g.currentLayout[i]); j++ {
				if g.currentLayout[i][j] == 1-currentValue {
					g.currentLayout[i][j] = currentValue
				}
			}
		}
	}

	if exists {
		g.currentLayout[int(g.currentChar.Y)][int(g.currentChar.X)] = reversedValue
	}
}

func CopyLayout(lin [][]int) [][]int {
	rows := len(lin)
	cols := len(lin)
	var lout [][]int
	lout = make([][]int, rows)
	for i := range lout {
		lout[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			lout[i][j] = lin[i][j]
		}
	}
	return lout
}

func (g *Gameplay) TraversePathFrom(pos *vector.Vector2D) bool {
	log.Print("Traversing path from: ", pos)
	currentValue := g.currentLayout[int(pos.Y)][int(pos.X)]
	log.Print("Current value: ", currentValue)
	dir := movePassageMap[currentValue]
	next := pos.Add(&dir)
	if g.PosInLayoutBounds(next) {
		nextValue := g.currentLayout[int(next.Y)][int(next.X)]
		log.Print("Next value: ", nextValue)
		log.Print("For array of available codes: ", passageFinderMap[currentValue])
		for _, code := range passageFinderMap[currentValue] {
			log.Print("Checking code: ", code)
			if nextValue == codeExit {
				log.Print("Found exit!")
				g.lvlCompleted = true
				return true
			} else if nextValue == code {
				log.Print("Able to pass.")
				g.TraversePathFrom(next)
			} else {
				log.Print("No passage.")
			}
		}
	}
	return false
}

func (g *Gameplay) FindPassage() {
	pos := g.FindInLayout(codeStart)
	current := vector.Vector2D{X: float64(pos[0][0]), Y: float64(pos[0][1])}

	directions := []vector.Vector2D{up, down, right, left}

	availablePaths := []*vector.Vector2D{}
	for _, dir := range directions {
		next := current.Add(&dir)
		log.Print("Checking direction: ", dir)
		log.Print("Checking path: ", next)
		if g.PosInLayoutBounds(next) {
			if movePassageMap[g.currentLayout[int(next.Y)][int(next.X)]] == dir {
				availablePaths = append(availablePaths, next)
				log.Print("Available path found: ", next)
			}
		}
	}

	for _, pos := range availablePaths {
		g.TraversePathFrom(pos)
	}
}

func (g *Gameplay) PosInLayoutBounds(pos *vector.Vector2D) bool {

	rows := len(g.lvls[g.currentLvl].Layout)
	cols := len(g.lvls[g.currentLvl].Layout[0])

	if pos.Y < 0 {
		return false
	} else if int(pos.Y) >= rows {
		return false
	}
	if pos.X < 0 {
		return false
	} else if int(pos.X) >= cols {
		return false
	}
	return true
}

func (g *Gameplay) Update() {
	if g.lvlCompleted {
		log.Print("Level successfully completed!")
		os.Exit(1)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.changeStateTo = STATE_MENU
		g.changeState = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.RestartCurrentLevel()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.currentChar.X += 1
		log.Print(g.currentChar)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.currentChar.X -= 1
		log.Print(g.currentChar)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.currentChar.Y += 1
		log.Print(g.currentChar)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.currentChar.Y -= 1
		log.Print(g.currentChar)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.ReverseCurrentChar()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.FindPassage()
	}

	rows := len(g.lvls[g.currentLvl].Layout)
	cols := len(g.lvls[g.currentLvl].Layout[0])

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
		g.currentLayout = CopyLayout(g.lvls[0].Layout)
		startPos := g.FindInLayout(codeStart)
		g.currentChar.X, g.currentChar.Y = float64(startPos[0][0]), float64(startPos[0][1])
		g.assignLayout = false
		g.lvlCompleted = false
	}

	rows := len(g.currentLayout)
	cols := len(g.currentLayout[0])

	//Calculations so that matrix is drawn around the center of it's "center" element
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
