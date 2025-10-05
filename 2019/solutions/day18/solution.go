package day18

func solvePart1() int {
	grid := LoadGrid(Input)
	startState := SearchState{LastKeys: "@", KeyMask: 0}
	return grid.findShortestDistance(startState)
}

func solvePart2() int {
	grid := LoadGrid(Input2)
	startState := SearchState{LastKeys: "0123", KeyMask: 0}
	return grid.findShortestDistance(startState)
}

func GetAnswers() (int, int) {
	return solvePart1(), solvePart2()
}
