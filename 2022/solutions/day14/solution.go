package day14

import (
	"fmt"
	"strings"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

const (
	Rock  = 1
	Sand  = 2
	Floor = 3
)

type Location struct{ X, Y int }

func Solve() (part1 int, part2 int) {
	lines := utils.ReadInput("./solutions/day14/input.txt", mappers.ToString)

	cave := make(map[Location]int)

	for _, v := range lines {
		locStrings := strings.Split(v, " -> ")

		var locs []Location
		for _, l := range locStrings {
			var x, y int
			fmt.Sscanf(l, "%d,%d", &x, &y)

			locs = append(locs, Location{x, y})
		}

		for i := 1; i < len(locs); i++ {
			a, b := locs[i-1], locs[i]

			if a.X == b.X {
				for y := a.Y; y <= b.Y; y++ {
					cave[Location{a.X, y}] = Rock
				}
				for y := b.Y; y <= a.Y; y++ {
					cave[Location{a.X, y}] = Rock
				}
			}

			if a.Y == b.Y {
				for x := a.X; x <= b.X; x++ {
					cave[Location{x, a.Y}] = Rock
				}
				for x := b.X; x <= a.X; x++ {
					cave[Location{x, a.Y}] = Rock
				}
			}
		}
	}

	var maxY int
	for rock := range cave {
		if rock.Y > maxY {
			maxY = rock.Y
		}
	}

	const x = 500
	const y = 0

	for i := 1; ; i++ {
		for xx, yy := x, y; ; {
			if _, ok := cave[Location{xx, yy + 1}]; !ok && yy < maxY+1 {
				yy++
				continue
			}
			if _, ok := cave[Location{xx - 1, yy + 1}]; !ok && yy < maxY+1 {
				xx--
				yy++
				continue
			}
			if _, ok := cave[Location{xx + 1, yy + 1}]; !ok && yy < maxY+1 {
				xx++
				yy++
				continue
			}

			if yy > maxY && part1 == 0 {
				part1 = i - 1
			}

			cave[Location{xx, yy}] = Sand
			break
		}

		if cave[Location{x, y}] == Sand {
			part2 = i
			break
		}
	}

	return part1, part2
}
