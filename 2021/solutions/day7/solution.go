package day7

import "math"

func minMax(crabs []int) (int, int) {
	min, max := math.MaxInt64, 0

	for _, v := range crabs {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min, max
}

func sum(pos int, crabs []int, progressive bool) (s int) {
	for _, v := range crabs {
		diff := v - pos
		if diff < 0 {
			diff = -diff
		}
		if progressive {
			s += diff * (diff + 1) / 2
		} else {
			s += diff
		}
	}
	return
}

func Solve() (minFuel, fuel int) {
	costs := make(map[int]int)
	costsProgressive := make(map[int]int)
	minPos, maxPos := minMax(Input)
	for i := minPos; i <= maxPos; i++ {
		costs[i] = sum(i, Input, false)
		costsProgressive[i] = sum(i, Input, true)
	}

	minFuel = math.MaxInt64
	fuel = math.MaxInt64
	for i := minPos; i <= maxPos; i++ {
		if costs[i] < minFuel {
			minFuel = costs[i]
		}
		if costsProgressive[i] < fuel {
			fuel = costsProgressive[i]
		}
	}

	return
}
