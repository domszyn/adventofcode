package day13

type GridRow map[int]int
type Grid map[int]GridRow

const (
	EmptyTile = iota
	Wall
	Block
	Paddle
	Ball
)

func InitGrid() Grid {
	return make(map[int]GridRow)
}

func (g Grid) Row(y int) (row GridRow) {
	var found bool
	if row, found = g[y]; !found {
		row = make(map[int]int)
		g[y] = row
	}
	return
}

func (g Grid) CountTiles(id int) int {
	var count int
	for _, row := range g {
		for _, tile := range row {
			if tile == id {
				count++
			}
		}
	}

	return count
}

func (g Grid) CountBlocks() int {
	return g.CountTiles(Block)
}

func (g Grid) findTile(id int) (int, int) {
	for y, row := range g {
		for x, tile := range row {
			if tile == id {
				return x, y
			}
		}
	}

	return -1, -1
}

func (g Grid) FindPaddle() (int, int) {
	return g.findTile(Paddle)
}

func (g Grid) FindBall() (int, int) {
	return g.findTile(Ball)
}
