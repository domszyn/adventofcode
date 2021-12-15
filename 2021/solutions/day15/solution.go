package day15

import (
	"bufio"
	"math"
	"strings"
)

type Point struct {
	Value   int
	Visited bool
	Risk    int
}

func findPointWithLowestRisk(points [][]Point) (int, int) {
	minRisk := math.MaxInt64

	for _, row := range points {
		for _, p := range row {
			if !p.Visited && p.Risk < minRisk {
				minRisk = p.Risk
			}
		}
	}

	for j, row := range points {
		for i, p := range row {
			if !p.Visited && p.Risk == minRisk {
				return i, j
			}
		}
	}

	return 0, 0
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func allVisited(points [][]Point) bool {
	for _, row := range points {
		for _, p := range row {
			if !p.Visited {
				return false
			}
		}
	}

	return true
}

func getMinRisk(points [][]Point) int {
	points[0][0].Risk = 0

	x, y := 0, 0
	for !points[len(points)-1][len(points[0])-1].Visited {
		if x+1 < len(points[0]) && !points[y][x+1].Visited {
			points[y][x+1].Risk = min(points[y][x+1].Risk, points[y][x].Risk+points[y][x].Value)
		}

		if y+1 < len(points) && !points[y+1][x].Visited {
			points[y+1][x].Risk = min(points[y+1][x].Risk, points[y][x].Risk+points[y][x].Value)
		}

		if x-1 > 0 && !points[y][x-1].Visited {
			points[y][x-1].Risk = min(points[y][x-1].Risk, points[y][x].Risk+points[y][x].Value)
		}

		if y-1 > 0 && !points[y-1][x].Visited {
			points[y-1][x].Risk = min(points[y-1][x].Risk, points[y][x].Risk+points[y][x].Value)
		}

		points[y][x].Visited = true
		x, y = findPointWithLowestRisk(points)
	}

	return points[len(points)-1][len(points[0])-1].Risk + points[len(points)-1][len(points[0])-1].Value - points[0][0].Value
}

func SolvePart1() (part1 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var rows []string
	var points [][]Point
	for scanner.Scan() {
		s := scanner.Text()
		rows = append(rows)
		var pp []Point
		for _, v := range s {
			value := int(v - '0')
			pp = append(pp, Point{Value: value, Risk: math.MaxInt64})
		}
		points = append(points, pp)
	}

	part1 = getMinRisk(points)

	// for _, row := range points {
	// 	for _, p := range row {
	// 		fmt.Printf("%d(%d)\t", p.Value, p.Risk)
	// 	}
	// 	fmt.Println()
	// }

	return
}

func SolvePart2() (part1 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var rows []string
	var pattern [][]Point
	for scanner.Scan() {
		s := scanner.Text()
		rows = append(rows)
		var pp []Point
		for _, v := range s {
			pp = append(pp, Point{Value: int(v - '0')})
		}
		pattern = append(pattern, pp)
	}

	pattern[0][0].Risk = 0

	points := make([][]Point, len(pattern)*5)

	for y := 0; y < len(points); y++ {
		points[y] = make([]Point, len(pattern)*5)

		for x := 0; x < len(points); x++ {
			offset := int(y/len(pattern)) + int(x/len(pattern))
			points[y][x] = Point{
				Value:   pattern[y%len(pattern)][x%len(pattern)].Value + offset,
				Visited: false,
				Risk:    math.MaxInt64,
			}

			if points[y][x].Value > 9 {
				points[y][x].Value -= 9
			}
		}
	}

	part1 = getMinRisk(points)

	return
}
