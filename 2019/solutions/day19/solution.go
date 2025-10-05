package day19

import (
	"math"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

func run(program *toolbox.Program, inputs [2]int) int {
	input := make(chan int)
	output := make(chan int)
	done := make(chan bool, 1)

	go program.IntCode(input, output, done)

	inputIdx := 0

loop:
	for {
		var inputChan chan int
		var nextInput int

		if inputIdx < len(inputs) {
			inputChan = input
			nextInput = inputs[inputIdx]
		}

		select {
		case inputChan <- nextInput:
			inputIdx++
		case val := <-output:
			return val
		case <-done:
			break loop
		}
	}

	return -1
}

func solvePart1() int {
	var affectedCount int
	program := toolbox.LoadProgram(Input)
	var beamGrid [][]int

	for i := 0; i < 50; i++ {
		beamGrid = append(beamGrid, make([]int, 50))
	}

	for i := range 50 {
		for j := range 50 {
			val := run(program, [2]int{i, j})
			beamGrid[j][i] = val
			affectedCount += val
		}
	}

	return affectedCount
}

type Point struct{ X, Y int }

func solvePart2() int {
	program := toolbox.LoadProgram(Input)
	sizeMap := make(map[Point]int)
	visited := make(map[Point]bool)
	frontier := toolbox.NewQueue[Point]()
	gridMap := make(map[Point]int)

	start := Point{14, 8}
	frontier.Enqueue(start)
	visited[start] = true
	sizeMap[start] = 1

	for frontier.Size() > 0 {
		current, err := frontier.Dequeue()
		if err != nil {
			return -1
		}

		neighbors := []Point{
			{current.X + 1, current.Y},
			{current.X, current.Y + 1},
		}

		for _, n := range neighbors {
			if visited[n] {
				continue
			}

			visited[n] = true
			if _, found := gridMap[n]; !found {
				gridMap[n] = run(program, [2]int{n.X, n.Y})
			}
			val := gridMap[n]

			if val == 1 {
				s1 := float64(sizeMap[Point{n.X - 1, n.Y}])
				s2 := float64(sizeMap[Point{n.X, n.Y - 1}])
				s3 := float64(sizeMap[Point{n.X - 1, n.Y - 1}])
				newSize := 1 + int(math.Min(math.Min(s1, s2), s3))
				if newSize >= 100 {
					// n.X n.Y is bottom right
					return 10_000*(n.X-99) + n.Y - 99
				}

				sizeMap[n] = newSize
				frontier.Enqueue(n)
			}
		}
	}

	return -1
}

func GetAnswers() (int, int) {
	return solvePart1(), solvePart2()
}
