package day23

import (
	"bufio"
	"strconv"
	"strings"
)

type InstructionFn = func(i Instruction, registers map[string]int) int

type Instruction struct {
	Address  int
	Register string
	Offset   int
	Run      InstructionFn
}

func hlf(i Instruction, registers map[string]int) int {
	registers[i.Register] /= 2
	return i.Address + 1
}

func tpl(i Instruction, registers map[string]int) int {
	registers[i.Register] *= 3
	return i.Address + 1
}

func inc(i Instruction, registers map[string]int) int {
	registers[i.Register]++
	return i.Address + 1
}

func jmp(i Instruction, _ map[string]int) int {
	return i.Address + i.Offset
}

func jie(i Instruction, registers map[string]int) int {
	if registers[i.Register]%2 == 0 {
		return jmp(i, registers)
	}

	return i.Address + 1
}

func jio(i Instruction, registers map[string]int) int {
	if registers[i.Register] == 1 {
		return jmp(i, registers)
	}

	return i.Address + 1
}

func (i Instruction) Execute(registers map[string]int) int {
	return i.Run(i, registers)
}

var instructions = map[string]InstructionFn{
	"hlf": hlf,
	"tpl": tpl,
	"inc": inc,
	"jmp": jmp,
	"jie": jie,
	"jio": jio,
}

func parseInput() (program []Instruction) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var address int
	for scanner.Scan() {
		line := scanner.Text()

		instCode := line[:3]
		instruction := Instruction{Address: address, Run: instructions[instCode]}
		switch instCode {
		case "jmp":
			instruction.Offset, _ = strconv.Atoi(line[4:])
		case "jie", "jio":
			instruction.Register = line[4:5]
			instruction.Offset, _ = strconv.Atoi(line[7:])
		case "hlf", "tpl", "inc":
			instruction.Register = line[4:5]
		}
		program = append(program, instruction)
		address++
	}

	return
}

func runProgram(program []Instruction, seed int) int {
	var registers = map[string]int{
		"a": seed,
		"b": 0,
	}
	var address int
	for {
		if address < 0 || address >= len(program) {
			break
		}

		address = program[address].Execute(registers)
	}
	return registers["b"]
}

func Solve() (int, int) {
	program := parseInput()

	return runProgram(program, 0), runProgram(program, 1)
}
