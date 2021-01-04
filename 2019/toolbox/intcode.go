package toolbox

import (
	"fmt"
	"strconv"
)

type Replacement struct {
	Position int
	Value    int
}

const (
	PositionMode  = iota
	ImmediateMode = iota
	RelativeMode  = iota
)

type Program []int

func (p *Program) Copy() Program {
	var program = make(Program, len(*p)*100)
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
	Value        int
	Mode         int
	RelativeBase int
}

type Instruction struct {
	Opcode     int
	Parameters []Parameter
}

func numParams(instruction int) int {
	switch instruction {
	case 1, 2, 7, 8:
		return 3
	case 3, 4, 9:
		return 1
	case 5, 6:
		return 2
	default:
		return 0
	}
}

func readNextInstruction(memory []int, relativeBase int) (i Instruction) {
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
		if paramMode == RelativeMode {
			i.Parameters[j].RelativeBase = relativeBase
		}
	}

	return
}

func (p *Program) ReadValue(parameter Parameter) (value int) {
	value = parameter.Value

	if parameter.Mode != ImmediateMode {
		value = (*p)[parameter.RelativeBase+value]
	}

	return
}

func (p *Program) WriteValue(parameter Parameter, value int) {
	if parameter.Mode != ImmediateMode {
		(*p)[parameter.Value+parameter.RelativeBase] = value
	}
}

func (p *Program) IntCode(input <-chan int, output chan<- int, done chan<- bool) Program {
	program := p.Copy()

	instructionPointer := 0
	relativeBase := 0
	for instructionPointer < len(program) {
		instruction := readNextInstruction(program[instructionPointer:], relativeBase)

		if instruction.Opcode == 99 {
			break
		}

		instructionPointer += 1 + len(instruction.Parameters)

		switch instruction.Opcode {
		case 1:
			value1 := program.ReadValue(instruction.Parameters[0])
			value2 := program.ReadValue(instruction.Parameters[1])
			program.WriteValue(instruction.Parameters[2], value1+value2)
			break
		case 2:
			value1 := program.ReadValue(instruction.Parameters[0])
			value2 := program.ReadValue(instruction.Parameters[1])
			program.WriteValue(instruction.Parameters[2], value1*value2)
			break
		case 3:
			select {
			case v := <-input:
				program.WriteValue(instruction.Parameters[0], v)
			}
			break
		case 4:
			output <- program.ReadValue(instruction.Parameters[0])
			break
		case 5:
			value := program.ReadValue(instruction.Parameters[0])

			if value > 0 {
				instructionPointer = program.ReadValue(instruction.Parameters[1])
			}
			break
		case 6:
			value := program.ReadValue(instruction.Parameters[0])

			if value == 0 {
				instructionPointer = program.ReadValue(instruction.Parameters[1])
			}
			break
		case 7:
			value1 := program.ReadValue(instruction.Parameters[0])
			value2 := program.ReadValue(instruction.Parameters[1])

			if value1 < value2 {
				program.WriteValue(instruction.Parameters[2], 1)
			} else {
				program.WriteValue(instruction.Parameters[2], 0)
			}
			break
		case 8:
			value1 := program.ReadValue(instruction.Parameters[0])
			value2 := program.ReadValue(instruction.Parameters[1])

			if value1 == value2 {
				program.WriteValue(instruction.Parameters[2], 1)
			} else {
				program.WriteValue(instruction.Parameters[2], 0)
			}
			break
		case 9:
			relativeBase += program.ReadValue(instruction.Parameters[0])
			break
		}
	}

	done <- true
	return program
}
