package day14

import (
	"bufio"
	"math"
	"strings"
)

func getPairs(template string) map[string]int {
	pp := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pp[template[i:i+2]]++
	}
	return pp
}

func makePolymer(rules map[string][]string, steps int) int {
	polymer := Template
	pairs := getPairs(polymer)
	for i := 0; i < steps; i++ {
		nextPairs := make(map[string]int)
		for pair, count := range pairs {
			if p, found := rules[pair]; found {
				for _, pp := range p {
					nextPairs[pp] += count
				}
			}
		}
		pairs = nextPairs
	}

	counts := make(map[string]int)
	counts[Template[len(Template)-1:len(Template)]]++
	for pair, count := range pairs {
		counts[string(pair[0])] += count
	}

	minCount, maxCount := math.MaxInt64, int(0)
	for _, v := range counts {
		if v < minCount {
			minCount = v
		}

		if v > maxCount {
			maxCount = v
		}
	}

	return maxCount - minCount
}

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Rules))
	scanner.Split(bufio.ScanLines)

	rules := make(map[string][]string)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		rules[parts[0]] = []string{
			string(parts[0][0]) + parts[1],
			parts[1] + string(parts[0][1]),
		}
	}

	part1 = makePolymer(rules, 10)
	part2 = makePolymer(rules, 40)

	return
}
