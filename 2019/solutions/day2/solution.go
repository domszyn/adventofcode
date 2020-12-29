package day2

import (
	"strconv"
	"strings"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

func readInput() []int {
	var ints []int
	for _, s := range strings.Split(Input, ",") {
		number, _ := strconv.Atoi(s)
		ints = append(ints, number)
	}

	return ints
}

func SolvePart1() int {
	return toolbox.IntCode(readInput(), []toolbox.Replacement{
		{Position: 1, Value: 12},
		{Position: 2, Value: 2},
	})
}

func SolvePart2() int {
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			result := toolbox.IntCode(readInput(), []toolbox.Replacement{
				{Position: 1, Value: n},
				{Position: 2, Value: v},
			})
			if result == 19690720 {
				return 100*n + v
			}
		}
	}

	return -1
}
