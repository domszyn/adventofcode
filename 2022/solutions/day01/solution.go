package day01

import (
	"strconv"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func Solve() (int, int) {
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

	largest := utils.NLargest(3, calories)

	return largest[0], utils.Sum(largest)
}
