package day3

import (
	"bufio"
	"fmt"
	"strings"
)

type Triangle struct{ A, B, C int }

func (t Triangle) IsPossible() bool {
	return t.A+t.B > t.C && t.A+t.C > t.B && t.B+t.C > t.A
}

func parseInput() (triangles []Triangle) {
	t := Triangle{}
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "  %3d  %3d  %3d", &t.A, &t.B, &t.C)
		triangles = append(triangles, t)
	}
	return
}

func parseInputVertically() (triangles []Triangle) {
	var t1, t2, t3 Triangle
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)
	for eof := true; eof; {
		var lines []string
		for i := 0; i < 3 && eof; i++ {
			eof = scanner.Scan()
			lines = append(lines, scanner.Text())
		}
		if !eof {
			break
		}
		fmt.Sscanf(lines[0], "  %3d  %3d  %3d", &t1.A, &t2.A, &t3.A)
		fmt.Sscanf(lines[1], "  %3d  %3d  %3d", &t1.B, &t2.B, &t3.B)
		fmt.Sscanf(lines[2], "  %3d  %3d  %3d", &t1.C, &t2.C, &t3.C)
		triangles = append(triangles, t1, t2, t3)
	}
	return
}

func Solve() (part1 int, part2 int) {
	for _, t := range parseInput() {
		if t.IsPossible() {
			part1++
		}
	}

	for _, t := range parseInputVertically() {
		if t.IsPossible() {
			part2++
		}
	}

	return
}
