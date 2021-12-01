package day5

import (
	"bufio"
	"strings"
)

func isNice(s string) bool {
	var vowels int
	var hasDoubleLetter bool
	for i, r := range s {
		if strings.ContainsRune("aeiou", r) {
			vowels++
		}

		if i > 0 && r == rune(s[i-1]) {
			hasDoubleLetter = true
		}

		if i > 0 && strings.Contains("ab#cd#pq#xy", s[i-1:i+1]) {
			return false
		}
	}

	return vowels >= 3 && hasDoubleLetter
}

func isReallyNice(s string) bool {
	var containsPair bool
	var containsRepeat bool
	for i := 1; i < len(s); i++ {
		if i > 1 && s[i] == s[i-2] {
			containsRepeat = true
		}

		if strings.Contains(s[i+1:], s[i-1:i+1]) {
			containsPair = true
		}

		if i > 2 && strings.Contains(s[:i-1], s[i-1:i+1]) {
			containsPair = true
		}
	}

	return containsPair && containsRepeat
}

func CountNiceStrings() (nice, reallyNice int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		if isNice(s) {
			nice++
		}

		if isReallyNice(s) {
			reallyNice++
		}
	}

	return
}
