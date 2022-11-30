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
	done := make(chan bool, 2)
	go program.IntCode(input, output, done)
	go program.IntCode(input, output, done)
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
	done := make(chan bool, 2)
	go program.IntCode(input, output, done)
	go program.IntCode(input, output, done)
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
	done := make(chan bool, 2)
	go program.IntCode(input, output, done)
	go program.IntCode(input, output, done)
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
	done := make(chan bool, 2)
	go program.IntCode(input, output, done)
	go program.IntCode(input, output, done)
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
	done := make(chan bool, 2)
	go program.IntCode(input, output, done)
	go program.IntCode(input, output, done)
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
	done := make(chan bool, 2)
	go program.IntCode(input, output, done)
	go program.IntCode(input, output, done)
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
	done := make(chan bool, 3)
	go program.IntCode(input, output, done)
	go program.IntCode(input, output, done)
	go program.IntCode(input, output, done)
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

func TestRelativeBase1(t *testing.T) {
	var program Program = []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}

	input := make(chan int)
	output := make(chan int, 16)
	done := make(chan bool, 1)
	go program.IntCode(input, output, done)
	<-done
	close(output)

	var result []int
	for v := range output {
		result = append(result, v)
	}

	if len(program) != len(result) {
		t.Errorf("IntCode(%#v) should return itself, got: %#v", program, program)
	}

	for i := 0; i < len(program); i++ {
		if program[i] != result[i] {
			t.Errorf("IntCode(%#v) should return itself, got: %#v", program, program)
			break
		}
	}
}

func TestRelativeBase2(t *testing.T) {
	var program Program = []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}

	input := make(chan int)
	output := make(chan int)
	done := make(chan bool, 1)
	go program.IntCode(input, output, done)
	select {
	case result := <-output:
		assertProgram(t, program, 0, result, 1219070632396864)
	}
}

func TestRelativeBase3(t *testing.T) {
	var program Program = []int{104, 1125899906842624, 99}

	input := make(chan int)
	output := make(chan int)
	done := make(chan bool, 1)
	go program.IntCode(input, output, done)
	select {
	case result := <-output:
		assertProgram(t, program, 0, result, 1125899906842624)
	}
}
