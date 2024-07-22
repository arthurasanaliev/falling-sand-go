package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	grid [][]Particle
}

func NewGame() *Game {
	grid := make([][]Particle, GRID_HEIGHT)
	for i := range grid {
		grid[i] = make([]Particle, GRID_WIDTH)
	}
	return &Game{
		grid: grid,
	}
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		gridX := x / CELL_WIDTH
		gridY := y / CELL_HEIGHT
		if gridX >= 0 && gridX < GRID_WIDTH && gridY >= 0 && gridY < GRID_HEIGHT {
			if g.grid[gridY][gridX] == nil {
				g.grid[gridY][gridX] = NewSand()
			}
		}
	}
	for y := GRID_HEIGHT - 1; y >= 0; y-- {
		for x := 0; x < GRID_WIDTH; x++ {
			if g.grid[y][x] != nil {
				g.grid[y][x].Update(g, x, y)
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for y, rows := range g.grid {
		for x := range rows {
			if g.grid[y][x] != nil {
				g.grid[y][x].Draw(screen, x, y)
			}
		}
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
