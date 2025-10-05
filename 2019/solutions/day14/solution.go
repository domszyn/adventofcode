package day14

func SolvePart1(input string) int {
	factory := InitFactory(input)
	factory.Needed["FUEL"] = 1
	return factory.CalculateOre()
}

func SolvePart2(input string) int {
	minFuel := 0
	maxFuel := 1000000000000
	var result int
	for minFuel <= maxFuel {
		factory := InitFactory(input)
		fuel := minFuel + int(maxFuel-minFuel)/2
		factory.Needed["FUEL"] = fuel
		ore := factory.CalculateOre()
		if ore > 1_000_000_000_000 {
			maxFuel = fuel - 1
		} else {
			result = fuel
			minFuel = fuel + 1
		}
	}

	return result
}

func GetAnswers() (int, int) {
	return SolvePart1(Input), SolvePart2(Input)
}
