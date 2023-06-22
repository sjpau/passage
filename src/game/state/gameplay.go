package state

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sjpau/passage/src/graphics"
	"github.com/sjpau/passage/src/help"
	"github.com/sjpau/passage/src/lvl"
	"github.com/tinne26/etxt"
)

type Gameplay struct {
	changeState   bool
	changeStateTo int
	text          *etxt.Renderer
	LvlFS         *embed.FS
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
	basePath := "src/lvl/conf/"

	//TODO create slice of lvls to append to
	lvlEscape, e := lvl.LoadFromJSON(g.LvlFS, basePath+"escape.json")
	log.Print(lvlEscape)
	help.Check(e)

}

func (g *Gameplay) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.changeStateTo = STATE_MENU
		g.changeState = true
	}
}

func (g *Gameplay) Draw(screen *ebiten.Image) {
	screen.Fill(graphics.COLOR_GREY)
}
