package day13

import (
	"time"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

func SolvePart1() int {
	p := toolbox.LoadProgram(Input)
	input := make(chan int)
	output := make(chan int, 100)
	done := make(chan bool, 1)
	complete := false
	grid := InitGrid()

	go func(input, output chan int, done chan bool) {
		p.IntCode(input, output, make(chan bool, 1))
		done <- true
	}(input, output, done)

	for !complete {
		var x, y, id int
		select {
		case complete = <-done:
		default:
		}

		if complete {
			break
		}

		for i := 0; i < 3; i++ {
			select {
			case v := <-output:
				switch i {
				case 0:
					x = v
				case 1:
					y = v
				case 2:
					id = v
				}
			}
		}

		grid.Row(y)[x] = id
	}

	return grid.CountBlocks()
}

func SolvePart2(animate, fast bool) int {
	maxRounds := 1
	if !fast {
		maxRounds = 2000
	}

	inputs := make([]int, 10000)
	if fast {
		copy(inputs, Part2Inputs)
	}

	prevUsed := 0
	var score Score

	for r := 0; r < maxRounds; r++ {
		if animate {
			printRound(r)
		}

		p := toolbox.LoadProgram(Input).Patch([]toolbox.Replacement{{Position: 0, Value: 2}})
		input := make(chan int, 10000)
		output := make(chan int)
		done := make(chan bool, 1)
		complete := false
		grid := InitGrid()

		go func(input, output chan int, done chan bool) {
			p.IntCode(input, output, make(chan bool, 1))
			done <- true
		}(input, output, done)

		for _, i := range inputs {
			input <- i
		}

		for !complete {
			if animate {
				score.Print()
			}

			var x, y, id int
			select {
			case complete = <-done:
			default:
			}

			if complete {
				break
			}

			for i := 0; i < 3; i++ {
				select {
				case v := <-output:
					switch i {
					case 0:
						x = v
					case 1:
						y = v
					case 2:
						id = v
					}
				}
			}

			if x == -1 && y == 0 {
				if id > 0 {
					score.Set(id)
				}
			} else {
				grid.Row(y)[x] = id
				if animate {
					printTile(x, y, id)
				}
			}

			if animate {
				time.Sleep(time.Millisecond * 10)
			}
		}

		close(input)
		used := 10000 - len(input)

		if grid.CountBlocks() > 0 {
			if animate {
				printGameOver()
			}

			if !fast {
				time.Sleep(time.Millisecond * 500)
			}
		} else {
			if animate {
				printWin()
			}
			if !fast {
				time.Sleep(time.Second * 2)
			}
			break
		}

		ballX, _ := grid.FindBall()
		paddleX, _ := grid.FindPaddle()
		offset := 0
		if score.Endgame() {
			offset = 1
		}

		if used == prevUsed {
			if paddleX == -1 {
				paddleX = ballX + 1
			}

			for i := used - 1; i > 0; i-- {
				if inputs[i] != 0 {
					for j := 0; j < paddleX-ballX+offset; j++ {
						inputs[i-j] = -1
					}
					for j := 0; j < ballX-paddleX+offset; j++ {
						inputs[i-j] = 1
					}
					break
				}
			}
		} else {
			prevUsed = used

			if paddleX > ballX {
				for i := 1; i < paddleX-ballX+offset; i++ {
					inputs[used-i] = -1
				}
			} else if ballX > paddleX {
				for i := 1; i < ballX-paddleX+offset; i++ {
					inputs[used-i] = 1
				}
			}
		}
	}

	return score.Score
}

func GetAnswers() (int, int) {
	return SolvePart1(), SolvePart2(false, true)
}
