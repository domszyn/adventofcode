package day7

import (
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

func readInput() []int {
	var ints []int
	for _, s := range strings.Split(Input, ",") {
		number, _ := strconv.Atoi(s)
		ints = append(ints, number)
	}

	return ints
}

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

			go program.IntCode(input, output, nil)
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
	var wg sync.WaitGroup

	for _, phaseSettings := range permutations {
		channels := make([]chan int, 5)
		thrust := 0

		for i := 0; i < 5; i++ {
			channels[i] = make(chan int, 10)
		}

		for i := 0; i < 5; i++ {
			wg.Add(1)
			channels[i] <- phaseSettings[i]
			if i == 0 {
				channels[0] <- 0
			}
			go program.IntCode(channels[i], channels[(i+1)%5], &wg)
		}

		wg.Wait()
		for v := range channels[0] {
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
	return CalculateMaximumThrust(readInput()), CalculateThrustWithFeedbackLoop(readInput())
}
