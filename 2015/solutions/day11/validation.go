package day11

import "regexp"

func hasStraightThrees(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i]+1 == s[i+1] && s[i]+2 == s[i+2] {
			return true
		}
	}

	return false
}

func hasForbiddenLetters(s string) bool {
	match, _ := regexp.MatchString("i|o|l", s)
	return match
}

func hasTwoPairs(s string) bool {
	pairs := make(map[string]bool)

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			pairs[s[i:i+2]] = true
			i++
			continue
		}
	}

	return len(pairs) > 1
}

func isValid(s string) bool {
	return !hasForbiddenLetters(s) && hasStraightThrees(s) && hasTwoPairs(s)
}
