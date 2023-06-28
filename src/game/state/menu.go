package state

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sjpau/passage/src/graphics"
	"github.com/tinne26/etxt"
	"github.com/tinne26/fonts/liberation/lbrtserif"
)

type Menu struct {
	changeState   bool
	changeStateTo int
	text          *etxt.Renderer
	currentOpt    int
	lenOpt        int
}

var MenuOptions = []string{
	"Start",
	"HowTo",
	"Quit",
}

const (
	START = iota
	HOWTO
	QUIT
)

func (m *Menu) Change() int {
	if m.changeState {
		m.changeState = false
		return m.changeStateTo
	} else {
		return -1
	}
}

func (m *Menu) Init() {
	m.changeState = false
	renderer := etxt.NewStdRenderer()
	renderer.SetFont(lbrtserif.Font())
	cache := etxt.NewDefaultCache(1024 * 1024)
	renderer.SetCacheHandler(cache.NewHandler())
	renderer.SetColor(graphics.COLOR_RED)
	renderer.SetSizePx(int(fontSize * ebiten.DeviceScaleFactor()))
	renderer.SetAlign(etxt.YCenter, etxt.XCenter)
	m.text = renderer
	m.lenOpt = len(MenuOptions)
}

func (m *Menu) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) && m.currentOpt != 0 {
		m.currentOpt--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) && m.currentOpt < m.lenOpt-1 {
		m.currentOpt++
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch m.currentOpt {
		case START:
			m.changeStateTo = STATE_GAME
			m.changeState = true
			break
		case HOWTO:
			m.changeStateTo = STATE_HOWTO
			m.changeState = true
			break
		case QUIT:
			os.Exit(0)
			break
		}
	}

}

//TODO change colors to consts
func (m *Menu) Draw(screen *ebiten.Image) {
	screen.Fill(graphics.COLOR_GREY)
	m.text.SetTarget(screen)
	bounds := screen.Bounds()
	x, y := bounds.Dx()/2, bounds.Dy()/2
	for i, option := range MenuOptions {
		text := option
		if i == m.currentOpt {
			m.text.SetColor(graphics.COLOR_PINK)
		} else {
			m.text.SetColor(graphics.COLOR_RED)
		}
		offsetY := (i - len(MenuOptions)/2) * int(verticalSP*ebiten.DeviceScaleFactor())
		yPos := y + offsetY
		m.text.Draw(text, x, yPos)
	}
}
