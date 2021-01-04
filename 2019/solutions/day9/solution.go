package day9

import (
	"strconv"
	"strings"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

func readInput() toolbox.Program {
	var ints []int
	for _, s := range strings.Split(Input, ",") {
		number, _ := strconv.Atoi(s)
		ints = append(ints, number)
	}

	return ints
}

func BOOST(i int) int {
	program := readInput()
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