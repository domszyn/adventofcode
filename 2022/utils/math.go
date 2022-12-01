package utils

import "sort"

func Max(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	tmp := make([]int, len(numbers))
	copy(tmp, numbers)
	sort.Ints(tmp)
	return tmp[len(numbers)-1]
}

func Min(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	tmp := make([]int, len(numbers))
	copy(tmp, numbers)
	sort.Ints(tmp)
	return tmp[0]
}

func Sum(numbers []int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return
}
