package day8

import (
	"bufio"
	"strings"
)

type Entry struct {
	Patterns []string
	Digits   []string
}

func sameDigits(d1, d2 string) bool {
	if len(d1) != len(d2) {
		return false
	}

	for _, v := range d1 {
		if strings.IndexRune(d2, v) == -1 {
			return false
		}
	}

	return true
}

func convertDigit(digit string, connections map[rune]rune) (s string) {
	s = ""
	for _, v := range digit {
		s += string(connections[v])
	}

	return
}

func getDigit(digit int) string {
	switch digit {
	case 0:
		return "abcefg"
	case 1:
		return "cf"
	case 2:
		return "acdeg"
	case 3:
		return "acdfg"
	case 4:
		return "bcdf"
	case 5:
		return "abdfg"
	case 6:
		return "abdefg"
	case 7:
		return "acf"
	case 8:
		return "abcdefg"
	case 9:
		return "abcdfg"
	}

	return ""
}

func rewire(entry Entry) (number int) {
	connections := make(map[rune]rune)

	mapped := ""
	for _, v := range entry.Patterns {
		if len(v) == 2 {
			connections['c'] = rune(v[0])
			connections['f'] = rune(v[1])
			mapped += v
			break
		}
	}
	for _, v := range entry.Patterns {
		if len(v) == 3 {
			for _, c := range v {
				if strings.IndexRune(mapped, c) == -1 {
					connections['a'] = c
					mapped += string(c)
					break
				}
			}
			break
		}
	}
	for _, v := range entry.Patterns {
		if len(v) == 4 {
			for _, c := range v {
				if strings.IndexRune(mapped, c) == -1 {
					if _, found := connections['b']; !found {
						connections['b'] = c
					} else {
						connections['d'] = c
					}
					mapped += string(c)
				}
			}
			break
		}
	}
	for _, v := range entry.Patterns {
		if len(v) == 6 {
			var unmapped []rune
			for _, c := range v {
				if strings.IndexRune(mapped, c) == -1 {
					unmapped = append(unmapped, c)
				}
			}

			if len(unmapped) == 1 {
				connections['g'] = unmapped[0]
				mapped += string(unmapped[0])

				for c := 'a'; c <= 'g'; c++ {
					if strings.IndexRune(mapped, c) == -1 {
						connections['e'] = c
						mapped += string(c)
						break
					}
				}
				break
			}
		}
	}

	decodedDigits := make(map[int]string)
	var swapbd, swapfc bool
	for len(decodedDigits) < 10 {
		for _, v := range entry.Patterns {
			for i := 0; i < 10; i++ {
				d := getDigit(i)
				cd := convertDigit(d, connections)
				if sameDigits(v, cd) {
					decodedDigits[i] = v
					break
				}
			}
		}
		if len(decodedDigits) < 10 && !swapfc {
			connections['c'], connections['f'] = connections['f'], connections['c']
			swapfc = true
		} else if len(decodedDigits) < 10 && !swapbd {
			connections['b'], connections['d'] = connections['d'], connections['b']
			swapbd = true
		} else if len(decodedDigits) < 10 && swapfc {
			connections['c'], connections['f'] = connections['f'], connections['c']
			swapfc = false
		}
	}

	for _, v := range entry.Digits {
		if v == "" {
			continue
		}
		for i := 0; i < 10; i++ {
			if sameDigits(v, decodedDigits[i]) {
				number += i
				break
			}
		}
		number *= 10
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
		entries = append(entries, Entry{
			Patterns: strings.Split(parts[0][:len(parts[0])-1], " "),
			Digits:   strings.Split(parts[1][1:], " "),
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
