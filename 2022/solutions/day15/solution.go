package day15

import (
	"fmt"
	"sort"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func distance(a, b Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

type Point struct{ X, Y int }

type Range struct{ Min, Max int }

type Space struct {
	Sensor, Beacon Point
	Distance       int
}

func beaconsImpossible(space []Space, y int) (Ranges, map[int]bool) {
	var ranges Ranges

	for _, s := range space {
		dy := abs(s.Sensor.Y - y)
		dx := s.Distance - dy
		if dx < 0 {
			continue
		}
		ranges = append(ranges, Range{
			s.Sensor.X - dx,
			s.Sensor.X + dx})
	}

	sort.Sort(ranges)

	for {
		hasMerged := false
		var tmp Ranges
		for i := 1; i < len(ranges); i++ {
			min1, min2 := ranges[i-1].Min, ranges[i].Min
			max1, max2 := ranges[i-1].Max, ranges[i].Max

			if min2 > max1 {
				tmp = append(tmp, ranges[i-1])
				continue
			}

			tmp = append(tmp, Range{utils.Min([]int{min1, min2}), utils.Max([]int{max1, max2})})
			tmp = append(tmp, ranges[i+1:]...)
			ranges = tmp
			hasMerged = true
			break
		}

		if !hasMerged {
			break
		}
	}

	beaconsAtY := make(map[int]bool)

	for _, s := range space {
		if s.Beacon.Y == y {
			beaconsAtY[s.Beacon.X] = true
		}
	}

	return ranges, beaconsAtY
}

func Solve() (part1 int, part2 int) {
	lines := utils.ReadInput("./solutions/day15/input.txt", mappers.ToString)
	var space []Space

	for _, v := range lines {
		var sx, sy, bx, by int
		fmt.Sscanf(v, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensor := Point{sx, sy}
		beacon := Point{bx, by}
		space = append(space, Space{
			sensor,
			beacon,
			distance(sensor, beacon),
		})
	}

	rr, beaconsAt2000000 := beaconsImpossible(space, 2000000)

	for _, r := range rr {
		part1 += r.Max - r.Min + 1
	}
	part1 -= len(beaconsAt2000000)

	maxY := 4000000
	for y := 0; y < maxY; y++ {
		r, _ := beaconsImpossible(space, y)
		if len(r) == 1 && r[0].Min < 0 && r[0].Max >= maxY {
			continue
		}

		if len(r) == 2 {
			part2 = maxY*(r[0].Max+1) + y
			break
		}
	}

	return
}

type Ranges []Range

func (a Ranges) Len() int           { return len(a) }
func (a Ranges) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Ranges) Less(i, j int) bool { return a[i].Min < a[j].Min }
