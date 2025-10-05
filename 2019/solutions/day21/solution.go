package day21

import (
	"fmt"
	"strings"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

func solvePart1() int {
	springscript := `NOT T T
AND A T
AND B T
AND C T
NOT T J
AND D J
WALK
`

	input := make(chan int)
	output := make(chan int)
	done := make(chan bool, 1)
	program := toolbox.LoadProgram(Input)
	go program.IntCode(input, output, done)

	inputIdx := 0

	var sb strings.Builder
loop:
	for {
		var inputChan chan int
		var nextInput int

		if inputIdx < len(springscript) {
			inputChan = input
			nextInput = int(springscript[inputIdx])
		}

		select {
		case inputChan <- nextInput:
			inputIdx++
		case val := <-output:
			if val > 256 {
				fmt.Println(sb.String())

				return val
			}

			sb.WriteRune(rune(val))
		case <-done:
			break loop
		}
	}

	fmt.Println(sb.String())
	return -1
}

func solvePart2() int {
	springscript := `NOT T T
AND A T
AND B T
AND C T
NOT T J
AND D J
NOT J T
AND J T
OR E T
OR H T
AND T J
RUN
`

	input := make(chan int)
	output := make(chan int)
	done := make(chan bool, 1)
	program := toolbox.LoadProgram(Input)
	go program.IntCode(input, output, done)

	inputIdx := 0

	var sb strings.Builder
loop:
	for {
		var inputChan chan int
		var nextInput int

		if inputIdx < len(springscript) {
			inputChan = input
			nextInput = int(springscript[inputIdx])
		}

		select {
		case inputChan <- nextInput:
			inputIdx++
		case val := <-output:
			if val > 256 {
				fmt.Println(sb.String())

				return val
			}

			sb.WriteRune(rune(val))
		case <-done:
			break loop
		}
	}

	fmt.Println(sb.String())
	return -1
}

func GetAnswers() (int, int) {
	return solvePart1(), solvePart2()
}
