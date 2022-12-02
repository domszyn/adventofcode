package day02

import (
	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

const (
	Rock     = "A"
	Paper    = "B"
	Scissors = "C"
)

func Solve() (part1 int, part2 int) {
	scores := map[string]int{
		"X":    1,
		"Y":    2,
		"Z":    3,
		"Lose": 0,
		"Draw": 3,
		"Win":  6,
	}

	strategy1 := map[string]int{
		"A X": scores["X"] + scores["Draw"], // Rock on rock
		"A Y": scores["Y"] + scores["Win"],  // Paper wins rock
		"A Z": scores["Z"] + scores["Lose"], // Scissors lose to rock
		"B X": scores["X"] + scores["Lose"], // Rock loses to paper
		"B Y": scores["Y"] + scores["Draw"], // Paper on paper
		"B Z": scores["Z"] + scores["Win"],  // Scissors win paper
		"C X": scores["X"] + scores["Win"],  // Rock wins scissors
		"C Y": scores["Y"] + scores["Lose"], // Paper loses to scissors
		"C Z": scores["Z"] + scores["Draw"], // Scissors on scissors
	}

	strategy2 := map[string]int{
		"A X": strategy1["A Z"], // Lose to rock
		"A Y": strategy1["A X"], // Draw on rock
		"A Z": strategy1["A Y"], // Win rock with paper
		"B X": strategy1["B X"], // Lose to paper
		"B Y": strategy1["B Y"], // Draw on paper
		"B Z": strategy1["B Z"], // Win paper with scissors
		"C X": strategy1["C Y"], // Lose to scissors
		"C Y": strategy1["C Z"], // Draw on scissors
		"C Z": strategy1["C X"], // Win scissors with rock
	}

	lines := utils.ReadInput("./solutions/day02/input.txt", mappers.ToString)

	for _, l := range lines {
		part1 += strategy1[l]
		part2 += strategy2[l]
	}

	return
}
