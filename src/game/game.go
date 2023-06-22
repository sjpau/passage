package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/passage/src/game/state"
)

type Game struct {
	states       []state.State
	currentState int
}

func NewGame() *Game {
	g := &Game{
		states:       []state.State{&state.Menu{}, &state.Gameplay{}},
		currentState: 0,
	}
	for i := range g.states {
		g.states[i].Init()
	}
	return g
}

func (self *Game) Update() error {
	self.states[self.currentState].Update()
	if self.states[self.currentState].Change() {
		self.currentState += 1
	}
	return nil
}

func (self *Game) Draw(screen *ebiten.Image) {
	self.states[self.currentState].Draw(screen)
}

func (self *Game) Layout(w, h int) (int, int) {
	f := ebiten.DeviceScaleFactor()
	return w * int(f), h * int(f)
}
