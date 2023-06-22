package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/passage/src/game"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ebiten.SetWindowFloating(true)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	ebiten.SetWindowTitle("Ethereal")
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
