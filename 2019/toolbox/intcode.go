package toolbox

import (
	"fmt"
	"strconv"
	"sync"
)

type Replacement struct {
	Position int
	Value    int
}

const (
	PositionMode  = iota
	ImmediateMode = iota
)

type Program []int

func (p *Program) Copy() Program {
	var program = make(Program, len(*p))
	copy(program, *p)
	return program
}

func (p *Program) Patch(replacements []Replacement) *Program {
	program := p.Copy()
	for _, r := range replacements {
		program[r.Position] = r.Value
	}

	return &program
}

type Parameter struct {
	Value int
	Mode  int
}

type Instruction struct {
	Opcode     int
	Parameters []Parameter
	Input      int
	Output     int
}

func numParams(instruction int) int {
	switch instruction {
	case 1, 2, 7, 8:
		return 3
	case 3, 4:
		return 1
	case 5, 6:
		return 2
	default:
		return 0
	}
}

func readNextInstruction(memory []int) (i Instruction) {
	instructionValue := fmt.Sprintf("%05d", memory[0])
	i.Opcode, _ = strconv.Atoi(instructionValue[len(instructionValue)-2:])
	i.Parameters = make([]Parameter, numParams(i.Opcode))

	for j := 0; j < len(i.Parameters); j++ {
		paramPos := len(instructionValue) - j - 3
		paramMode, _ := strconv.Atoi(instructionValue[paramPos : paramPos+1])
		i.Parameters[j] = Parameter{
			Value: memory[j+1],
			Mode:  paramMode,
		}
	}

	return
}

func (p *Program) IntCode(input <-chan int, output chan<- int, wg *sync.WaitGroup) Program {
	program := p.Copy()
	if wg != nil {
		defer wg.Done()
	}

	instructionPointer := 0
	for instructionPointer < len(program) {
		instruction := readNextInstruction(program[instructionPointer:])

		if instruction.Opcode == 99 {
			break
		}

		instructionPointer += 1 + len(instruction.Parameters)

		switch instruction.Opcode {
		case 1:
			value1 := instruction.Parameters[0].Value
			if instruction.Parameters[0].Mode == PositionMode {
				value1 = program[value1]
			}
			value2 := instruction.Parameters[1].Value
			if instruction.Parameters[1].Mode == PositionMode {
				value2 = program[value2]
			}
			program[instruction.Parameters[2].Value] = value1 + value2
			break
		case 2:
			value1 := instruction.Parameters[0].Value
			if instruction.Parameters[0].Mode == PositionMode {
				value1 = program[value1]
			}
			value2 := instruction.Parameters[1].Value
			if instruction.Parameters[1].Mode == PositionMode {
				value2 = program[value2]
			}
			program[instruction.Parameters[2].Value] = value1 * value2
			break
		case 3:
			select {
			case v := <-input:
				program[instruction.Parameters[0].Value] = v
			}
			break
		case 4:
			value := instruction.Parameters[0].Value

			if instruction.Parameters[0].Mode == PositionMode {
				value = program[instruction.Parameters[0].Value]
			}

			output <- value
			break
		case 5:
			value := instruction.Parameters[0].Value

			if instruction.Parameters[0].Mode == PositionMode {
				value = program[instruction.Parameters[0].Value]
			}

			if value > 0 && instruction.Parameters[1].Mode == PositionMode {
				instructionPointer = program[instruction.Parameters[1].Value]
			} else if value > 0 {
				instructionPointer = instruction.Parameters[1].Value
			}
			break
		case 6:
			value := instruction.Parameters[0].Value

			if instruction.Parameters[0].Mode == PositionMode {
				value = program[instruction.Parameters[0].Value]
			}

			if value == 0 && instruction.Parameters[1].Mode == PositionMode {
				instructionPointer = program[instruction.Parameters[1].Value]
			} else if value == 0 {
				instructionPointer = instruction.Parameters[1].Value
			}
			break
		case 7:
			value1 := instruction.Parameters[0].Value
			if instruction.Parameters[0].Mode == PositionMode {
				value1 = program[value1]
			}
			value2 := instruction.Parameters[1].Value
			if instruction.Parameters[1].Mode == PositionMode {
				value2 = program[value2]
			}
			if value1 < value2 {
				program[instruction.Parameters[2].Value] = 1
			} else {
				program[instruction.Parameters[2].Value] = 0
			}
			break
		case 8:
			value1 := instruction.Parameters[0].Value
			if instruction.Parameters[0].Mode == PositionMode {
				value1 = program[value1]
			}
			value2 := instruction.Parameters[1].Value
			if instruction.Parameters[1].Mode == PositionMode {
				value2 = program[value2]
			}
			if value1 == value2 {
				program[instruction.Parameters[2].Value] = 1
			} else {
				program[instruction.Parameters[2].Value] = 0
			}
			break
		}

	}

	return program
}
