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
	// TODO: add flow direction
	// TODO: make cleaner
	if withinBounds(x, y+1) && emptyCell(g, x, y+1) {
		g.grid[y+1][x] = g.grid[y][x]
		g.grid[y][x] = nil
	} else if withinBounds(x-1, y+1) && emptyCell(g, x-1, y+1) {
		g.grid[y+1][x-1] = g.grid[y][x]
		g.grid[y][x] = nil
	} else if withinBounds(x+1, y+1) && emptyCell(g, x+1, y+1) {
		g.grid[y+1][x+1] = g.grid[y][x]
		g.grid[y][x] = nil
	} else if withinBounds(x-1, y) && emptyCell(g, x-1, y) {
		g.grid[y][x-1] = g.grid[y][x]
		g.grid[y][x] = nil
	} else if withinBounds(x+1, y) && emptyCell(g, x+1, y) {
		g.grid[y][x+1] = g.grid[y][x]
		g.grid[y][x] = nil
	}
}

func withinBounds(x, y int) bool {
	return x >= 0 && x < GRID_WIDTH && y >= 0 && y < GRID_HEIGHT
}

func emptyCell(g *Game, x, y int) bool {
	return g.grid[y][x] == nil
}

func (w *Water) Draw(screen *ebiten.Image, x, y int) {
	vector.DrawFilledRect(screen, float32(x*CELL_WIDTH), float32(y*CELL_HEIGHT), CELL_WIDTH, CELL_HEIGHT, w.color, false)
}
