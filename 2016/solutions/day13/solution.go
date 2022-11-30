package day13

import (
	"fmt"
	"math"
	"math/bits"
)

func isOpenSpace(x, y uint) bool {
	return bits.OnesCount(x*x+3*x+2*x*y+y+y*y+1352)%2 == 0
}

type Point struct {
	X, Y uint
}

type FloorTile struct {
	Distance uint
	Visited  bool
}

func findShortestDistance(floor map[Point]*FloorTile) (uint, Point) {
	var minDist uint = math.MaxUint
	var minPoint Point

	for point, tile := range floor {
		if !tile.Visited && tile.Distance > 0 && tile.Distance < minDist {
			minDist = tile.Distance
			minPoint = point
		}
	}

	return minDist, minPoint
}

func Solve() {
	floor := make(map[Point]*FloorTile)

	var x, y uint
	for x = 0; x < 50; x++ {
		for y = 0; y < 50; y++ {
			if isOpenSpace(x, y) {
				floor[Point{x, y}] = &FloorTile{Distance: math.MaxUint}
			}
		}
	}

	currentTile := Point{1, 1}
	floor[currentTile].Distance = 0

	for !floor[Point{31, 39}].Visited {
		neighbours := []*FloorTile{
			floor[Point{currentTile.X - 1, currentTile.Y}],
			floor[Point{currentTile.X + 1, currentTile.Y}],
			floor[Point{currentTile.X, currentTile.Y - 1}],
			floor[Point{currentTile.X, currentTile.Y + 1}],
		}

		for _, n := range neighbours {
			if n != nil && !n.Visited {
				n.Distance = floor[currentTile].Distance + 1
			}
		}

		floor[currentTile].Visited = true
		_, currentTile = findShortestDistance(floor)
	}

	fmt.Printf("%v\n", floor[Point{31, 39}].Distance)

	part2 := 0
	for _, tile := range floor {
		if tile.Distance <= 50 {
			part2++
		}
	}

	fmt.Printf("%d\n", part2)
}
