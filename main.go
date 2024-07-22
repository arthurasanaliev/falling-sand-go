package main

import (
	"falling-sand-go/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowSize(game.SCREEN_WIDTH, game.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Falling Sand")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
