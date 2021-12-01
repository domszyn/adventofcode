package solutions

import (
	"fmt"
	"time"

	"github.com/domszyn/adventofcode/2016/solutions/day1"
	"github.com/domszyn/adventofcode/2016/solutions/day10"
	"github.com/domszyn/adventofcode/2016/solutions/day12"
	"github.com/domszyn/adventofcode/2016/solutions/day2"
	"github.com/domszyn/adventofcode/2016/solutions/day3"
	"github.com/domszyn/adventofcode/2016/solutions/day4"
	"github.com/domszyn/adventofcode/2016/solutions/day5"
	"github.com/domszyn/adventofcode/2016/solutions/day6"
	"github.com/domszyn/adventofcode/2016/solutions/day7"
	"github.com/domszyn/adventofcode/2016/solutions/day8"
	"github.com/domszyn/adventofcode/2016/solutions/day9"
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

func (s *Solution) SolveDay6() {
	start := time.Now()
	s.Part1, s.Part2 = day6.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay7() {
	start := time.Now()
	s.Part1, s.Part2 = day7.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay8() {
	start := time.Now()
	s.Part1 = day8.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay9() {
	start := time.Now()
	s.Part1, s.Part2 = day9.Solve(), day9.Solve2()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay10() {
	start := time.Now()
	s.Part1, s.Part2 = day10.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay12() {
	start := time.Now()
	s.Part1, s.Part2 = day12.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}
