package day1

func SolvePart1() int {
	var increases int

	for i := 0; i < len(Input)-1; i++ {
		if Input[i] < Input[i+1] {
			increases++
		}
	}

	return increases
}

func SolvePart2() int {
	var windows []int

	for i := 0; i < len(Input)-2; i++ {
		windows = append(windows, Input[i]+Input[i+1]+Input[i+2])
	}

	var increases int
	for i := 0; i < len(windows)-1; i++ {
		if windows[i] < windows[i+1] {
			increases++
		}
	}

	return increases
}
