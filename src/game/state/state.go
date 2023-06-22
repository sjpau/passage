package state

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Draw(screen *ebiten.Image)
	Update()
	Init()        //TODO: pass slice of ints as options?
	Change() bool //TODO: return int as a state number
}
