package day16

type Pattern []int

type PatternList struct {
	Head *PatternElement
}

type PatternElement struct {
	Value int
	Next  *PatternElement
}

func CreatePattern(repeat, length int) Pattern {
	basePattern := []int{0, 1, 0, -1}
	var pattern []int

	curIdx := -1
	for curIdx+1 < length {
		for _, i := range basePattern {
			switch i {
			case 1:
				for j := 0; j < repeat && curIdx+j < length; j++ {
					pattern = append(pattern, curIdx+j)
				}
			case -1:
				for j := 0; j < repeat && curIdx+j < length; j++ {
					pattern = append(pattern, -curIdx-j)
				}
			}

			curIdx += repeat
		}
	}

	return pattern
}

func CreatePatternList(repeat int, maxLength int) PatternList {
	basePattern := []int{0, 1, 0, -1}
	patternList := PatternList{}
	var currentElement *PatternElement
	var chainLength int

	for _, i := range basePattern {
		for j := 0; j < repeat && chainLength <= maxLength; j++ {
			elem := &PatternElement{Value: i}

			if currentElement == nil {
				patternList.Head = elem
			} else {
				currentElement.Next = elem
			}

			currentElement = elem
			chainLength++
		}
	}

	if chainLength <= maxLength {
		currentElement.Next = patternList.Head
	}

	return patternList
}
