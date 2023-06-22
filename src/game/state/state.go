package state

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Draw(screen *ebiten.Image)
	Update()
	Init()       //TODO: pass slice of ints as options?
	Change() int //TODO: return int as a state number
}

const (
	STATE_MENU = iota
	STATE_GAME
	STATE_HOWTO
)
