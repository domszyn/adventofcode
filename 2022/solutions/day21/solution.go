package day21

import (
	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func Solve() (int, int) {
	lines := utils.ReadInput("./solutions/day21/input.txt", mappers.ToString)

	return len(lines), len(lines)
}
