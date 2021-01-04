package day11

type Position struct {
	X, Y int
}

func (p *Position) Up() Position {
	return Position{X: p.X, Y: p.Y - 1}
}

func (p *Position) Down() Position {
	return Position{X: p.X, Y: p.Y + 1}
}

func (p *Position) Left() Position {
	return Position{X: p.X - 1, Y: p.Y}
}

func (p *Position) Right() Position {
	return Position{X: p.X + 1, Y: p.Y}
}
