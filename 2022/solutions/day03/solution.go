package day03

import (
	"strings"
	"unicode"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func count(item rune) int32 {
	if unicode.IsLower(item) {
		return item - 'a' + 1
	}

	return item - 'A' + 27
}

func Solve() (part1 int32, part2 int32) {
	lines := utils.ReadInput("./solutions/day03/input.txt", mappers.ToString)

	for _, rucksack := range lines {
		comp1, comp2 := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]

		for _, item := range comp1 {
			if strings.ContainsRune(comp2, item) {
				part1 += count(item)
				break
			}
		}
	}

	var groups [][]string
	for i, v := range lines {
		if i%3 == 0 {
			groups = append(groups, make([]string, 0, 3))
		}

		groups[len(groups)-1] = append(groups[len(groups)-1], v)
	}

	for _, g := range groups {
		for _, item := range g[0] {
			if strings.ContainsRune(g[1], item) && strings.ContainsRune(g[2], item) {
				part2 += count(item)
				break
			}
		}

	}

	return part1, part2
}
