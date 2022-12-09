package day09

import (
	"fmt"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

type Knot struct {
	X int
	Y int
}

type Rope struct {
	Knots   []*Knot
	Visited map[Knot]bool
}

func (r *Rope) MakeKnots(numKnots int) {
	r.Knots = make([]*Knot, numKnots)
	for i := 0; i < numKnots; i++ {
		r.Knots[i] = &Knot{}
	}

	r.Visited = map[Knot]bool{}
}

func (r *Rope) Head() *Knot {
	return r.Knots[0]
}

func (r *Rope) Tail() *Knot {
	return r.Knots[len(r.Knots)-1]
}

func (tail *Knot) MoveTo(head Knot) {
	dx := head.X - tail.X
	dy := head.Y - tail.Y

	if abs(dy) > 1 || abs(dx) > 1 {
		if dy != 0 {
			tail.Y += dy / abs(dy)
		}

		if dx != 0 {
			tail.X += dx / abs(dx)
		}
	}
}

func (r *Rope) MoveTail() {
	for j := 1; j < len(r.Knots); j++ {
		head := r.Knots[j-1]
		tail := r.Knots[j]

		tail.MoveTo(*head)
	}
	r.Visited[*r.Tail()] = true
}

func (r *Rope) Move(direction string, numSteps int) {
	for i := 0; i < numSteps; i++ {
		switch direction {
		case "U":
			r.Head().Y--
		case "D":
			r.Head().Y++
		case "L":
			r.Head().X--
		case "R":
			r.Head().X++
		}
		r.MoveTail()
	}
}

func Solve() (int, int) {
	lines := utils.ReadInput("./solutions/day09/input.txt", mappers.ToString)

	rope1 := Rope{}
	rope1.MakeKnots(2)
	rope2 := Rope{}
	rope2.MakeKnots(10)
	for _, v := range lines {
		var direction string
		var numSteps int
		fmt.Sscanf(v, "%s %d", &direction, &numSteps)

		rope1.Move(direction, numSteps)
		rope2.Move(direction, numSteps)
	}

	return len(rope1.Visited), len(rope2.Visited)
}
