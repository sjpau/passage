package main

import (
	"embed"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/passage/src/game"
)

//go:embed src/lvl/conf/*
var lvlFS embed.FS

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ebiten.SetWindowFloating(true)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	ebiten.SetWindowTitle("The Passage")
	g := game.NewGame(&lvlFS)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
