package game

func withinBounds(x, y int) bool {
	return x >= 0 && x < GRID_WIDTH && y >= 0 && y < GRID_HEIGHT
}

func emptyCell(g *Game, x, y int) bool {
	return g.grid[y][x] == nil
}

func sandCell(g *Game, x, y int) bool {
	_, ok := g.grid[y][x].(*Sand)
	return ok
}
