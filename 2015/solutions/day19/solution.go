package day19

import (
	"bufio"
	"regexp"
	"strings"
)

type RulesMap map[string][]string

func parseRules(rules string) RulesMap {
	rulesMap := make(RulesMap)
	scanner := bufio.NewScanner(strings.NewReader(rules))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		rule := scanner.Text()
		parts := strings.Split(rule, " => ")
		input, output := parts[0], parts[1]
		rulesMap[input] = append(rulesMap[input], output)
	}

	return rulesMap
}

func splitMolecules(input string) []string {
	r, _ := regexp.Compile("e|([A-Z]([a-z])*)")

	return r.FindAllString(input, -1)
}

func calculateReplacements(rulesMap RulesMap, input string) map[string]bool {
	molecules := splitMolecules(input)

	replacements := make(map[string]bool)
	for i := 0; i < len(molecules); i++ {
		molecule := molecules[i]
		for _, r := range rulesMap[molecule] {
			before := concatenate(molecules[:i]...)
			after := concatenate(molecules[i+1:]...)
			replacement := concatenate(before, r, after)
			replacements[replacement] = true
		}
	}

	return replacements
}

func concatenate(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

func Solve() (part1, part2 int) {
	rulesMap := parseRules(Rules)
	part1 = len(calculateReplacements(rulesMap, Input))

	molecules := splitMolecules(Input)
	part2 = len(molecules) - 1
	for _, v := range molecules {
		switch v {
		case "Rn", "Ar":
			part2--
		case "Y":
			part2 -= 2
		}
	}

	return
}
