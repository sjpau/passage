package state

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Gameplay struct {
	changeState bool
	SW          float64
	SH          float64
}

func (g *Gameplay) Change() bool {
	return g.changeState
}

func (g *Gameplay) Init() {
}

func (g *Gameplay) Update() {
}

func (g *Gameplay) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
}
