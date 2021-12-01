package day11

import "regexp"

func increment(s string) string {
	if match, _ := regexp.MatchString("^[a-z]{8}$", s); !match {
		return s
	}

	output := make([]byte, 8)
	carryOver := false
	for i := len(s) - 1; i >= 0; i-- {
		var inc byte
		if i == len(s)-1 {
			inc++
		}
		if carryOver {
			inc++
		}

		output[i] = s[i] + inc
		carryOver = output[i] > 'z'
		if carryOver {
			output[i] = 'a'
		}
	}

	return string(output)
}
