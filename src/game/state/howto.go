package state

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sjpau/passage/src/graphics"
	"github.com/tinne26/etxt"
	"github.com/tinne26/fonts/liberation/lbrtserif"
)

type Howto struct {
	changeState   bool
	changeStateTo int
	text          *etxt.Renderer
}

func (h *Howto) Init() {
	renderer := etxt.NewStdRenderer()
	renderer.SetFont(lbrtserif.Font())
	cache := etxt.NewDefaultCache(1024 * 1024)
	renderer.SetCacheHandler(cache.NewHandler())
	renderer.SetColor(graphics.COLOR_RED)
	renderer.SetSizePx(int(fontSize / 2 * ebiten.DeviceScaleFactor()))
	renderer.SetAlign(etxt.YCenter, etxt.XCenter)
	h.text = renderer

}

func (h *Howto) Change() int {
	if h.changeState {
		h.changeState = false
		return h.changeStateTo
	} else {
		return -1
	}
}

func (h *Howto) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		h.changeStateTo = STATE_MENU
		h.changeState = true
	}
}

func (h *Howto) Draw(screen *ebiten.Image) {
	screen.Fill(graphics.COLOR_GREY)
	h.text.SetTarget(screen)
	bounds := screen.Bounds()
	x, y := bounds.Dx()/2, bounds.Dy()/2
	h.text.SetAlign(etxt.YCenter, etxt.XCenter)
	txt := fmt.Sprintf("Move around the field using WASD keys.\nPress SPACE to reverse a glyph.\nCreate a path from * to O and press ENTER to complete a level.\n")
	h.text.Draw(txt, x, y)
	h.text.SetAlign(etxt.Top, etxt.Left)
	h.text.Draw("Esc", 0, 0)
}
