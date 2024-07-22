package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Sand struct {
	BaseParticle
}

func NewSand() *Sand {
	return &Sand{
		BaseParticle: BaseParticle{
			color:          color.RGBA{255, 255, 100, 255},
			alreadyUpdated: false,
		},
	}
}

func (s *Sand) Update(g *Game, x, y int) {
	// TODO: modify sand-water swap
	if s.IsAlreadyUpdated() {
		return
	}
	if !withinBounds(x, y+1) {
		return
	}
	if !sandCell(g, x, y+1) {
		g.grid[y+1][x], g.grid[y][x] = g.grid[y][x], g.grid[y+1][x]
	} else if withinBounds(x-1, y+1) && !sandCell(g, x-1, y+1) {
		g.grid[y+1][x-1], g.grid[y][x] = g.grid[y][x], g.grid[y+1][x-1]
	} else if withinBounds(x+1, y+1) && !sandCell(g, x+1, y+1) {
		g.grid[y+1][x+1], g.grid[y][x] = g.grid[y][x], g.grid[y+1][x+1]
	}
	s.SetAlreadyUpdated(true)
}

func (s *Sand) Draw(screen *ebiten.Image, x, y int) {
	vector.DrawFilledRect(screen, float32(x*CELL_WIDTH), float32(y*CELL_HEIGHT), CELL_WIDTH, CELL_HEIGHT, s.color, false)
}
