package game

import "github.com/hajimehoshi/ebiten/v2"

type Particle interface {
	Update(g *Game, x, y int)
	Draw(screen *ebiten.Image, x, y int)
	SetAlreadyUpdated(updated bool)
	IsAlreadyUpdated() bool
}
