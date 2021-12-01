package solutions

import (
	"fmt"
	"time"

	"github.com/domszyn/adventofcode/2015/solutions/day1"
	"github.com/domszyn/adventofcode/2015/solutions/day2"
	"github.com/domszyn/adventofcode/2015/solutions/day3"
	"github.com/domszyn/adventofcode/2015/solutions/day4"
	"github.com/domszyn/adventofcode/2015/solutions/day5"
	"github.com/domszyn/adventofcode/2015/solutions/day6"
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
