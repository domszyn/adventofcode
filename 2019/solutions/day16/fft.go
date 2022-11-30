package day16

func convertToDigits(input string) []int {
	digits := make([]int, len(input)+1)
	for i := 0; i < len(input); i++ {
		digits[i+1] = int(input[i] - '0')
	}

	return digits
}

func convertToString(digits []int, offset int) string {
	output := make([]rune, 8)
	for i := offset; i < offset+8; i++ {
		output[i-offset] = rune(digits[i+1]) + '0'
	}
	return string(output)
}

func sum(slice []int) int {
	var result int
	for _, val := range slice {
		result += val
	}

	return result
}

func transform(digits []int, pos int, offset int) int {
	var output int

	if pos < offset {
		return 0
	}

	for i := 0; i < len(digits)-pos; i += pos * 4 {
		var addSlice, subSlice []int
		if len(digits) < i+pos*2 {
			addSlice = digits[i+pos:]
		} else {
			addSlice = digits[i+pos : i+pos*2]
		}

		if i+pos*3 < len(digits) && len(digits) < i+pos*4 {
			subSlice = digits[i+pos*3:]
		} else if i+pos*3 < len(digits) {
			subSlice = digits[i+pos*3 : i+pos*4]
		}

		output += sum(addSlice)
		output -= sum(subSlice)
	}

	if output < 0 {
		output *= -1
	}

	return output
}

func FFT(input string, phaseCount int, messageOffset int) string {
	inputDigits := convertToDigits(input)
	outputDigits := make([]int, len(inputDigits))

	for phase := 1; phase <= phaseCount; phase++ {
		for i := 1; i < len(inputDigits); i++ {
			outputDigits[i] = transform(inputDigits, i, messageOffset) % 10
		}

		inputDigits = outputDigits
	}

	return convertToString(inputDigits, messageOffset)
}

func FFT2(input string, phaseCount int, messageOffset int) string {
	if messageOffset <= len(input)/2 {
		panic("Unapplicable")
	}

	inputDigits := convertToDigits(input)
	outputDigits := make([]int, len(inputDigits))

	for phase := 1; phase <= phaseCount; phase++ {
		sum := transform(inputDigits, messageOffset, messageOffset)
		outputDigits[messageOffset] = sum % 10
		for i := messageOffset + 1; i < len(inputDigits); i++ {
			sum -= inputDigits[i-1]
			outputDigits[i] = sum % 10
		}

		copy(inputDigits, outputDigits)
	}

	return convertToString(inputDigits, messageOffset)
}
