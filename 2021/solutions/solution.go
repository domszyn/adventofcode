package solutions

import (
	"fmt"
	"time"

	"github.com/domszyn/adventofcode/2021/solutions/day1"
	"github.com/domszyn/adventofcode/2021/solutions/day2"
	"github.com/domszyn/adventofcode/2021/solutions/day3"
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
	s.Part1 = day1.SolvePart1()
	s.Part2 = day1.SolvePart2()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay2() {
	start := time.Now()
	s.Part1, s.Part2 = day2.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay3() {
	start := time.Now()
	s.Part1, s.Part2 = day3.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}
