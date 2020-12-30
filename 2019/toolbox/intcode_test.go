package toolbox

import (
	"testing"
)

func assertProgram(t *testing.T, program Program, input, output, expected int) {
	if output != expected {
		t.Errorf("IntCode(%#v) with input %d = %d; want %d", program, input, output, expected)
	}
}

func TestEqualTo8UsingPositionMode(t *testing.T) {
	var program Program = []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	input := make(chan int)
	output := make(chan int)
	go program.IntCode(input, output, nil)
	go program.IntCode(input, output, nil)
	input <- 8
	select {
	case result := <-output:
		assertProgram(t, program, 8, result, 1)
	}

	input <- 0
	select {
	case result := <-output:
		assertProgram(t, program, 0, result, 0)
	}
}

func TestLessThan8UsingPositionMode(t *testing.T) {
	var program Program = []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	input := make(chan int)
	output := make(chan int)
	go program.IntCode(input, output, nil)
	go program.IntCode(input, output, nil)
	input <- 8
	select {
	case result := <-output:
		assertProgram(t, program, 8, result, 0)
	}

	input <- 0
	select {
	case result := <-output:
		assertProgram(t, program, 0, result, 1)
	}
}

func TestEqualTo8UsingImmediateMode(t *testing.T) {
	var program Program = []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	input := make(chan int)
	output := make(chan int)
	go program.IntCode(input, output, nil)
	go program.IntCode(input, output, nil)
	input <- 8
	select {
	case result := <-output:
		assertProgram(t, program, 8, result, 1)
	}

	input <- 0
	select {
	case result := <-output:
		assertProgram(t, program, 0, result, 0)
	}
}

func TestLessThan8UsingImmediateMode(t *testing.T) {
	var program Program = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	input := make(chan int)
	output := make(chan int)
	go program.IntCode(input, output, nil)
	go program.IntCode(input, output, nil)
	input <- 8
	select {
	case result := <-output:
		assertProgram(t, program, 8, result, 0)
	}

	input <- 0
	select {
	case result := <-output:
		assertProgram(t, program, 0, result, 1)
	}
}

func TestJumpUsingPositionMode(t *testing.T) {
	var program Program = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	input := make(chan int)
	output := make(chan int)
	go program.IntCode(input, output, nil)
	go program.IntCode(input, output, nil)
	input <- 8
	select {
	case result := <-output:
		assertProgram(t, program, 8, result, 1)
	}

	input <- 0
	select {
	case result := <-output:
		assertProgram(t, program, 0, result, 0)
	}
}

func TestJumpUsingImmediateMode(t *testing.T) {
	var program Program = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	input := make(chan int)
	output := make(chan int)
	go program.IntCode(input, output, nil)
	go program.IntCode(input, output, nil)
	input <- 8
	select {
	case result := <-output:
		assertProgram(t, program, 8, result, 1)
	}

	input <- 0
	select {
	case result := <-output:
		assertProgram(t, program, 0, result, 0)
	}
}

func TestCombined(t *testing.T) {
	var program Program = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}

	input := make(chan int)
	output := make(chan int)
	go program.IntCode(input, output, nil)
	go program.IntCode(input, output, nil)
	go program.IntCode(input, output, nil)
	input <- 7
	select {
	case result := <-output:
		assertProgram(t, program, 7, result, 999)
	}

	input <- 8
	select {
	case result := <-output:
		assertProgram(t, program, 8, result, 1000)
	}

	input <- 9
	select {
	case result := <-output:
		assertProgram(t, program, 9, result, 1001)
	}
}
