package day9

import (
	"bufio"
	"fmt"
	"math"
	"strings"
)

func parseDistance(s string) (from, to string, distance int) {
	fmt.Sscanf(s, "%s to %s = %d", &from, &to, &distance)
	return
}

type Distance struct {
	From     string
	To       string
	Distance int
}

func Solve() (int, int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	cities := make(map[string]int)

	var distances []Distance

	for scanner.Scan() {
		s := scanner.Text()
		from, to, distance := parseDistance(s)
		distances = append(distances, Distance{
			From:     from,
			To:       to,
			Distance: distance,
		})
	}

	var counter int
	for _, distance := range distances {
		if _, ok := cities[distance.From]; !ok {
			cities[distance.From] = counter
			counter++
		}
		if _, ok := cities[distance.To]; !ok {
			cities[distance.To] = counter
			counter++
		}
	}

	grid := make([][]int, len(cities))
	for i := 0; i < len(cities); i++ {
		grid[i] = make([]int, len(cities))
	}

	for _, distance := range distances {
		from := cities[distance.From]
		to := cities[distance.To]
		grid[from][to] = distance.Distance
		grid[to][from] = distance.Distance
	}

	return findMinDistance(grid), findMaxDistance(grid)
}

func findMinDistance(grid [][]int) int {
	minDistance := math.MaxInt32
	for i := 0; i < len(grid); i++ {
		var route []int
		inRoute := make(map[int]bool)
		distance := 0

		from := i
		for len(route) < len(grid) {
			inRoute[from] = true
			route = append(route, from)
			to := findClosestCity(inRoute, grid[from])
			if to == -1 {
				break
			}
			distance += grid[from][to]
			from = to
		}

		if distance < minDistance {
			minDistance = distance
		}
	}

	return minDistance
}

func findMaxDistance(grid [][]int) int {
	maxDistance := 0
	for i := 0; i < len(grid); i++ {
		var route []int
		inRoute := make(map[int]bool)
		distance := 0

		from := i
		for len(route) < len(grid) {
			inRoute[from] = true
			route = append(route, from)
			to := findFarthestCity(inRoute, grid[from])
			if to == -1 {
				break
			}
			distance += grid[from][to]
			from = to
		}

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func findClosestCity(inRoute map[int]bool, distances []int) int {
	minDistance := math.MaxInt32
	nextCity := -1
	for i := 0; i < len(distances); i++ {
		if inRoute[i] || distances[i] == 0 {
			continue
		}

		if distances[i] < minDistance {
			minDistance = distances[i]
			nextCity = i
		}
	}

	return nextCity
}

func findFarthestCity(inRoute map[int]bool, distances []int) int {
	maxDistance := 0
	nextCity := -1
	for i := 0; i < len(distances); i++ {
		if inRoute[i] || distances[i] == 0 {
			continue
		}

		if distances[i] > maxDistance {
			maxDistance = distances[i]
			nextCity = i
		}
	}

	return nextCity
}
