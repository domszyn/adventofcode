package solutions

import (
	"fmt"
	"time"

	"github.com/domszyn/adventofcode/2015/solutions/day1"
	"github.com/domszyn/adventofcode/2015/solutions/day10"
	"github.com/domszyn/adventofcode/2015/solutions/day11"
	"github.com/domszyn/adventofcode/2015/solutions/day12"
	"github.com/domszyn/adventofcode/2015/solutions/day13"
	"github.com/domszyn/adventofcode/2015/solutions/day14"
	"github.com/domszyn/adventofcode/2015/solutions/day15"
	"github.com/domszyn/adventofcode/2015/solutions/day16"
	"github.com/domszyn/adventofcode/2015/solutions/day17"
	"github.com/domszyn/adventofcode/2015/solutions/day18"
	"github.com/domszyn/adventofcode/2015/solutions/day19"
	"github.com/domszyn/adventofcode/2015/solutions/day2"
	"github.com/domszyn/adventofcode/2015/solutions/day20"
	"github.com/domszyn/adventofcode/2015/solutions/day21"
	"github.com/domszyn/adventofcode/2015/solutions/day22"
	"github.com/domszyn/adventofcode/2015/solutions/day23"
	"github.com/domszyn/adventofcode/2015/solutions/day24"
	"github.com/domszyn/adventofcode/2015/solutions/day3"
	"github.com/domszyn/adventofcode/2015/solutions/day4"
	"github.com/domszyn/adventofcode/2015/solutions/day5"
	"github.com/domszyn/adventofcode/2015/solutions/day6"
	"github.com/domszyn/adventofcode/2015/solutions/day7"
	"github.com/domszyn/adventofcode/2015/solutions/day8"
	"github.com/domszyn/adventofcode/2015/solutions/day9"
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
	s.Part1 = day1.CalculateFloor()
	s.Part2 = day1.CalculateBasementPosition()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay2() {
	start := time.Now()
	s.Part1, s.Part2 = day2.CalculatePaperAndRibbon()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay3() {
	start := time.Now()
	s.Part1 = day3.CountHouses()
	s.Part2 = day3.CountHousesWithRoboSanta()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay4() {
	start := time.Now()
	s.Part1 = day4.FindHashWithPrefix("00000")
	s.Part2 = day4.FindHashWithPrefix("000000")
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay5() {
	start := time.Now()
	s.Part1, s.Part2 = day5.CountNiceStrings()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay6() {
	start := time.Now()
	s.Part1 = day6.SolvePart1()
	s.Part2 = day6.SolvePart2()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay7() {
	start := time.Now()
	s.Part1 = day7.Solve(day7.Input)
	s.Part2 = day7.Solve(day7.Input2)
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

func (s *Solution) SolveDay11() {
	start := time.Now()
	pwd1 := day11.GetNewPassword("cqjxjnds")
	pwd2 := day11.GetNewPassword(pwd1)
	s.Part1 = pwd1
	s.Part2 = pwd2
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay12() {
	start := time.Now()
	s.Part1 = day12.SolvePart1()
	s.Part2 = day12.SolvePart2()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay13() {
	start := time.Now()
	s.Part1, s.Part2 = day13.Solve(false), day13.Solve(true)
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay14() {
	start := time.Now()
	s.Part1, s.Part2 = day14.Solve(2503)
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay15() {
	start := time.Now()
	s.Part1, s.Part2 = day15.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay16() {
	start := time.Now()
	s.Part1, s.Part2 = day16.SolvePart1(), day16.SolvePart2()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay17() {
	start := time.Now()
	s.Part1, s.Part2 = day17.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay18() {
	start := time.Now()
	s.Part1, s.Part2 = day18.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay19() {
	start := time.Now()
	s.Part1, s.Part2 = day19.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay20() {
	start := time.Now()
	s.Part1, s.Part2 = day20.SolvePart1(), day20.SolvePart2()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay21() {
	start := time.Now()
	s.Part1, s.Part2 = day21.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay22() {
	start := time.Now()
	s.Part1 = day22.SolvePart1()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay23() {
	start := time.Now()
	s.Part1, s.Part2 = day23.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay24() {
	start := time.Now()
	s.Part1, s.Part2 = day24.SolvePart1(), day24.SolvePart2()
	s.ExecTime = time.Since(start)
	s.Print()
}
