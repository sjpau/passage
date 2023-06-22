package state

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Menu struct {
	changeState bool
	SW          float64
	SH          float64
}

func (m *Menu) Change() bool {
	return m.changeState
}

func (m *Menu) Init() {
	m.changeState = false
}

func (m *Menu) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		m.changeState = true
	}
}

func (m *Menu) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 20, 255})
}
