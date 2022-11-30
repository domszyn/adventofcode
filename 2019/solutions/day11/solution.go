package day11

import "github.com/domszyn/adventofcode/2019/toolbox"

type HullPaintingRobot struct {
	Direction Direction
	Position  Position
	Hull      map[Position]int
}

func paintHull(startingColor int) map[Position]int {
	paintProgram := toolbox.LoadProgram(Input)
	robot := HullPaintingRobot{
		Direction: 0,
		Position:  Position{X: 0, Y: 0},
		Hull:      make(map[Position]int),
	}
	robot.Hull[robot.Position] = startingColor
	input := make(chan int)
	output := make(chan int)
	done := make(chan bool, 1)

	go func(input, output chan int, done chan bool) {
		paintProgram.IntCode(input, output, make(chan bool, 1))
		done <- true
	}(input, output, done)

	for {
		var paintComplete bool
		color := robot.Hull[robot.Position]

		select {
		case paintComplete = <-done:
		case input <- color:
		}

		if paintComplete {
			break
		}

		select {
		case color := <-output:
			robot.Hull[robot.Position] = color
		}

		select {
		case direction := <-output:
			if direction == 0 {
				robot.Direction = robot.Direction.RotateLeft()
			} else {
				robot.Direction = robot.Direction.RotateRight()
			}
		}

		switch robot.Direction {
		case 0:
			robot.Position = robot.Position.Up()
		case 90:
			robot.Position = robot.Position.Right()
		case 180:
			robot.Position = robot.Position.Down()
		case 270:
			robot.Position = robot.Position.Left()
		}
	}

	return robot.Hull
}

func GetAnswers() (answer1 int, identifier string) {
	answer1 = len(paintHull(0))
	hull := paintHull(1)
	maxPos := Position{X: 0, Y: 0}
	for pos := range hull {
		if pos.X > maxPos.X {
			maxPos.X = pos.X
		}
		if pos.Y > maxPos.Y {
			maxPos.Y = pos.Y
		}
	}

	identifier = "\n"
	for y := 0; y <= maxPos.Y; y++ {
		for x := 0; x <= maxPos.X; x++ {
			if hull[Position{X: x, Y: y}] == 0 {
				identifier += " "
			} else {
				identifier += "#"
			}
		}
		identifier += "\n"
	}
	return
}
