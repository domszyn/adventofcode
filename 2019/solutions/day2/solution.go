package day2

import (
	"strconv"
	"strings"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

func readInput() toolbox.Program {
	var ints []int
	for _, s := range strings.Split(Input, ",") {
		number, _ := strconv.Atoi(s)
		ints = append(ints, number)
	}

	return ints
}

func SolvePart1() int {
	program := readInput()

	patches := []toolbox.Replacement{
		{Position: 1, Value: 12},
		{Position: 2, Value: 2},
	}

	result := program.Patch(patches).IntCode(
		make(chan int),
		make(chan int),
		make(chan bool, 1))

	return result[0]
}

func SolvePart2() int {
	program := readInput()

	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			patches := []toolbox.Replacement{
				{Position: 1, Value: n},
				{Position: 2, Value: v},
			}

			result := program.Patch(patches).IntCode(
				make(chan int),
				make(chan int),
				make(chan bool, 1))

			if result[0] == 19690720 {
				return 100*n + v
			}
		}
	}

	return -1
}
