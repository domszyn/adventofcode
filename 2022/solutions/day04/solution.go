package day04

import (
	"fmt"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

type Range struct{ From, To int }

func rangesOverlapFully(a, b Range) bool {
	return (a.From >= b.From && a.To <= b.To) || (b.From >= a.From && b.To <= a.To)
}

func rangesOverlapAtAll(a, b Range) bool {
	return (b.From >= a.From && b.From <= a.To) || (b.To >= a.From && b.To <= a.To) ||
		(a.From >= b.From && a.From <= b.To) || (a.To >= b.From && a.To <= b.To)
}

func Solve() (part1 int, part2 int) {
	lines := utils.ReadInput("./solutions/day04/input.txt", mappers.ToString)

	for _, line := range lines {
		ranges := [2]Range{{}, {}}

		fmt.Sscanf(line, "%d-%d,%d-%d", &ranges[0].From, &ranges[0].To, &ranges[1].From, &ranges[1].To)

		if rangesOverlapFully(ranges[0], ranges[1]) {
			part1++
		}

		if rangesOverlapAtAll(ranges[0], ranges[1]) {
			part2++
		}
	}

	return
}
