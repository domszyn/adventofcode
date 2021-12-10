package day10

import (
	"bufio"
	"sort"
	"strings"
)

func moreReplacements(s string) bool {
	return strings.Contains(s, "()") ||
		strings.Contains(s, "[]") ||
		strings.Contains(s, "{}") ||
		strings.Contains(s, "<>")
}

func isClosing(c rune) bool { return c == ')' || c == ']' || c == '}' || c == '>' }

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var scores []int
	for scanner.Scan() {
		r, corrupted := scanner.Text(), false

		for moreReplacements(r) {
			r = strings.ReplaceAll(r, "()", "")
			r = strings.ReplaceAll(r, "[]", "")
			r = strings.ReplaceAll(r, "{}", "")
			r = strings.ReplaceAll(r, "<>", "")
		}

		for _, c := range r {
			switch rune(c) {
			case ')':
				part1 += 3
			case ']':
				part1 += 57
			case '}':
				part1 += 1197
			case '>':
				part1 += 25137
			}

			if isClosing(c) {
				corrupted = true
				break
			}
		}

		if !corrupted {
			score := 0
			for i := len(r) - 1; i >= 0; i-- {
				score *= 5
				switch r[i] {
				case '(':
					score++
				case '[':
					score += 2
				case '{':
					score += 3
				case '<':
					score += 4
				}
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	part2 = scores[len(scores)/2]
	return
}
