package utils

import (
	"math"
)

func Max(numbers []int) (max int) {
	if len(numbers) == 0 {
		return math.MaxInt64
	}

	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	return
}

func Min(numbers []int) int {
	min := math.MaxInt64
	if len(numbers) == 0 {
		return math.MinInt64
	}

	for _, n := range numbers {
		if n < min {
			min = n
		}
	}

	return min
}

func NLargest(n int, numbers []int) []int {
	if n < 1 {
		return []int{}
	}

	largest := make([]int, n)

	if len(numbers) == 0 {
		return largest
	}

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < n; j++ {
			if numbers[i] > largest[j] {
				for k := n - 1; k > j; k-- {
					largest[k] = largest[k-1]
				}
				largest[j] = numbers[i]
				break
			}
		}
	}

	return largest
}

func NSmallest(n int, numbers []int) []int {
	if n < 1 || len(numbers) == 0 {
		return []int{}
	}

	smallest := make([]int, n)
	for i := 0; i < n; i++ {
		smallest[i] = math.MaxInt64
	}

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < n; j++ {
			if numbers[i] < smallest[j] {
				for k := j; k < n-1; k++ {
					smallest[k+1] = smallest[k]
				}
				smallest[j] = numbers[i]
				break
			}
		}
	}

	return smallest
}

func Sum(numbers []int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return
}
