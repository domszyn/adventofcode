package day3

import (
	"math"
	"strconv"
	"strings"
)

func readInput() []string {
	return strings.Split(Input, "\n")
}

type Point struct {
	X, Y int
}

type Segment struct {
	Start, End Point
}

func getLinePoints(line string) []Point {
	points := []Point{{X: 0, Y: 0}}
	paths := strings.Split(line, ",")

	for _, path := range paths {
		direction := path[:1]
		length, _ := strconv.Atoi(path[1:])
		prevPoint := points[len(points)-1]
		switch direction {
		case "L":
			points = append(points, Point{X: prevPoint.X - length, Y: prevPoint.Y})
		case "R":
			points = append(points, Point{X: prevPoint.X + length, Y: prevPoint.Y})
		case "U":
			points = append(points, Point{X: prevPoint.X, Y: prevPoint.Y + length})
		case "D":
			points = append(points, Point{X: prevPoint.X, Y: prevPoint.Y - length})
		}
	}

	return points
}

func getIntersection(a, b Segment) Point {
	a1 := float64(a.End.X - a.Start.X)
	a2 := float64(a.End.Y - a.Start.Y)
	b1 := float64(b.Start.X - b.End.X)
	b2 := float64(b.Start.Y - b.End.Y)
	c1 := float64(b.Start.X - a.Start.X)
	c2 := float64(b.Start.Y - a.Start.Y)

	if a1*b2 == a2*b1 {
		return Point{}
	}

	s := float64(c1*b2-c2*b1) / float64(a1*b2-a2*b1)
	t := float64(a1*c2-a2*c1) / float64(a1*b2-a2*b1)

	if s >= 0 && t >= 0 && s <= 1 && t <= 1 {
		return Point{
			X: int(math.Round(float64(a.Start.X) + s*a1)),
			Y: int(math.Round(float64(a.Start.Y) + s*a2)),
		}
	}

	return Point{}
}

func findIntersections(lineA, lineB []Point) []Point {
	var intersections []Point
	for i := 0; i < len(lineA)-1; i++ {
		for j := 0; j < len(lineB)-1; j++ {
			intersection := getIntersection(Segment{Start: lineA[i], End: lineA[i+1]}, Segment{Start: lineB[j], End: lineB[j+1]})
			if intersection.X != 0 && intersection.Y != 0 {
				intersections = append(intersections, intersection)
			}
		}
	}

	return intersections
}

func (p *Point) distanceTo() int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
}

func (p *Point) stepsTo(line []Point) int {
	steps := 0
	for i := 0; i < len(line)-1; i++ {
		point1 := line[i]
		point2 := line[i+1]

		if p.X == point1.X && p.X == point2.X {
			steps += int(math.Abs(float64(p.Y - point1.Y)))
			break
		}

		if p.Y == point1.Y && p.Y == point2.Y {
			steps += int(math.Abs(float64(p.X - point1.X)))
			break
		}

		steps += int(math.Abs(float64(point2.X-point1.X)) + math.Abs(float64(point2.Y-point1.Y)))
	}

	return steps
}

func GetAnswers() (closestIntersection, minSteps int) {
	input := readInput()
	lineA := getLinePoints(input[0])
	lineB := getLinePoints(input[1])
	intersections := findIntersections(lineA, lineB)
	closestIntersection = math.MaxInt32
	minSteps = math.MaxInt32

	for _, intersection := range intersections {
		if d := intersection.distanceTo(); d < closestIntersection {
			closestIntersection = d
		}

		steps := intersection.stepsTo(lineA) + intersection.stepsTo(lineB)

		if steps < minSteps {
			minSteps = steps
		}
	}

	return
}
