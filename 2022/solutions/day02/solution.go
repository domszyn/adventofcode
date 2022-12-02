package day02

import (
	"strings"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

const (
	Rock     = "A"
	Paper    = "B"
	Scissors = "C"
)

func Solve() (part1 int, part2 int) {
	lines := utils.ReadInput("./solutions/day02/input.txt", mappers.ToString)

	var moves [][]string
	for _, l := range lines {
		moves = append(moves, strings.Split(l, " "))
	}

	for _, m := range moves {
		switch m[0] {
		case Rock:
			switch m[1] {
			case "X":
				part1 += 1
				part2 += 3
			case "Y":
				part1 += 8
				part2 += 4
			case "Z":
				part1 += 3
				part2 += 8
			}
		case Paper:
			switch m[1] {
			case "X":
				part1 += 1
				part2 += 1
			case "Y":
				part1 += 5
				part2 += 5
			case "Z":
				part1 += 9
				part2 += 9
			}
		case Scissors:
			switch m[1] {
			case "X":
				part1 += 7
				part2 += 2
			case "Y":
				part1 += 2
				part2 += 6
			case "Z":
				part1 += 6
				part2 += 7
			}
		}
	}

	return
}
