package day9

import (
	"bufio"
	"math"
	"sort"
	"strings"
)

type Point struct{ X, Y int }

func getHeight(x, y int, rows []string) int {
	if x < 0 || x >= len(rows[0]) || y < 0 || y >= len(rows) {
		return math.MaxInt64
	}

	return int(rows[y][x] - '0')
}

func isLowPoint(x, y int, rows []string) (int, bool) {
	height := getHeight(x, y, rows)

	if getHeight(x-1, y, rows) > height &&
		getHeight(x+1, y, rows) > height &&
		getHeight(x, y-1, rows) > height &&
		getHeight(x, y+1, rows) > height {
		return int(1 + height), true
	}

	return -1, false
}

func findAdjacent(p Point, rows []string) (adjacent []Point) {
	height := getHeight(p.X, p.Y, rows)

	points := []Point{
		{X: p.X - 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y},
		{X: p.X, Y: p.Y - 1},
		{X: p.X, Y: p.Y + 1},
	}
	for _, pp := range points {
		if h := getHeight(pp.X, pp.Y, rows); h > height && h < 9 {
			adjacent = append(adjacent, pp)
		}
	}

	return
}

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var rows []string
	for scanner.Scan() {
		s := scanner.Text()

		rows = append(rows, s)
	}

	var basinSizes []int
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[y]); x++ {
			if height, found := isLowPoint(x, y, rows); found {
				part1 += height

				adjacentPoints := []Point{{X: x, Y: y}}
				basinPoints := make(map[Point]bool)
				for {
					var nextRing []Point
					for _, p := range adjacentPoints {
						nextRing = append(nextRing, findAdjacent(p, rows)...)
					}
					if len(nextRing) == 0 {
						break
					}
					for _, point := range nextRing {
						basinPoints[point] = true
					}
					adjacentPoints = nextRing
				}
				basinSizes = append(basinSizes, len(basinPoints)+1)
			}

			if len(basinSizes) > 3 {
				sort.Ints(basinSizes)
				basinSizes = basinSizes[len(basinSizes)-3:]
			}
		}
	}

	part2 = 1
	for _, size := range basinSizes {
		part2 *= size
	}

	return
}
