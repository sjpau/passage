package game

import (
	"embed"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/passage/src/game/state"
)

type Game struct {
	states       []state.State
	currentState int
}

func NewGame(lvlFS *embed.FS) *Game {
	g := &Game{
		states:       []state.State{&state.Menu{}, &state.Gameplay{LvlFS: lvlFS}},
		currentState: state.STATE_MENU,
	}
	for i := range g.states {
		g.states[i].Init()
	}
	return g
}

func (self *Game) Update() error {
	self.states[self.currentState].Update()
	changeStatus := self.states[self.currentState].Change()
	if changeStatus != -1 {
		self.currentState = changeStatus
	}
	return nil
}

func (self *Game) Draw(screen *ebiten.Image) {
	self.states[self.currentState].Draw(screen)
}

func (self *Game) Layout(w, h int) (int, int) {
	f := ebiten.DeviceScaleFactor()
	sw := int(math.Ceil(float64(w) * f))
	sh := int(math.Ceil(float64(h) * f))
	return sw, sh
}
