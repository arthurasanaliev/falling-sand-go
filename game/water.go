package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Water struct {
	BaseParticle
}

func NewWater() *Water {
	return &Water{
		BaseParticle: BaseParticle{
			color: color.RGBA{0, 150, 255, 255},
		},
	}
}

func (w *Water) Update(g *Game, x, y int) {
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

func (w *Water) Draw(screen *ebiten.Image, x, y int) {
	vector.DrawFilledRect(screen, float32(x*CELL_WIDTH), float32(y*CELL_HEIGHT), CELL_WIDTH, CELL_HEIGHT, w.color, false)
}
