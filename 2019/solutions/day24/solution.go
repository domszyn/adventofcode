package day24

import (
	"math"
	"strings"
)

// ....#
// #..#.
// ##..#
// #.###
// .####

func getNeighbors(i int) []int {
	var neighbors []int

	if i >= 5 {
		neighbors = append(neighbors, i-5)
	}

	if i%5 != 0 {
		neighbors = append(neighbors, i-1)
	}

	if i%5 != 4 {
		neighbors = append(neighbors, i+1)
	}

	if i < 20 {
		neighbors = append(neighbors, i+5)
	}

	return neighbors
}

func nextLayout(layout string) string {
	next := make([]rune, len(layout))

	for i, curr := range layout {
		var numBugsAround int
		for _, n := range getNeighbors(i) {
			if n < 0 || n >= len(layout) {
				continue
			}

			if layout[n] == '#' {
				numBugsAround++
			}
		}

		next[i] = curr

		if curr == '#' && numBugsAround != 1 {
			next[i] = '.'
		} else if curr == '.' && (numBugsAround == 1 || numBugsAround == 2) {
			next[i] = '#'
		}
	}

	return string(next)
}

func rating(layout string) int {
	var rating int
	for i, c := range layout {
		if c == '#' {
			rating += int(math.Pow(2, float64(i)))
		}
	}

	return rating
}

func solvePart1() int {
	layouts := make(map[string]bool)
	layout := "....##..#.##..##.###.####"
	layouts[layout] = true

	for {
		next := nextLayout(layout)

		if _, found := layouts[next]; found {
			return rating(next)
		}

		layouts[next] = true
		layout = next
	}
}

type RFS struct {
	Levels             map[int][]rune
	MinLevel, MaxLevel int
}

type Position struct{ Level, Index int }

func NewSpace(layout string) RFS {
	levels := make(map[int][]rune)
	levels[0] = []rune(layout)
	return RFS{
		Levels:   levels,
		MinLevel: 0,
		MaxLevel: 0,
	}
}

func (p Position) IsOnEdge() bool {
	// 0  1  2  3  4
	// 5           9
	// 10          14
	// 15          19
	// 20 21 22 23 24

	return p.Index%5 == 0 ||
		(p.Index+1)%5 == 0 ||
		p.Index < 5 ||
		p.Index > 20
}

func (p Position) IsAdjacentToCenter() bool {
	return p.Index == 7 ||
		p.Index == 11 ||
		p.Index == 13 ||
		p.Index == 17
}

func (space RFS) GetNeighbors(pos Position) []Position {
	var neighbors []Position

	for _, n := range getNeighbors(pos.Index) {
		if n == 12 {
			continue // skip center position
		}

		neighbors = append(neighbors, Position{Level: pos.Level, Index: n})
	}

	if !pos.IsOnEdge() && !pos.IsAdjacentToCenter() {
		return neighbors
	}

	if pos.IsOnEdge() {
		if pos.Index < 5 {
			neighbors = append(neighbors, Position{Level: pos.Level - 1, Index: 7})
		}

		if pos.Index >= 20 {
			neighbors = append(neighbors, Position{Level: pos.Level - 1, Index: 17})
		}

		if pos.Index%5 == 0 {
			neighbors = append(neighbors, Position{Level: pos.Level - 1, Index: 11})
		}

		if (pos.Index+1)%5 == 0 {
			neighbors = append(neighbors, Position{Level: pos.Level - 1, Index: 13})
		}

		return neighbors
	}

	var outerNeighbors []Position
	switch pos.Index {
	case 7:
		outerNeighbors = []Position{
			{Level: pos.Level + 1, Index: 0},
			{Level: pos.Level + 1, Index: 1},
			{Level: pos.Level + 1, Index: 2},
			{Level: pos.Level + 1, Index: 3},
			{Level: pos.Level + 1, Index: 4},
		}
	case 11:
		outerNeighbors = []Position{
			{Level: pos.Level + 1, Index: 0},
			{Level: pos.Level + 1, Index: 5},
			{Level: pos.Level + 1, Index: 10},
			{Level: pos.Level + 1, Index: 15},
			{Level: pos.Level + 1, Index: 20},
		}
	case 13:
		outerNeighbors = []Position{
			{Level: pos.Level + 1, Index: 4},
			{Level: pos.Level + 1, Index: 9},
			{Level: pos.Level + 1, Index: 14},
			{Level: pos.Level + 1, Index: 19},
			{Level: pos.Level + 1, Index: 24},
		}
	case 17:
		outerNeighbors = []Position{
			{Level: pos.Level + 1, Index: 20},
			{Level: pos.Level + 1, Index: 21},
			{Level: pos.Level + 1, Index: 22},
			{Level: pos.Level + 1, Index: 23},
			{Level: pos.Level + 1, Index: 24},
		}
	}

	neighbors = append(neighbors, outerNeighbors...)
	return neighbors
}

func (space RFS) Expand() RFS {
	newSpace := RFS{
		Levels:   make(map[int]([]rune)),
		MinLevel: space.MinLevel - 1,
		MaxLevel: space.MaxLevel + 1,
	}

	for level := newSpace.MinLevel; level <= newSpace.MaxLevel; level++ {
		newSpace.Levels[level] = []rune(strings.Repeat(".", 25))

		for i := range 25 {
			if i == 12 {
				continue
			}

			pos := Position{Level: level, Index: i}
			neighbors := space.GetNeighbors(pos)

			var numBugsAround int
			for _, n := range neighbors {
				if oldLevel, ok := space.Levels[n.Level]; ok {
					if oldLevel[n.Index] == '#' {
						numBugsAround++
					}
				}
			}

			currentState := '.'
			if oldLevel, ok := space.Levels[level]; ok {
				currentState = oldLevel[i]
			}

			nextState := currentState

			if currentState == '#' && numBugsAround != 1 {
				nextState = '.'
			} else if currentState == '.' && (numBugsAround == 1 || numBugsAround == 2) {
				nextState = '#'
			}

			newSpace.Levels[pos.Level][pos.Index] = nextState
		}
	}

	return newSpace
}

func (space RFS) TotalBugs() int {
	var totalBugs int
	for level := space.MinLevel; level <= space.MaxLevel; level++ {
		for i := range 25 {
			if space.Levels[level][i] == '#' {
				totalBugs++
			}
		}
	}

	return totalBugs
}

func solvePart2() int {
	space := NewSpace("....##..#.##..##.###.####")

	for range 200 {
		space = space.Expand()
	}

	return space.TotalBugs()
}

func GetAnswers() (int, int) {
	return solvePart1(), solvePart2()
}
