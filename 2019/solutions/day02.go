package solutions

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

func readInput() []int {
	input, err := ioutil.ReadFile("./input/day02.txt")
	if err != nil {
		log.Fatal(err)
	}

	var ints []int
	for _, s := range strings.Split(string(input), ",") {
		number, _ := strconv.Atoi(s)
		ints = append(ints, number)
	}

	return ints
}

func (s *Solution) SolveDay2() {
	start := time.Now()
	input := readInput()

	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			result := toolbox.IntCode(input, []toolbox.Replacement{
				{Position: 1, Value: n},
				{Position: 2, Value: v},
			})
			if n == 12 && v == 2 {
				s.Part1 = result
			}
			if result == 19690720 {
				s.Part2 = 100*n + v
			}

			if s.Part1 != nil && s.Part2 != nil {
				break
			}
		}

		if s.Part1 != nil && s.Part2 != nil {
			break
		}
	}
	s.ExecTime = time.Since(start)
}
