package day17

import "math"

var containers = []int{11,
	30,
	47,
	31,
	32,
	36,
	3,
	1,
	5,
	3,
	32,
	36,
	15,
	11,
	46,
	26,
	28,
	1,
	19,
	3}

// var containers = []int{20, 15, 10, 5, 5}

type ContainerTree struct {
	Volume     int
	Containers []ContainerTree
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}

	return i * factorial(i-1)
}

func buildCombinations(containers []int, volume int) [][]int {
	if len(containers) == 1 && containers[0] == volume {
		return [][]int{{containers[0]}}
	}

	var combinations [][]int

	for i, c := range containers {
		if c == volume {
			combinations = append(combinations, []int{c})
			continue
		} else if c > volume {
			continue
		}

		remainingContainers := make([]int, 0, len(containers)-i-1)
		// remainingContainers = append(remainingContainers, containers[:i]...)
		remainingContainers = append(remainingContainers, containers[i+1:]...)
		if sum(remainingContainers)+c < volume || len(remainingContainers) == 0 {
			continue
		}

		subcombinations := buildCombinations(remainingContainers, volume-c)
		if subcombinations == nil {
			continue
		}

		for _, sc := range subcombinations {
			var combination []int
			combination = append(combination, c)
			combination = append(combination, sc...)
			combinations = append(combinations, combination)
		}
	}

	return combinations
}

func sum(slice []int) (result int) {
	for _, v := range slice {
		result += v
	}
	return
}

func countMinimalCombinations(combinations [][]int) (count int) {
	combinationLengths := make(map[int]int)
	minLength := math.MaxInt64
	for _, c := range combinations {
		l := len(c)
		combinationLengths[l]++
		if l < minLength {
			minLength = l
		}
	}

	return combinationLengths[minLength]
}

func findCombinations(containers []int, total int) (combinations [][]int) {
	if total < 0 {
		return nil
	}

	for i, c := range containers {
		if total == c {
			combinations = append(combinations, []int{c})
			return
		}

		remainingContainers := make([]int, 0, len(containers)-1)
		// remainingContainers = append(remainingContainers, containers[:i]...)
		remainingContainers = append(remainingContainers, containers[i+1:]...)
		possibleCombinations := findCombinations(remainingContainers, total-c)
		if possibleCombinations != nil {
			for _, pc := range possibleCombinations {
				combination := []int{c}
				combination = append(combination, pc...)
				combinations = append(combinations, combination)
			}
		}
	}

	return
}

func Solve() (int, int) {
	combinations := buildCombinations(containers, 150)
	return len(combinations), countMinimalCombinations(combinations)
	//return len(findCombinations(containers, 150))
}
