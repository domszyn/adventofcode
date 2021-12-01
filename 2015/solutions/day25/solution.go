package day25

func nextCode(code int) int {
	return (code * 252533) % 33554393
}

func Solve() int {
	row2978 := []int{}

	code := 20151125
	for row := 1; len(row2978) < 3083; row++ {
		for i := row; i > 0; i-- {
			if i == 2978 {
				row2978 = append(row2978, code)
			}

			code = nextCode(code)
		}
	}

	return row2978[len(row2978)-1]
}
