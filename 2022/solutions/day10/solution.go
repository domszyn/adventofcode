package day10

import (
	"fmt"
	"strings"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func signalStrength(numbers []int, numCycles int) int {
	return utils.Sum(numbers[:numCycles-1]) + 1
}

func Solve() (part1 int, part2 string) {
	lines := utils.ReadInput("./solutions/day10/input.txt", mappers.ToString)

	var numbers []int
	for _, v := range lines {
		if v == "noop" {
			numbers = append(numbers, 0)
			continue
		}

		var n int
		fmt.Sscanf(v, "addx %d", &n)
		numbers = append(numbers, 0)
		numbers = append(numbers, n)
	}

	for i := 20; i <= 220; i += 40 {
		part1 += i * signalStrength(numbers, i)
	}

	var output []string

	for i := 0; i < 6; i++ {
		output = append(output, strings.Repeat(".", 40))
	}

	for i := 0; i < 240; i++ {
		x := i % 40
		y := i / 40

		s := signalStrength(numbers, i+1)
		if s >= x-1 && s <= x+1 {
			output[y] = output[y][:x] + "#" + output[y][x+1:]
		}
	}

	part2 = "\n\n" + strings.Join(output, "\n")

	return
}
