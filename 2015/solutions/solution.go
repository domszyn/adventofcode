package solutions

import (
	"fmt"
	"time"

	"github.com/domszyn/adventofcode/2015/solutions/day1"
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
