package day24

var packages = []int{1, 2, 3, 5, 7, 13, 17, 19, 23, 29, 31, 37, 41, 43, 53,
	59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
}

func weight(packages []int) (w int) {
	for _, p := range packages {
		w += p
	}
	return
}

func qe(packages []int) (w int) {
	w = 1
	for _, p := range packages {
		w *= p
	}
	return
}

func solve(numGroups int) int {
	totalWeight := weight(packages)
	maxWeight := totalWeight / numGroups

	var groups = make([][]int, 4)
	for i := 0; i < numGroups; i++ {
		groups[i] = []int{}
	}

	for i := len(packages); i > 0; i-- {
		for j := 0; j < numGroups; j++ {
			p := packages[i-1]

			if weight(groups[j])+p <= maxWeight {
				groups[j] = append(groups[j], p)
				break
			}
		}
	}

	minGroup := 0
	for i := 1; i < numGroups; i++ {
		if len(groups[i]) < len(groups[minGroup]) {
			minGroup = i
		}
	}

	return qe(groups[minGroup])
}

func SolvePart1() int {
	return solve(3)
}

func SolvePart2() int {
	return solve(4)
}
