package day5

import (
	"github.com/domszyn/adventofcode/2019/toolbox"
)

func GetAnswers() (part1 int, part2 int) {
	program := toolbox.LoadProgram(Input)

	input := make(chan int, 2)
	input <- 1
	input <- 5
	outputPart1 := make(chan int)
	outputPart2 := make(chan int)
	go program.IntCode(input, outputPart1, make(chan bool, 1))
	go program.IntCode(input, outputPart2, make(chan bool, 1))

	for v := range outputPart1 {
		if v != 0 {
			part1 = v
			break
		}
	}

	for v := range outputPart2 {
		if v != 0 {
			part2 = v
			break
		}
	}

	return
}
