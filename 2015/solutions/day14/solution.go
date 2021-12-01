package day14

import (
	"bufio"
	"fmt"
	"math"
	"strings"
)

type Reindeer struct {
	Name        string
	Speed       int
	FlyingTime  int
	RestingTime int
}

func parseInput(s string) (deer Reindeer) {
	fmt.Sscanf(s, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &deer.Name, &deer.Speed, &deer.FlyingTime, &deer.RestingTime)
	return
}

func Solve(raceTime int) (int, int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var stable []Reindeer

	for scanner.Scan() {
		s := scanner.Text()
		deer := parseInput(s)
		stable = append(stable, deer)
	}

	race := make(map[string]int, len(stable))
	points := make(map[string]int, len(stable))

	for i := 1; i <= raceTime; i++ {
		for _, deer := range stable {
			race[deer.Name] = calculateDistance(i, deer)
		}

		maxDistance := findMax(race)
		for deer, distance := range race {
			if distance == maxDistance {
				points[deer]++
			}
		}
	}

	return findMax(race), findMax(points)
}

func calculateDistance(time int, deer Reindeer) int {
	cycleTime := deer.FlyingTime + deer.RestingTime
	numCycles := time / cycleTime
	remainingFlyingTime := int(math.Min(float64(deer.FlyingTime), float64(time%cycleTime)))
	return deer.Speed * (deer.FlyingTime*numCycles + remainingFlyingTime)
}

func findMax(race map[string]int) (maxDistance int) {
	for _, distance := range race {
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}
