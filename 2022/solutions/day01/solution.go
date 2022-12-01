package day01

import (
	"sort"
	"strconv"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func Solve() (part1 int, part2 int) {
	lines := utils.ReadInput("./solutions/day01/input.txt", mappers.ToString)

	var calories []int
	calories = append(calories, 0)
	for _, v := range lines {
		if len(v) == 0 {
			calories = append(calories, 0)
		}

		c, _ := strconv.Atoi(v)
		calories[len(calories)-1] += c
	}

	part1 = utils.Max(calories)

	sort.Ints(calories)
	part2 = utils.Sum(calories[len(calories)-3:])

	return
}
