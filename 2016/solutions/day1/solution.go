package day1

import (
	"fmt"
	"strconv"
	"strings"
)

const Input = "R2, L3, R2, R4, L2, L1, R2, R4, R1, L4, L5, R5, R5, R2, R2, R1, L2, L3, L2, L1, R3, L5, R187, R1, R4, L1, R5, L3, L4, R50, L4, R2, R70, L3, L2, R4, R3, R194, L3, L4, L4, L3, L4, R4, R5, L1, L5, L4, R1, L2, R4, L5, L3, R4, L5, L5, R5, R3, R5, L2, L4, R4, L1, R3, R1, L1, L2, R2, R2, L3, R3, R2, R5, R2, R5, L3, R2, L5, R1, R2, R2, L4, L5, L1, L4, R4, R3, R1, R2, L1, L2, R4, R5, L2, R3, L4, L5, L5, L4, R4, L2, R1, R1, L2, L3, L2, R2, L4, R3, R2, L1, L3, L2, L4, L4, R2, L3, L3, R2, L4, L3, R4, R3, L2, L1, L4, R4, R2, L4, L4, L5, L1, R2, L5, L2, L3, R2, L2"

type Direction struct {
	LR       string
	Distance int
}

type Location struct{ X, Y, R int }

func (l *Location) Move(d Direction) {
	switch d.LR {
	case "L":
		l.R -= 90
	case "R":
		l.R += 90
	}
	if l.R < 0 {
		l.R += 360
	}

	switch l.R % 360 {
	case 0:
		l.Y += d.Distance
	case 90:
		l.X += d.Distance
	case 180:
		l.Y -= d.Distance
	case 270:
		l.X -= d.Distance
	}
	fmt.Printf("%s%d [%d, %d]\n", d.LR, d.Distance, l.X, l.Y)
}

func getDirections() (directions []Direction) {
	for _, s := range strings.Split(Input, ", ") {
		d := Direction{LR: s[0:1]}
		d.Distance, _ = strconv.Atoi(s[1:])
		directions = append(directions, d)
	}

	return
}

func SolvePart1() int {
	l := Location{}
	for _, d := range getDirections() {
		l.Move(d)
	}

	if l.X < 0 {
		l.X = -l.X
	}
	if l.Y < 0 {
		l.Y = -l.Y
	}
	return l.X + l.Y
}

func SolvePart2() int {
	l := Location{}
	visited := make(map[Location]int)
	for _, d := range getDirections() {
		pl := Location{X: l.X, Y: l.Y}
		l.Move(d)

		var vl Location
		var bail bool
		for y := pl.Y + 1; y <= l.Y; y++ {
			vl = Location{X: pl.X, Y: y}
			if visited[vl]++; visited[vl] > 1 {
				bail = true
				break
			}
		}
		for y := pl.Y - 1; y >= l.Y; y-- {
			vl = Location{X: pl.X, Y: y}
			if visited[vl]++; visited[vl] > 1 {
				bail = true
				break
			}
		}
		for x := pl.X + 1; x <= l.X; x++ {
			vl = Location{X: x, Y: pl.Y}
			if visited[vl]++; visited[vl] > 1 {
				bail = true
				break
			}
		}
		for x := pl.X - 1; x >= l.X; x-- {
			vl = Location{X: x, Y: pl.Y}
			if visited[vl]++; visited[vl] > 1 {
				bail = true
				break
			}
		}
		if bail {
			l = vl
			break
		}
	}

	if l.X < 0 {
		l.X = -l.X
	}
	if l.Y < 0 {
		l.Y = -l.Y
	}
	return l.X + l.Y
}
