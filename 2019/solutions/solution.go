package solutions

import (
	"fmt"
	"time"

	"github.com/domszyn/adventofcode/2019/solutions/day1"
	"github.com/domszyn/adventofcode/2019/solutions/day10"
	"github.com/domszyn/adventofcode/2019/solutions/day11"
	"github.com/domszyn/adventofcode/2019/solutions/day12"
	"github.com/domszyn/adventofcode/2019/solutions/day13"
	"github.com/domszyn/adventofcode/2019/solutions/day14"
	"github.com/domszyn/adventofcode/2019/solutions/day15"
	"github.com/domszyn/adventofcode/2019/solutions/day16"
	"github.com/domszyn/adventofcode/2019/solutions/day17"
	"github.com/domszyn/adventofcode/2019/solutions/day18"
	"github.com/domszyn/adventofcode/2019/solutions/day19"
	"github.com/domszyn/adventofcode/2019/solutions/day2"
	"github.com/domszyn/adventofcode/2019/solutions/day20"
	"github.com/domszyn/adventofcode/2019/solutions/day21"
	"github.com/domszyn/adventofcode/2019/solutions/day22"
	"github.com/domszyn/adventofcode/2019/solutions/day23"
	"github.com/domszyn/adventofcode/2019/solutions/day24"
	"github.com/domszyn/adventofcode/2019/solutions/day25"
	"github.com/domszyn/adventofcode/2019/solutions/day3"
	"github.com/domszyn/adventofcode/2019/solutions/day4"
	"github.com/domszyn/adventofcode/2019/solutions/day5"
	"github.com/domszyn/adventofcode/2019/solutions/day6"
	"github.com/domszyn/adventofcode/2019/solutions/day7"
	"github.com/domszyn/adventofcode/2019/solutions/day8"
	"github.com/domszyn/adventofcode/2019/solutions/day9"
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
	s.Part1, s.Part2 = day3.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay4() {
	start := time.Now()
	s.Part1, s.Part2 = day4.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay5() {
	start := time.Now()
	s.Part1, s.Part2 = day5.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay6() {
	start := time.Now()
	s.Part1, s.Part2 = day6.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay7() {
	start := time.Now()
	s.Part1, s.Part2 = day7.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay8() {
	start := time.Now()
	s.Part1, s.Part2 = day8.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay9() {
	start := time.Now()
	s.Part1, s.Part2 = day9.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay10() {
	start := time.Now()
	s.Part1, s.Part2 = day10.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay11() {
	start := time.Now()
	s.Part1, s.Part2 = day11.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay12() {
	start := time.Now()
	s.Part1, s.Part2 = day12.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay13() {
	start := time.Now()
	s.Part1, s.Part2 = day13.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay14() {
	start := time.Now()
	s.Part1, s.Part2 = day14.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay15() {
	start := time.Now()
	s.Part1, s.Part2 = day15.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay16() {
	start := time.Now()
	s.Part1, s.Part2 = day16.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay17() {
	start := time.Now()
	s.Part1, s.Part2 = day17.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay18() {
	start := time.Now()
	s.Part1, s.Part2 = day18.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay19() {
	start := time.Now()
	s.Part1, s.Part2 = day19.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay20() {
	start := time.Now()
	s.Part1, s.Part2 = day20.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay21() {
	start := time.Now()
	s.Part1, s.Part2 = day21.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay22() {
	start := time.Now()
	s.Part1, s.Part2 = day22.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay23() {
	start := time.Now()
	s.Part1, s.Part2 = day23.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay24() {
	start := time.Now()
	s.Part1, s.Part2 = day24.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay25() {
	start := time.Now()
	s.Part1, s.Part2 = day25.GetAnswers()
	s.ExecTime = time.Since(start)
	s.Print()
}
