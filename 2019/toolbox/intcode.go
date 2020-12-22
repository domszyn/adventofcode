package toolbox

type Replacement struct {
	Position int
	Value    int
}

func numParams(instruction int) int {
	switch instruction {
	case 1, 2:
		return 3
	case 3:
		return 1
	default:
		return 0
	}
}

func IntCode(program []int, replacements []Replacement) int {
	var memory = make([]int, len(program))
	copy(memory, program)
	for _, r := range replacements {
		memory[r.Position] = r.Value
	}

	instructionPointer := 0

	for instructionPointer < len(memory) {
		instruction := memory[instructionPointer]
		paramPointers := memory[instructionPointer+1 : instructionPointer+numParams(instruction)+1]
		instructionPointer += 1 + len(paramPointers)

		if instruction == 99 {
			break
		}

		if len(paramPointers) < 3 ||
			paramPointers[0] >= len(memory) ||
			paramPointers[1] >= len(memory) ||
			paramPointers[2] >= len(memory) {
			break
		}

		switch instruction {
		case 1:
			memory[paramPointers[2]] = memory[paramPointers[0]] + memory[paramPointers[1]]
			break
		case 2:
			memory[paramPointers[2]] = memory[paramPointers[0]] * memory[paramPointers[1]]
			break
		}
	}

	return memory[0]
}
