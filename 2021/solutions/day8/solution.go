package day8

import (
	"bufio"
	"sort"
	"strings"
)

type Entry struct {
	Patterns []string
	Digits   []string
}

func swap(s string, a, b rune) string {
	var result []rune
	for _, v := range s {
		switch v {
		case a:
			result = append(result, b)
		case b:
			result = append(result, a)
		default:
			result = append(result, v)
		}
	}
	return string(result)
}

func removeRune(s string, a rune) string {
	var result []rune
	for _, v := range s {
		if v != a {
			result = append(result, v)
		}
	}
	return string(result)
}

func removeMultiple(s string, remove string) string {
	for _, r := range remove {
		s = removeRune(s, rune(r))
	}
	return s
}

func sortRunes(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func getDigit(digit int, layout string) (result string) {
	switch digit {
	case 0:
		result = layout[:3] + layout[4:]
	case 1:
		result = layout[2:3] + layout[5:6]
	case 2:
		result = layout[:1] + layout[2:5] + layout[6:]
	case 3:
		result = layout[:1] + layout[2:4] + layout[5:]
	case 4:
		result = layout[1:4] + layout[5:6]
	case 5:
		result = layout[:2] + layout[3:4] + layout[5:]
	case 6:
		result = layout[:2] + layout[3:]
	case 7:
		result = layout[:1] + layout[2:3] + layout[5:6]
	case 8:
		result = layout
	case 9:
		result = layout[:4] + layout[5:]
	}

	return sortRunes(result)
}

func isLayoutValid(layout string, inputs []string) bool {
	zero, six, nine := getDigit(0, layout), getDigit(6, layout), getDigit(9, layout)

	for _, v := range inputs {
		if len(v) == 6 && v != zero && v != six && v != nine {
			return false
		}
	}

	return true
}

func findPatterns(patterns []string) (one, four, seven, nine string) {
	var nines []string
	for _, pattern := range patterns {
		switch len(pattern) {
		case 2:
			one = pattern
		case 3:
			seven = pattern
		case 4:
			four = pattern
		case 6:
			nines = append(nines, pattern)
		}
	}

	for _, n := range nines {
		if len(removeMultiple(n, one+four+seven)) == 1 {
			nine = n
			break
		}
	}
	return
}

func getEncodedDigits(layout string) map[string]int {
	digits := make(map[string]int)
	for i := 0; i < 10; i++ {
		digits[getDigit(i, layout)] = i
	}
	return digits
}

func rewire(entry Entry) (number int) {
	layout := "abcdefg"
	one, four, seven, nine := findPatterns(entry.Patterns)

	c, f := rune(one[0]), rune(one[1])
	four = removeMultiple(four, one)
	seven = removeMultiple(seven, one)
	a, b, d := rune(seven[0]), rune(four[0]), rune(four[1])
	nine = removeMultiple(nine, one+four+seven)
	g := rune(nine[0])

	layout = swap(layout, rune(layout[2]), c)
	layout = swap(layout, rune(layout[5]), f)
	layout = swap(layout, rune(layout[0]), a)
	layout = swap(layout, rune(layout[1]), b)
	layout = swap(layout, rune(layout[3]), d)
	layout = swap(layout, rune(layout[6]), g)
	layouts := []string{
		layout,
		swap(layout, c, f),
		swap(layout, b, d),
		swap(swap(layout, c, f), b, d),
	}

	for _, l := range layouts {
		if isLayoutValid(l, entry.Patterns) {
			layout = l
			break
		}
	}

	digits := getEncodedDigits(layout)
	for i, v := range entry.Digits {
		number += digits[v]
		if i < 3 {
			number *= 10
		}
	}
	return
}

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var entries []Entry
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, "|")
		patterns := strings.Split(parts[0][:len(parts[0])-1], " ")
		for i, p := range patterns {
			patterns[i] = sortRunes(p)
		}
		digits := strings.Split(parts[1][1:], " ")
		for i, d := range digits {
			digits[i] = sortRunes(d)
		}
		entries = append(entries, Entry{
			Patterns: patterns,
			Digits:   digits,
		})
	}

	for _, e := range entries {
		for _, digit := range e.Digits {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				part1++
			}
		}
		part2 += rewire(e)
	}

	return
}
