package day12

import (
	"bufio"
	"strconv"
	"strings"
)

type InstructionFn = func(i Instruction, registers map[string]int) int

type Instruction struct {
	Address  int
	Register string
	Value    string
	Run      InstructionFn
}

func cpy(i Instruction, registers map[string]int) int {
	value, err := strconv.Atoi(i.Value)
	if err != nil {
		value = registers[i.Value]
	}

	registers[i.Register] = value
	return i.Address + 1
}

func inc(i Instruction, registers map[string]int) int {
	registers[i.Register]++
	return i.Address + 1
}

func dec(i Instruction, registers map[string]int) int {
	registers[i.Register]--
	return i.Address + 1
}

func jnz(i Instruction, registers map[string]int) int {
	reg, err := strconv.Atoi(i.Register)
	if err != nil {
		reg = registers[i.Register]
	}

	value, err := strconv.Atoi(i.Value)
	if err != nil {
		value = registers[i.Value]
	}

	if reg != 0 {
		return i.Address + value
	}

	return i.Address + 1
}

func (i Instruction) Execute(registers map[string]int) int {
	return i.Run(i, registers)
}

var instructions = map[string]InstructionFn{
	"cpy": cpy,
	"inc": inc,
	"dec": dec,
	"jnz": jnz,
}

func parseInput() (program []Instruction) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var address int
	for scanner.Scan() {
		line := scanner.Text()

		instCodes := strings.Split(line, " ")
		instCode := instCodes[0]
		instruction := Instruction{Address: address, Run: instructions[instCode]}
		switch instCode {
		case "cpy":
			instruction.Value = instCodes[1]
			instruction.Register = instCodes[2]
		case "inc", "dec":
			instruction.Register = instCodes[1]
		case "jnz":
			instruction.Register = instCodes[1]
			instruction.Value = instCodes[2]
		}
		program = append(program, instruction)
		address++
	}

	return
}

func runProgram(program []Instruction, seed int) int {
	var registers = map[string]int{
		"a": 0,
		"b": 0,
		"c": seed,
		"d": 0,
	}
	var address int
	for {
		if address < 0 || address >= len(program) {
			break
		}

		address = program[address].Execute(registers)
	}
	return registers["a"]
}

func Solve() (int, int) {
	program := parseInput()

	return runProgram(program, 0), runProgram(program, 1)
}
