package day5

import (
	"bufio"
	"fmt"
	"strings"
)

type Line struct {
	X1, X2, Y1, Y2 int
}

type Coordinate struct {
	X, Y int
}

func Solve(countDiagonals bool) (count int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var lines []Line
	for scanner.Scan() {
		s := scanner.Text()
		line := Line{}
		fmt.Sscanf(s, "%d,%d -> %d,%d", &line.X1, &line.Y1, &line.X2, &line.Y2)
		lines = append(lines, line)
	}

	lineMap := make(map[Coordinate]int)

	for _, line := range lines {
		if line.X1 == line.X2 {
			y1, y2 := line.Y1, line.Y2
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				lineMap[Coordinate{X: line.X1, Y: y}]++
			}
		} else if line.Y1 == line.Y2 {
			x1, x2 := line.X1, line.X2
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				lineMap[Coordinate{X: x, Y: line.Y1}]++
			}
		} else if countDiagonals {
			x1, x2 := line.X1, line.X2
			y1, y2 := line.Y1, line.Y2
			if x1 > x2 {
				x1, x2 = x2, x1
				y1, y2 = y2, y1
			}
			for i := 0; i <= x2-x1; i++ {
				if y2 > y1 {
					lineMap[Coordinate{X: x1 + i, Y: y1 + i}]++
				} else {
					lineMap[Coordinate{X: x1 + i, Y: y1 - i}]++
				}
			}
		}
	}

	for _, v := range lineMap {
		if v > 1 {
			count++
		}
	}

	return
}
