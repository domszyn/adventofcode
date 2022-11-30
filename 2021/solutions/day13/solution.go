package day13

import (
	"bufio"
	"strconv"
	"strings"
)

type Dot struct{ x, y int }

func Solve() (part1 int, part2 string) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	paper := make(map[Dot]bool)
	foldsCount := 0
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}
		if strings.Index(s, "fold along") == 0 {
			folds := strings.Split(s[11:], "=")
			foldSize, _ := strconv.Atoi(folds[1])
			if folds[0] == "x" {
				for dot, val := range paper {
					if !val {
						continue
					}
					if dot.x > foldSize {
						paper[Dot{x: 2*foldSize - dot.x, y: dot.y}] = true
						paper[dot] = false
					}
				}
			}

			if folds[0] == "y" {
				for dot, val := range paper {
					if !val {
						continue
					}
					if dot.y > foldSize {
						paper[Dot{x: dot.x, y: 2*foldSize - dot.y}] = true
						paper[dot] = false
					}
				}
			}

			if foldsCount == 0 {
				for _, val := range paper {
					if val {
						part1++
					}
				}
			}

			foldsCount++
		} else {
			coords := strings.Split(s, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			paper[Dot{x, y}] = true
		}
	}

	var maxX, maxY int
	for dot, val := range paper {
		if val && dot.x > maxX {
			maxX = dot.x
		}
		if val && dot.y > maxY {
			maxY = dot.y
		}
	}

	part2 = "\n\n"
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if paper[Dot{x, y}] {
				part2 += "#"
			} else {
				part2 += "."
			}
		}
		part2 += "\n"
	}

	return
}
