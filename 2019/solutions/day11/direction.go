package day11

type Direction int

func (d *Direction) RotateLeft() Direction {
	return (*d - 90 + 360) % 360
}

func (d *Direction) RotateRight() Direction {
	return (*d + 90) % 360
}
