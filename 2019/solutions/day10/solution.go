package day10

import (
	"bufio"
	"math"
	"sort"
	"strings"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

type Point struct{ X, Y int }
type Target struct {
	Position Point
	Angle    float64
	Distance float64
}

func points(x1, y1, x2, y2 int) [][2]int {
	dx := x2 - x1
	dy := y2 - y1
	adx := dx
	ady := dy
	if adx < 0 {
		adx *= -1
	}
	if ady < 0 {
		ady *= -1
	}
	g := int(toolbox.Gcd(int64(adx), int64(ady)))
	if g == 1 {
		return make([][2]int, 0)
	}

	stepX := dx / g
	stepY := dy / g

	result := make([][2]int, g-1)

	for i := 0; i < g-1; i++ {
		result[i] = [2]int{y1 + stepY*(i+1), x1 + stepX*(i+1)}
	}

	return result
}

func detect(x, y int, grid [][]bool) int {
	detected := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == y && j == x {
				continue
			}

			if grid[i][j] {
				between := points(x, y, j, i)

				if len(between) == 0 {
					detected++
				} else {
					cleanPath := true
					for _, p := range between {
						if grid[p[0]][p[1]] {
							cleanPath = false
							break
						}
					}

					if cleanPath {
						detected++
					}
				}
			}
		}
	}

	return detected
}

func parseGrid(input string) ([][]bool, []Point) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var grid [][]bool
	var asteroids []Point
	for scanner.Scan() {
		s := scanner.Text()
		grid = append(grid, make([]bool, len(s)))

		for x, c := range s {
			if c == '#' {
				y := len(grid) - 1
				grid[y][x] = true
				asteroids = append(asteroids, Point{X: x, Y: y})
			}
		}
	}

	return grid, asteroids
}

func findMonitoringStationLocation(grid [][]bool, asteroids []Point) (Point, int) {
	maxDetected := 0
	location := Point{}

	for _, p := range asteroids {
		detected := detect(p.X, p.Y, grid)
		if detected > maxDetected {
			maxDetected = detected
			location = p
		}
	}

	return location, maxDetected
}

func vaporizeAsteroids(grid [][]bool, asteroids []Point) int {
	station, _ := findMonitoringStationLocation(grid, asteroids)

	var targets []Target
	for _, asteroid := range asteroids {
		if asteroid == station {
			continue
		}

		dx := float64(asteroid.X - station.X)
		dy := float64(asteroid.Y - station.Y)
		angle := math.Atan2(dx, -dy)
		if angle < 0 {
			angle += 2 * math.Pi
		}

		distance := dx*dx + dy*dy
		targets = append(targets, Target{
			Position: asteroid,
			Angle:    angle,
			Distance: distance,
		})
	}

	groupedByAngle := make(map[float64][]Target)
	for _, t := range targets {
		groupedByAngle[t.Angle] = append(groupedByAngle[t.Angle], t)
	}

	var sortedAngles []float64
	for angle := range groupedByAngle {
		sort.Slice(groupedByAngle[angle], func(i, j int) bool {
			return groupedByAngle[angle][i].Distance < groupedByAngle[angle][j].Distance
		})
		sortedAngles = append(sortedAngles, angle)
	}

	sort.Float64s(sortedAngles)

	vaporizedCount := 0
	for sweep := 0; ; sweep++ {
		for _, angle := range sortedAngles {
			targetsAtAngle := groupedByAngle[angle]
			if sweep < len(targetsAtAngle) {
				vaporizedCount++
				if vaporizedCount == 200 {
					// Found the 200th asteroid.
					winner := targetsAtAngle[sweep].Position
					return winner.X*100 + winner.Y
				}
			}
		}
	}
}

func GetAnswers() (int, int) {
	grid, asteroids := parseGrid(Input)
	_, detected := findMonitoringStationLocation(grid, asteroids)
	vaporized := vaporizeAsteroids(grid, asteroids)
	return detected, vaporized
}
