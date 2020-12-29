package solutions

import (
	"fmt"
	"time"

	"github.com/domszyn/adventofcode/2019/solutions/day1"
	"github.com/domszyn/adventofcode/2019/solutions/day2"
	"github.com/domszyn/adventofcode/2019/solutions/day3"
)

type Solution struct {
	Part1    interface{}
	Part2    interface{}
	ExecTime time.Duration
}

func (s *Solution) Print() {
	fmt.Printf("% +v", *s)
	fmt.Println()
}

func (s *Solution) SolveDay1() {
	start := time.Now()
	s.Part1 = day1.FuelRequirement(false)
	s.Part2 = day1.FuelRequirement(true)
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay2() {
	start := time.Now()

	s.Part1 = day2.SolvePart1()
	s.Part2 = day2.SolvePart2()

	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay3() {
	start := time.Now()
	answers := day3.GetAnswers()

	s.Part1 = answers[0]
	s.Part2 = answers[1]

	s.ExecTime = time.Since(start)
	s.Print()
}
