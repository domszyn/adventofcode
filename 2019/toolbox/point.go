package toolbox

type Point struct{ X, Y int }

func (p Point) GetNeighbors() []Point {
	return []Point{
		{p.X, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X, p.Y + 1},
		{p.X - 1, p.Y},
	}
}
