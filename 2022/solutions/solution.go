package solutions

import (
	"fmt"
	"time"

	"github.com/domszyn/adventofcode/2022/solutions/day01"
	"github.com/domszyn/adventofcode/2022/solutions/day02"
	"github.com/domszyn/adventofcode/2022/solutions/day03"
	"github.com/domszyn/adventofcode/2022/solutions/day04"
	"github.com/domszyn/adventofcode/2022/solutions/day05"
	"github.com/domszyn/adventofcode/2022/solutions/day06"
	"github.com/domszyn/adventofcode/2022/solutions/day07"
	"github.com/domszyn/adventofcode/2022/solutions/day08"
	"github.com/domszyn/adventofcode/2022/solutions/day09"
	"github.com/domszyn/adventofcode/2022/solutions/day10"
	"github.com/domszyn/adventofcode/2022/solutions/day11"
	"github.com/domszyn/adventofcode/2022/solutions/day12"
	"github.com/domszyn/adventofcode/2022/solutions/day13"
	"github.com/domszyn/adventofcode/2022/solutions/day14"
	"github.com/domszyn/adventofcode/2022/solutions/day15"
	"github.com/domszyn/adventofcode/2022/solutions/day16"
	"github.com/domszyn/adventofcode/2022/solutions/day17"
	"github.com/domszyn/adventofcode/2022/solutions/day18"
	"github.com/domszyn/adventofcode/2022/solutions/day19"
	"github.com/domszyn/adventofcode/2022/solutions/day20"
	"github.com/domszyn/adventofcode/2022/solutions/day21"
	"github.com/domszyn/adventofcode/2022/solutions/day22"
	"github.com/domszyn/adventofcode/2022/solutions/day23"
	"github.com/domszyn/adventofcode/2022/solutions/day24"
	"github.com/domszyn/adventofcode/2022/solutions/day25"
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

func (s *Solution) SolveDay01() {
	start := time.Now()
	s.Part1, s.Part2 = day01.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay02() {
	start := time.Now()
	s.Part1, s.Part2 = day02.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay03() {
	start := time.Now()
	s.Part1, s.Part2 = day03.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay04() {
	start := time.Now()
	s.Part1, s.Part2 = day04.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay05() {
	start := time.Now()
	s.Part1, s.Part2 = day05.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay06() {
	start := time.Now()
	s.Part1, s.Part2 = day06.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay07() {
	start := time.Now()
	s.Part1, s.Part2 = day07.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay08() {
	start := time.Now()
	s.Part1, s.Part2 = day08.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay09() {
	start := time.Now()
	s.Part1, s.Part2 = day09.Solve()
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
	s.Part1, s.Part2 = day11.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
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
	s.Part1, s.Part2 = day15.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay16() {
	start := time.Now()
	s.Part1, s.Part2 = day16.Solve()
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
	s.Part1, s.Part2 = day20.Solve()
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
	s.Part1, s.Part2 = day22.Solve()
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
	s.Part1, s.Part2 = day24.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}

func (s *Solution) SolveDay25() {
	start := time.Now()
	s.Part1, s.Part2 = day25.Solve()
	s.ExecTime = time.Since(start)
	s.Print()
}
