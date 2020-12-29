package toolbox

import "testing"

func assertProgram(t *testing.T, program Program, input, output, expected int) {
	if output != expected {
		t.Errorf("IntCode(%#v) with input %d = %d; want %d", program, input, output, expected)
	}
}

func TestEqualTo8UsingPositionMode(t *testing.T) {
	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	input := 8
	output := IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 1)

	input = 0
	output = IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 0)
}

func TestLessThan8UsingPositionMode(t *testing.T) {
	program := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	input := 0
	output := IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 1)

	input = 8
	output = IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 0)
}

func TestEqualTo8UsingImmediateMode(t *testing.T) {
	program := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	input := 8
	output := IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 1)

	input = 0
	output = IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 0)
}

func TestLessThan8UsingImmediateMode(t *testing.T) {
	program := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	input := 0
	output := IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 1)

	input = 8
	output = IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 0)
}

func TestJumpUsingPositionMode(t *testing.T) {
	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	input := 0
	output := IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 0)

	input = 5
	output = IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 1)
}

func TestJumpUsingImmediateMode(t *testing.T) {
	program := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	input := 0
	output := IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 0)

	input = 5
	output = IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 1)
}

func TestCombined(t *testing.T) {
	program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}

	input := 7
	output := IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 999)

	input = 8
	output = IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 1000)

	input = 9
	output = IntCodeWithInput(program, []Replacement{}, input)
	assertProgram(t, program, input, output, 1001)
}
