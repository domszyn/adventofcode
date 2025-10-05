package day18

import (
	"math"
	"strings"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

type Point struct{ X, Y int }

type PathInfo struct {
	Distance    int
	DoorsNeeded uint32
}

type Grid struct {
	Map    map[Point]rune
	Keys   map[rune]Point
	Doors  map[rune]Point
	Starts []Point
	Cache  map[Point]map[Point]PathInfo
}

func LoadGrid(input string) Grid {
	grid := Grid{
		Map:    make(map[Point]rune),
		Keys:   make(map[rune]Point),
		Doors:  make(map[rune]Point),
		Starts: make([]Point, 0, 4),
		Cache:  make(map[Point]map[Point]PathInfo),
	}
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		for x, c := range line {
			loc := Point{x, y}
			grid.Map[loc] = c
			if c == '@' || (c >= '0' && c <= '3') {
				grid.Starts = append(grid.Starts, loc)
			} else if c >= 'A' && c <= 'Z' {
				grid.Doors[c] = loc
			} else if c >= 'a' && c <= 'z' {
				grid.Keys[c] = loc
			}
		}
	}

	pointsOfInterest := append([]Point{}, grid.Starts...)
	for _, v := range grid.Keys {
		pointsOfInterest = append(pointsOfInterest, v)
	}

	for _, start := range pointsOfInterest {
		grid.Cache[start] = make(map[Point]PathInfo)
		distances, doors := grid.findPathData(start, pointsOfInterest)

		for _, end := range pointsOfInterest {
			if start == end {
				continue
			}

			if dist, found := distances[end]; found {
				grid.Cache[start][end] = PathInfo{
					Distance:    dist,
					DoorsNeeded: doors[end],
				}
			} else {
				grid.Cache[start][end] = PathInfo{
					Distance:    math.MaxInt64,
					DoorsNeeded: 0,
				}
			}
		}
	}

	return grid
}

func (g Grid) GetNeighbors(p Point) []Point {
	return []Point{
		{X: p.X, Y: p.Y - 1},
		{X: p.X, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y},
	}
}

func (g Grid) findPathData(from Point, endPoints []Point) (map[Point]int, map[Point]uint32) {
	queue := toolbox.NewQueue[Point]()
	distance := make(map[Point]int)
	doors := make(map[Point]uint32)
	endPointsMap := make(map[Point]struct{})
	for _, p := range endPoints {
		endPointsMap[p] = struct{}{}
	}

	var current Point
	queue.Enqueue(from)
	distance[from] = 0

	for queue.Size() > 0 {
		current, _ = queue.Dequeue()
		if _, isEndPoint := endPointsMap[current]; isEndPoint {
			delete(endPointsMap, current)
			// If we've found all our targets, we can stop early
			if len(endPointsMap) == 0 {
				return distance, doors
			}
		}

		for _, n := range g.GetNeighbors(current) {
			if _, visited := distance[n]; visited || g.Map[n] == '#' {
				continue
			}

			cell := g.Map[n]
			newDoorMask := doors[current]
			if cell >= 'A' && cell <= 'Z' {
				newDoorMask |= (1 << (cell - 'A'))
			}

			distance[n] = distance[current] + 1
			doors[n] = newDoorMask
			queue.Enqueue(n)
		}
	}

	return distance, doors
}

func (g Grid) getLocation(from rune) Point {
	if from == '@' {
		return g.Starts[0]
	} else if from >= '0' && from <= '3' {
		fromIdx := from - '0'
		return g.Starts[fromIdx]
	}
	return g.Keys[from]
}

func (g Grid) h(from SearchState) int {
	if from.HasAllKeys() {
		return 0
	}

	minDistance := math.MaxInt64

	for _, lk := range from.LastKeys {
		startPoint := g.getLocation(lk)

		for k := 'a'; k <= 'z'; k++ {
			// Skip keys we already have
			if from.HasKey(k) {
				continue
			}

			keyLoc := g.Keys[k]
			distance := g.Cache[startPoint][keyLoc].Distance
			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	return minDistance
}
