package day2

import (
	"bufio"
	"fmt"
	"strings"
)

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var x, y1, y2, aim int
	for scanner.Scan() {
		s := scanner.Text()
		var direction string
		var count int
		fmt.Sscanf(s, "%s%d", &direction, &count)

		switch direction {
		case "forward":
			x += count
			y2 += count * aim
		case "up":
			y1 -= count
			aim -= count
		case "down":
			y1 += count
			aim += count
		}
	}

	part1 = x * y1
	part2 = x * y2

	return
}
