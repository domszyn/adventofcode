package day15

import (
	"bufio"
	"fmt"
	"math"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getNeighbors(location Location, points map[Location]Point) (neighbours []Location) {
	locations := []Location{
		{X: location.X - 1, Y: location.Y},
		{X: location.X + 1, Y: location.Y},
		{X: location.X, Y: location.Y - 1},
		{X: location.X, Y: location.Y + 1},
	}

	for _, loc := range locations {
		if _, found := points[loc]; found {
			neighbours = append(neighbours, loc)
		}
	}

	return
}

func getMinRisk(points map[Location]Point, goal Location) int {
	openSet := make(map[Location]Point)
	openSet[Location{X: 0, Y: 0}] = points[Location{X: 0, Y: 0}]

	gScore := make(map[Location]int)
	fScore := make(map[Location]int)
	for location := range points {
		gScore[location] = math.MaxInt64
		fScore[location] = math.MaxInt64
	}
	gScore[Location{X: 0, Y: 0}] = 0
	fScore[Location{X: 0, Y: 0}] = 0 // 2 * len(points)

	for len(openSet) > 0 {
		minFScore := math.MaxInt64
		var current Location
		for location := range openSet {
			if fScore[location] < minFScore {
				minFScore = fScore[location]
				current = location
			}
		}

		if current.X == goal.X && current.Y == goal.Y {
			break
		}

		delete(openSet, current)

		for _, neighbor := range getNeighbors(current, points) {
			tentativeGScore := gScore[current] + points[current].Risk
			if tentativeGScore < gScore[neighbor] {
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore //+ 2*len(points) - neighbor.X - neighbor.Y

				if _, found := openSet[neighbor]; !found {
					openSet[neighbor] = points[neighbor]
				}
			}
		}

		fmt.Printf("Open neighbors: %d\n", len(openSet))
	}

	return gScore[goal]
}

func SolvePart1() (part1 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var gridSize int
	points := make(map[Location]Point)
	for y := 0; scanner.Scan(); y++ {
		s := scanner.Text()
		gridSize = len(s)
		for x, v := range s {
			value := int(v - '0')
			points[Location{X: x, Y: y}] = Point{
				X:         x,
				Y:         y,
				Risk:      value,
				TotalRisk: 2*len(s) - x - y,
			}
		}
	}

	part1 = getMinRisk(points, Location{X: gridSize - 1, Y: gridSize - 1})

	return
}

func SolvePart2() (part1 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var gridSize int
	pattern := make(map[Location]Point)
	for y := 0; scanner.Scan(); y++ {
		s := scanner.Text()
		gridSize = len(s)
		for x, v := range s {
			value := int(v - '0')
			pattern[Location{X: x, Y: y}] = Point{
				X:         x,
				Y:         y,
				Risk:      value,
				TotalRisk: 2*len(s) - x - y,
			}
		}
	}

	points := make(map[Location]Point)

	for y := 0; y < gridSize*5; y++ {
		for x := 0; x < gridSize*5; x++ {
			offset := int(y/gridSize) + int(x/gridSize)
			risk := pattern[Location{X: x % gridSize, Y: y % gridSize}].Risk + offset
			if risk > 9 {
				risk -= 9
			}
			points[Location{X: x, Y: y}] = Point{
				X:         x,
				Y:         y,
				Risk:      risk,
				TotalRisk: 10*gridSize - x - y,
			}
		}
	}

	part1 = getMinRisk(points, Location{X: 5*gridSize - 1, Y: 5*gridSize - 1})

	return
}
