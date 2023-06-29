package game

import (
	"bytes"
	"embed"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/sjpau/passage/src/game/state"
	"github.com/sjpau/passage/src/help"
	"github.com/sjpau/passage/src/sfx"
)

type Game struct {
	states       []state.State
	currentState int
	audioPlayer  *audio.Player
	audioContext *audio.Context
}

const (
	sampleRate     = 44100
	bytesPerSample = 4
)

func NewGame(lvlFS *embed.FS) *Game {
	g := &Game{
		states:       []state.State{&state.Menu{}, &state.Gameplay{LvlFS: lvlFS}, &state.Howto{}, &state.Win{}},
		currentState: state.STATE_MENU,
	}
	for i := range g.states {
		g.states[i].Init()
	}
	return g
}

func (self *Game) Update() error {
	go func() error {
		if self.audioPlayer != nil {
			return nil
		}

		if self.audioContext == nil {
			self.audioContext = audio.NewContext(sampleRate)
		}

		sfx_hfe, e := vorbis.DecodeWithoutResampling(bytes.NewReader(sfx.Here_for_eternity))
		help.Check(e)
		s := audio.NewInfiniteLoop(sfx_hfe, 15*bytesPerSample*sampleRate)

		self.audioPlayer, e = self.audioContext.NewPlayer(s)
		help.Check(e)

		self.audioPlayer.Play()
		return nil
	}()
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
