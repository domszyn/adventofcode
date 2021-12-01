package day9

import (
	"fmt"
	"strings"
)

func getDecompressedLength(input string, nested bool) (length int) {
	for i := 0; i < len(input); i++ {
		if input[i] != '(' {
			length++
			continue
		}
		end := strings.Index(input[i:], ")") + i + 1
		marker := input[i:end]
		var l, r int
		fmt.Sscanf(marker, "(%dx%d)", &l, &r)
		if nested {
			length += r * getDecompressedLength(input[end:end+l], nested)
		} else {
			length += l * r
		}
		i = end + l - 1
	}

	return
}

func Solve() int {
	return getDecompressedLength(Input, false)
}

func Solve2() int {
	return getDecompressedLength(Input, true)
}
