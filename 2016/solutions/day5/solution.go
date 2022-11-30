package day5

import (
	"crypto/md5"
	"fmt"
)

func Solve() string {
	characters := []rune{}
	const input = "abbhdwsy"

	for i := 0; len(characters) < 8; i++ {
		key := []byte(fmt.Sprintf("%s%d", input, i))
		hash := fmt.Sprintf("%x", md5.Sum(key))

		if hash[:5] == "00000" {
			characters = append(characters, rune(hash[5]))
		}
	}

	return string(characters)
}

func Solve2() string {
	characters := []rune{'z', 'z', 'z', 'z', 'z', 'z', 'z', 'z'}
	const input = "abbhdwsy"

	for i, filled := 0, 0; filled < 8; i++ {
		key := []byte(fmt.Sprintf("%s%d", input, i))
		hash := fmt.Sprintf("%x", md5.Sum(key))

		if hash[:5] == "00000" {
			pos := hash[5] - '0'
			if pos < 0 || pos > 7 || characters[pos] != 'z' {
				continue
			}

			characters[pos] = rune(hash[6])
			filled++
		}
	}

	return string(characters)
}
