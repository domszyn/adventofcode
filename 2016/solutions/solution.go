package solutions

import (
	"fmt"
	"time"

	"github.com/domszyn/adventofcode/2016/solutions/day1"
	"github.com/domszyn/adventofcode/2016/solutions/day2"
	"github.com/domszyn/adventofcode/2016/solutions/day3"
	"github.com/domszyn/adventofcode/2016/solutions/day4"
	"github.com/domszyn/adventofcode/2016/solutions/day5"
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
	s.Part1 = day2.Solve()
	s.Part2 = day2.SolvePart2()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay3() {
	start := time.Now()
	s.Part1, s.Part2 = day3.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay4() {
	start := time.Now()
	s.Part1, s.Part2 = day4.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay5() {
	start := time.Now()
	s.Part1, s.Part2 = day5.Solve(), day5.Solve2()
	s.ExecTime = time.Since(start)
	s.Print()
}
