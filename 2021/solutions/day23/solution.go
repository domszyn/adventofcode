package day23

import (
	"bufio"
	"strings"
)

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	part1 = len(rows)
	return
}
