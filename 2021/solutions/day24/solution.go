package day24

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type InstructionFn = func(i Instruction, registers map[string]int)
type GetValueFn = func(registers map[string]int) int

type Instruction struct {
	Text     string
	Register string
	GetValue GetValueFn
	Run      InstructionFn
	Type     string
}

func inp(i Instruction, registers map[string]int) {
	registers[i.Register] = i.GetValue(registers)
}

func add(i Instruction, registers map[string]int) {
	registers[i.Register] += i.GetValue(registers)
}

func mul(i Instruction, registers map[string]int) {
	registers[i.Register] *= i.GetValue(registers)
}

func div(i Instruction, registers map[string]int) {
	registers[i.Register] = int(registers[i.Register] / i.GetValue(registers))
}

func mod(i Instruction, registers map[string]int) {
	registers[i.Register] %= i.GetValue(registers)
}

func eql(i Instruction, registers map[string]int) {
	if registers[i.Register] == i.GetValue(registers) {
		registers[i.Register] = 1
	} else {
		registers[i.Register] = 0
	}
}

func (i Instruction) Execute(registers map[string]int) {
	i.Run(i, registers)
}

var instructions = map[string]InstructionFn{
	"inp": inp,
	"add": add,
	"mul": mul,
	"div": div,
	"mod": mod,
	"eql": eql,
}

func parseInput() (program []Instruction) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		instruction := Instruction{
			Type:     parts[0],
			Run:      instructions[parts[0]],
			Register: parts[1],
			Text:     line,
		}

		if len(parts) == 3 {
			if val, err := strconv.Atoi(parts[2]); err == nil {
				instruction.GetValue = func(_ map[string]int) int {
					return val
				}
			} else {
				instruction.GetValue = func(registers map[string]int) int {
					return registers[parts[2]]
				}
			}
		}

		program = append(program, instruction)
	}

	return
}

func runProgram(program []Instruction, inputs []int) (results []int) {
	var registers = map[string]int{
		"x": 0,
		"y": 0,
		"z": 0,
		"w": 0,
	}

	for _, instruction := range program {
		if instruction.Type == "inp" {
			instruction.GetValue = func(_ map[string]int) int {
				val := inputs[0]
				inputs = inputs[1:]
				return val
			}

			results = append(results, registers["z"])
		}

		instruction.Execute(registers)
	}

	results = append(results, registers["z"])
	return
}

func nextInput(input [14]int) [14]int {
	var result [14]int
	result[13] = -1
	for i := 13; i >= 0; i-- {
		result[i] = result[i] + input[i]

		if i > 0 && result[i] == 0 {
			result[i] = 9
			result[i-1]--
		}
	}

	return result
}

func validInput(input [14]int) bool {
	return input[3]-3 == input[4] &&
		input[6]-2 == input[7] &&
		input[5]-1 == input[8] &&
		input[2]-7 == input[9] &&
		input[10]-6 == input[11] &&
		input[1]+7 == input[12] &&
		input[0] == 1 && input[13] == 9
}

func Solve() (part1, part2 [14]int) {
	inputs := [14]int{
		1, 2, 9, 9, 9, 9, 9,
		7, 9, 9, 9, 9, 9, 9} // {1, 1, 4, 9, 6, 9, 9, 7, 9, 4, 9, 6, 9, 1}
	var validInputs [][14]int

	for len(validInputs) < 1 {
		fmt.Printf("%v\n", inputs)
		if validInput(inputs) {
			program := parseInput()
			results := runProgram(program, inputs[:])
			if results[len(results)-1] == 0 {
				validInputs = append(validInputs, inputs)
			}
			// fmt.Printf("%v => %v\n", inputs, results)
		}
		// fmt.Printf("%v\n", inputs)
		inputs = nextInput(inputs)
	}

	part1 = inputs

	return
}
