package day3

import (
	"bufio"
	"strconv"
	"strings"
)

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var numbers []int64
	var values []string
	for scanner.Scan() {
		s := scanner.Text()
		values = append(values, s)
		n, _ := strconv.ParseInt(s, 2, 0)
		numbers = append(numbers, n)
	}

	gammaRate := ""
	epsilonRate := ""
	for i := 0; i < len(values[0]); i++ {
		var digits = make(map[string]int)
		for j := 0; j < len(values); j++ {
			digits[string(values[j][i])]++
		}

		if digits["0"] > digits["1"] {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
	}

	gr, _ := strconv.ParseInt(gammaRate, 2, 0)
	er, _ := strconv.ParseInt(epsilonRate, 2, 0)

	part1 = int(gr) * int(er)

	var o2rating, co2rating string
	searchNumbers := make([]string, len(values))
	copy(searchNumbers, values)

	for len(searchNumbers) > 1 {
		for i := 0; i < len(searchNumbers[0]); i++ {
			if len(searchNumbers) == 1 {
				o2rating += string(searchNumbers[0][i])
				continue
			}

			var digits = make(map[string]int)
			for j := 0; j < len(searchNumbers); j++ {
				digits[string(searchNumbers[j][i])]++
			}

			if digits["1"] >= digits["0"] {
				o2rating += "1"
			} else {
				o2rating += "0"
			}

			var tmpNumbers []string
			for j := 0; j < len(searchNumbers); j++ {
				if string(searchNumbers[j][i]) == string(o2rating[len(o2rating)-1]) {
					tmpNumbers = append(tmpNumbers, searchNumbers[j])
				}
			}
			searchNumbers = tmpNumbers
		}
	}

	searchNumbers = make([]string, len(values))
	copy(searchNumbers, values)

	for len(searchNumbers) > 1 {
		for i := 0; i < len(values[0]); i++ {
			if len(searchNumbers) == 1 {
				co2rating += string(searchNumbers[0][i])
				continue
			}

			var digits = make(map[string]int)
			for j := 0; j < len(searchNumbers); j++ {
				digits[string(searchNumbers[j][i])]++
			}

			if digits["0"] <= digits["1"] {
				co2rating += "0"
			} else {
				co2rating += "1"
			}

			var tmpNumbers []string
			for j := 0; j < len(searchNumbers); j++ {
				if string(searchNumbers[j][i]) == string(co2rating[len(co2rating)-1]) {
					tmpNumbers = append(tmpNumbers, searchNumbers[j])
				}
			}
			searchNumbers = tmpNumbers
		}
	}

	o2, _ := strconv.ParseInt(o2rating, 2, 0)
	co2, _ := strconv.ParseInt(co2rating, 2, 0)
	part2 = int(o2) * int(co2)

	return
}
