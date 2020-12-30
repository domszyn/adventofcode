package day8

func splitString(input string, size int) []string {
	var layers []string
	for len(input) > 0 {
		layers = append(layers, input[0:size])
		input = input[size:]
	}
	return layers
}

func readInput() []string {
	input := Input
	return splitString(input, 150)
}

func countDigits(layer string, digit rune) (count int) {
	for _, d := range []rune(layer) {
		if d == digit {
			count++
		}
	}

	return
}

func combineLayers(layers []string) string {
	imageSize := len(layers[0])
	combined := make([]rune, imageSize)
	for i := 0; i < imageSize; i++ {
		for j := 0; j < len(layers); j++ {
			if layers[j][i] == '0' {
				combined[i] = ' '
				break
			}

			if layers[j][i] == '1' {
				combined[i] = '#'
				break
			}
		}
	}

	return string(combined)
}

func solvePart1() int {
	layers := readInput()
	minZeroes := len(layers[0])
	idx := len(layers)
	for i, layer := range layers {
		zeroCount := countDigits(layer, '0')
		if zeroCount < minZeroes {
			minZeroes = zeroCount
			idx = i
		}
	}

	return countDigits(layers[idx], '1') * countDigits(layers[idx], '2')
}

func decodeImage() string {
	layers := readInput()
	combined := splitString(combineLayers(layers), 25)
	image := "\n"
	for _, row := range combined {
		image += row + "\n"
	}
	return image
}

func GetAnswers() (int, string) {
	return solvePart1(), decodeImage()
}
