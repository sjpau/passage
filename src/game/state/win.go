package state

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sjpau/passage/src/graphics"
	"github.com/tinne26/etxt"
	"github.com/tinne26/fonts/liberation/lbrtserif"
)

type Win struct {
	changeState   bool
	changeStateTo int
	text          *etxt.Renderer
}

func (w *Win) Init() {
	renderer := etxt.NewStdRenderer()
	renderer.SetFont(lbrtserif.Font())
	cache := etxt.NewDefaultCache(1024 * 1024)
	renderer.SetCacheHandler(cache.NewHandler())
	renderer.SetColor(graphics.COLOR_RED)
	renderer.SetSizePx(int(fontSize / 2 * ebiten.DeviceScaleFactor()))
	renderer.SetAlign(etxt.YCenter, etxt.XCenter)
	w.text = renderer

}

func (w *Win) Change() int {
	if w.changeState {
		w.changeState = false
		return w.changeStateTo
	} else {
		return -1
	}
}

func (w *Win) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		w.changeStateTo = STATE_MENU
		w.changeState = true
	}
}

func (w *Win) Draw(screen *ebiten.Image) {
	screen.Fill(graphics.COLOR_GREY)
	w.text.SetTarget(screen)
	bounds := screen.Bounds()
	x, y := bounds.Dx()/2, bounds.Dy()/2
	w.text.SetAlign(etxt.YCenter, etxt.XCenter)
	txt := fmt.Sprintf("You have finished the game!\nYep, that's it.\nThanks for playing!")
	w.text.Draw(txt, x, y)
	w.text.SetAlign(etxt.Top, etxt.Left)
	w.text.Draw("Esc", 0, 0)
}
