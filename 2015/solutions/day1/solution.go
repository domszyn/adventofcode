package day1

func CalculateFloor() int {
	var floor int
	for _, c := range Input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
	}

	return floor
}

func CalculateBasementPosition() int {
	var floor int
	for i, c := range Input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}

		if floor == -1 {
			return i + 1
		}
	}

	return -1
}
