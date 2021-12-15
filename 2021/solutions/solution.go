package solutions

import (
	"fmt"
	"time"

	"github.com/domszyn/adventofcode/2021/solutions/day1"
	"github.com/domszyn/adventofcode/2021/solutions/day10"
	"github.com/domszyn/adventofcode/2021/solutions/day11"
	"github.com/domszyn/adventofcode/2021/solutions/day12"
	"github.com/domszyn/adventofcode/2021/solutions/day13"
	"github.com/domszyn/adventofcode/2021/solutions/day14"
	"github.com/domszyn/adventofcode/2021/solutions/day15"
	"github.com/domszyn/adventofcode/2021/solutions/day2"
	"github.com/domszyn/adventofcode/2021/solutions/day3"
	"github.com/domszyn/adventofcode/2021/solutions/day4"
	"github.com/domszyn/adventofcode/2021/solutions/day5"
	"github.com/domszyn/adventofcode/2021/solutions/day6"
	"github.com/domszyn/adventofcode/2021/solutions/day7"
	"github.com/domszyn/adventofcode/2021/solutions/day8"
	"github.com/domszyn/adventofcode/2021/solutions/day9"
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

func (s *Solution) SolveDay4() {
	start := time.Now()
	s.Part1, s.Part2 = day4.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay5() {
	start := time.Now()
	s.Part1, s.Part2 = day5.Solve(false), day5.Solve(true)
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
	s.Part1, s.Part2 = day8.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay9() {
	start := time.Now()
	s.Part1, s.Part2 = day9.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay10() {
	start := time.Now()
	s.Part1, s.Part2 = day10.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay11(animate bool) {
	start := time.Now()
	s.Part1, s.Part2 = day11.Solve(animate)
	s.ExecTime = time.Since(start)
	if !animate {
		s.Print()
	}
}

func (s *Solution) SolveDay12() {
	start := time.Now()
	s.Part1, s.Part2 = day12.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay13() {
	start := time.Now()
	s.Part1, s.Part2 = day13.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay14() {
	start := time.Now()
	s.Part1, s.Part2 = day14.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay15() {
	start := time.Now()
	s.Part1, s.Part2 = day15.SolvePart1(), day15.SolvePart2()
	s.ExecTime = time.Since(start)
	s.Print()
}
