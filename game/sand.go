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
			color: color.RGBA{255, 255, 0, 255},
		},
	}
}

func (s *Sand) Update(g *Game, x, y int) {
	if y+1 >= GRID_HEIGHT {
		return
	}
	if g.grid[y+1][x] == nil {
		g.grid[y+1][x] = g.grid[y][x]
		g.grid[y][x] = nil
	} else if x-1 >= 0 && g.grid[y+1][x-1] == nil {
		g.grid[y+1][x-1] = g.grid[y][x]
		g.grid[y][x] = nil
	} else if x+1 < GRID_WIDTH && g.grid[y+1][x+1] == nil {
		g.grid[y+1][x+1] = g.grid[y][x]
		g.grid[y][x] = nil
	}
}

func (s *Sand) Draw(screen *ebiten.Image, x, y int) {
	vector.DrawFilledRect(screen, float32(x*CELL_WIDTH), float32(y*CELL_HEIGHT), CELL_WIDTH, CELL_HEIGHT, s.color, false)
}
