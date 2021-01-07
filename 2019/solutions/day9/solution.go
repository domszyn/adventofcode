package day9

import (
	"github.com/domszyn/adventofcode/2019/toolbox"
)

func BOOST(i int) int {
	program := toolbox.LoadProgram(Input)
	input := make(chan int)
	output := make(chan int, 100)
	done := make(chan bool, 1)
	go program.IntCode(input, output, done)
	input <- i
	<-done

	return <-output
}

func GetAnswers() (int, int) {
	return BOOST(1), BOOST(2)
}
