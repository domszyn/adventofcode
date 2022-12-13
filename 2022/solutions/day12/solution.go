package day12

import (
	"math"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

type Point struct {
	X int
	Y int
}

type HeightMap = map[Point]int

func getNeighbors(point Point, heightMap HeightMap) (neighbours []Point) {
	points := []Point{
		{point.X - 1, point.Y},
		{point.X + 1, point.Y},
		{point.X, point.Y - 1},
		{point.X, point.Y + 1},
	}

	for _, p := range points {
		if _, found := heightMap[p]; found {
			neighbours = append(neighbours, p)
		}
	}

	return
}

func findPath(points map[Point]int, start, end Point) int {
	openSet := make(map[Point]int)
	openSet[start] = points[start]

	gScore := make(map[Point]int)
	fScore := make(map[Point]int)
	for p := range points {
		gScore[p] = math.MaxInt64
		fScore[p] = math.MaxInt64
	}
	gScore[start] = 0
	fScore[start] = 0

	for len(openSet) > 0 {
		minFScore := math.MaxInt64
		var current Point
		for p := range openSet {
			if fScore[p] < minFScore {
				minFScore = fScore[p]
				current = p
			}
		}

		if current.X == end.X && current.Y == end.Y {
			break
		}

		delete(openSet, current)

		for _, neighbor := range getNeighbors(current, points) {
			if points[neighbor]-points[current] > 1 {
				continue
			}

			tentativeGScore := gScore[current] + 1
			if tentativeGScore < gScore[neighbor] {
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore

				if _, found := openSet[neighbor]; !found {
					openSet[neighbor] = points[neighbor]
				}
			}
		}
	}

	return gScore[end]
}

func Solve() (int, int) {
	lines := utils.ReadInput("./solutions/day12/input.txt", mappers.ToString)

	heightMap := make(HeightMap)
	var start, end Point

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		for j := 0; j < len(line); j++ {
			switch line[j] {
			case 'S':
				start = Point{j, i}
				heightMap[start] = 0
			case 'E':
				end = Point{j, i}
				heightMap[end] = 0
			default:
				heightMap[Point{j, i}] = int(line[j] - 'a' + 1)

			}
		}
	}

	part1 := findPath(heightMap, start, end)

	var steps []int

	for p, h := range heightMap {
		if h == 1 {
			steps = append(steps, findPath(heightMap, p, end))
		}
	}

	part2 := utils.Min(steps)

	return part1, part2
}
