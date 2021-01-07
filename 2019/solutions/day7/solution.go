package day7

import (
	"math"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

func getPermutations(numbers []int) (permutations [][]int) {
	for i := 0; i < len(numbers); i++ {
		permutation := []int{numbers[i]}
		if len(numbers) > 1 {
			var remaining []int
			remaining = append(remaining, numbers[0:i]...)
			remaining = append(remaining, numbers[i+1:]...)

			for _, v := range getPermutations(remaining) {
				permutations = append(permutations, append(permutation, v...))
			}
		} else {
			permutations = append(permutations, permutation)
		}

	}

	return
}

func CalculateMaximumThrust(program toolbox.Program) int {
	permutations := getPermutations([]int{0, 1, 2, 3, 4})
	maxThrust := math.MinInt32

	for _, phaseSettings := range permutations {
		thrust := 0

		for _, phaseSetting := range phaseSettings {
			input := make(chan int)
			output := make(chan int)

			go program.IntCode(input, output, make(chan bool, 1))
			input <- phaseSetting
			input <- thrust

			thrust = <-output
		}

		if thrust > maxThrust {
			maxThrust = thrust
		}
	}

	return maxThrust
}

func CalculateThrustWithFeedbackLoop(program toolbox.Program) int {
	permutations := getPermutations([]int{5, 6, 7, 8, 9})
	maxThrust := math.MinInt32

	for _, phaseSettings := range permutations {
		amplifiers := make([]chan int, 5)
		done := make([]chan bool, 5)
		thrust := 0

		for i := 0; i < 5; i++ {
			amplifiers[i] = make(chan int, 10)
			done[i] = make(chan bool, 1)
		}

		for i := 0; i < 5; i++ {
			amplifiers[i] <- phaseSettings[i]
			if i == 0 {
				amplifiers[0] <- 0
			}
			go program.IntCode(amplifiers[i], amplifiers[(i+1)%5], done[i])
		}

		for i := 0; i < 5; i++ {
			<-done[i]
		}

		for v := range amplifiers[0] {
			if v != 0 {
				thrust = v
				break
			}
		}

		if thrust > maxThrust {
			maxThrust = thrust
		}
	}

	return maxThrust
}

func GetAnswers() (int, int) {
	program := toolbox.LoadProgram(Input)
	return CalculateMaximumThrust(*program), CalculateThrustWithFeedbackLoop(*program)
}
