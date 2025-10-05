package day15

import (
	"math"
	"slices"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

type Point struct{ X, Y int }

const (
	Wall = iota
	Free
	Oxygen
)

const (
	_ = iota
	North
	South
	West
	East
)

type Droid struct {
	Program *toolbox.Program
	input   chan int
	output  chan int
	done    chan bool
}

func InitDroid(program *toolbox.Program) Droid {
	d := Droid{
		Program: program,
		input:   make(chan int),
		output:  make(chan int, 100),
		done:    make(chan bool, 1),
	}

	go d.Program.IntCode(d.input, d.output, d.done)

	return d
}

func (d Droid) Move(direction int) int {
	d.input <- direction
	// <-d.done

	return <-d.output
}

func getDirection(from, to Point) int {
	dx := to.X - from.X
	dy := to.Y - from.Y

	if dy == -1 && dx == 0 {
		return North
	}
	if dy == 1 && dx == 0 {
		return South
	}
	if dy == 0 && dx == -1 {
		return West
	}
	if dy == 0 && dx == 1 {
		return East
	}

	return -1
}

func reconstructPath(cameFrom map[Point]Point, current Point) []int {
	path := []Point{current}
	for {
		prev, ok := cameFrom[current]
		if !ok {
			break
		}

		path = append(path, prev)
		current = prev
	}
	slices.Reverse(path)

	var directions []int
	for i := 1; i < len(path); i++ {
		directions = append(directions, getDirection(path[i-1], path[i]))
	}

	return directions
}

func findShortestPath(from, to Point, knownMap map[Point]int) []int {
	openSet := make(map[Point]int)
	openSet[from] = knownMap[from]
	cameFrom := make(map[Point]Point)

	gScore := make(map[Point]int)
	fScore := make(map[Point]int)
	for location := range knownMap {
		if knownMap[location] == Wall {
			continue
		}
		gScore[location] = math.MaxInt64
		fScore[location] = math.MaxInt64
	}
	gScore[from] = 0
	fScore[from] = 0 // 2 * len(points)

	for len(openSet) > 0 {
		minFScore := math.MaxInt64
		var current Point
		for location := range openSet {
			if fScore[location] < minFScore {
				minFScore = fScore[location]
				current = location
			}
		}

		if current == to {
			break
		}

		delete(openSet, current)

		neighbors := []Point{
			{current.X, current.Y - 1},
			{current.X, current.Y + 1},
			{current.X - 1, current.Y},
			{current.X + 1, current.Y},
		}

		for _, neighbor := range neighbors {
			if n, found := knownMap[neighbor]; !found || n == Wall {
				continue
			}

			tentativeGScore := gScore[current] + 1
			if tentativeGScore < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + int(math.Abs(float64(to.X-current.X))+math.Abs(float64(to.Y-current.Y)))

				if _, found := openSet[neighbor]; !found {
					openSet[neighbor] = knownMap[neighbor]
				}
			}
		}
	}

	return reconstructPath(cameFrom, to)
}

func oppositeDirection(direction int) int {
	switch direction {
	case North:
		return South
	case South:
		return North
	case West:
		return East
	case East:
		return West
	}

	return direction
}

func (d Droid) searchOxygen() (int, Point, map[Point]int) {
	visited := make(map[Point]bool)
	var queue []Point
	spaceMap := make(map[Point]int)
	parentMap := make(map[Point]Point)
	currentPos := Point{0, 0}

	spaceMap[currentPos] = Free
	queue = append(queue, currentPos)
	visited[currentPos] = true

	var oxygenSteps int
	var oxygenPos Point

	for len(queue) > 0 {
		targetPos := queue[0]
		queue = queue[1:]

		if currentPos != targetPos {
			directions := findShortestPath(currentPos, targetPos, spaceMap)
			for _, dir := range directions {
				d.Move(dir)
				switch dir {
				case North:
					currentPos.Y -= 1
				case South:
					currentPos.Y += 1
				case West:
					currentPos.X -= 1
				case East:
					currentPos.X += 1
				}
			}
		}

		neighbors := map[Point]int{
			{targetPos.X, targetPos.Y - 1}: North,
			{targetPos.X, targetPos.Y + 1}: South,
			{targetPos.X - 1, targetPos.Y}: West,
			{targetPos.X + 1, targetPos.Y}: East,
		}

		for neighborPos, neighborDir := range neighbors {
			if visited[neighborPos] {
				continue
			}

			status := d.Move(neighborDir)
			switch status {
			case Wall:
				spaceMap[neighborPos] = Wall
			case Free:
				visited[neighborPos] = true
				spaceMap[neighborPos] = Free
				parentMap[neighborPos] = targetPos
				queue = append(queue, neighborPos)
				d.Move(oppositeDirection(neighborDir))
			case Oxygen:
				oxygenPos = targetPos
				visited[neighborPos] = true
				spaceMap[neighborPos] = Oxygen
				parentMap[neighborPos] = targetPos

				tmpPos := neighborPos
				start := Point{0, 0}
				for tmpPos != start {
					tmpPos = parentMap[tmpPos]
					oxygenSteps++
				}

				queue = append(queue, neighborPos)
				d.Move(oppositeDirection(neighborDir))
			}
		}
	}

	return oxygenSteps, oxygenPos, spaceMap
}

func GetAnswers() (int, int) {
	program := toolbox.LoadProgram(Input)
	droid := InitDroid(program)
	steps, oxygenPos, spaceMap := droid.searchOxygen()

	queue := map[Point]bool{
		oxygenPos: true,
	}
	var time int
	for time = 0; len(queue) > 0; time++ {
		for pos := range queue {
			spaceMap[pos] = Oxygen
		}

		var freeCells []Point
		for current := range queue {
			neighbors := []Point{
				{current.X, current.Y - 1},
				{current.X, current.Y + 1},
				{current.X - 1, current.Y},
				{current.X + 1, current.Y},
			}

			for _, n := range neighbors {
				if spaceMap[n] == Free {
					freeCells = append(freeCells, n)
				}
			}
		}

		queue = make(map[Point]bool)
		for _, pos := range freeCells {
			queue[pos] = true
		}
	}

	return steps, time
}
