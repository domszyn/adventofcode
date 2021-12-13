package day13

import (
	"bufio"
	"strconv"
	"strings"

	tm "github.com/buger/goterm"
)

type Dot struct{ x, y int }

func Solve() (part1, part2 int) {
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
				tm.Clear()
				tm.Printf("Part 1: %d", part1)
			}

			foldsCount++
		} else {
			coords := strings.Split(s, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			paper[Dot{x, y}] = true
		}
	}

	var cleanPaper []Dot
	for dot, val := range paper {
		if val {
			cleanPaper = append(cleanPaper, dot)
			tm.MoveCursor(dot.x+1, dot.y+2)
			tm.Println("#")
			tm.Flush()
		}
	}

	return
}
