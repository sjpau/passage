package state

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Draw(screen *ebiten.Image)
	Update()
	Init() //TODO: pass slice of ints as options?
	Change() int
}

const (
	STATE_MENU = iota
	STATE_GAME
	STATE_HOWTO
)

const fontSize = 72
const verticalSP = fontSize + 5
const horizontalSP = fontSize + 5
