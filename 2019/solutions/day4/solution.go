package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func readInput() (min, max int) {
	rangeValues := strings.Split(Input, "-")
	min, _ = strconv.Atoi(rangeValues[0])
	max, _ = strconv.Atoi(rangeValues[1])

	return
}

func hasDoubleDigits(password string, skipLargerGroups bool) bool {
	if skipLargerGroups {
		var digitGroups []string
		startIdx := 0
		for i := 0; i < len(password); i++ {
			if i == len(password)-1 || password[i] != password[i+1] {
				digitGroups = append(digitGroups, password[startIdx:i+1])
				startIdx = i + 1
			}
		}
		for _, dg := range digitGroups {
			if len(dg) == 2 {
				return true
			}
		}
	} else {
		for i := 0; i < len(password)-1; i++ {
			if password[i] == password[i+1] {
				return true
			}
		}
	}

	return false
}

func validPassword(password int, skipLargerGroups bool) bool {
	pString := fmt.Sprintf("%d", password)

	var maxDigit byte
	for i := 0; i < 6; i++ {
		currentDigit := pString[i]
		if currentDigit < maxDigit {
			return false
		}
		maxDigit = currentDigit
	}

	return hasDoubleDigits(pString, skipLargerGroups)
}

func GetAnswers() (valid, validSkipLargerGroups int) {
	min, max := readInput()
	for p := min; p <= max; p++ {
		if validPassword(p, false) {
			valid++
		}
		if validPassword(p, true) {
			validSkipLargerGroups++
		}
	}

	return
}
