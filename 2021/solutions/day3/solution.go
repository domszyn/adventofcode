package day3

import (
	"bufio"
	"strconv"
	"strings"
)

func findO2rating(base int, o2values []uint16) (o2rating uint) {
	for i := base - 1; i >= 0; i-- {
		o2rating <<= 1
		if len(o2values) == 1 {
			continue
		}

		var midbase uint16 = 1 << i
		var zeroes, ones []uint16
		for _, v := range o2values {
			if v >= midbase {
				ones = append(ones, v%midbase)
			} else {
				zeroes = append(zeroes, v)
			}
		}

		if len(ones) >= len(zeroes) {
			o2rating++
			o2values = ones
		} else {
			o2values = zeroes
		}
	}

	return
}

func findCO2rating(base int, co2values []uint16) (co2rating uint) {
	for i := base - 1; i >= 0; i-- {
		co2rating <<= 1
		if len(co2values) == 1 {
			continue
		}

		var midbase uint16 = 1 << i
		var zeroes, ones []uint16
		for _, v := range co2values {
			if v >= midbase {
				ones = append(ones, v%midbase)
			} else {
				zeroes = append(zeroes, v)
			}
		}

		if len(ones) >= len(zeroes) {
			co2values = zeroes
		} else {
			co2rating++
			co2values = ones
		}
	}

	return
}

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var values []string
	var numbers []uint16
	for scanner.Scan() {
		s := scanner.Text()
		values = append(values, s)
		number, _ := strconv.ParseUint(s, 2, len(s))
		numbers = append(numbers, uint16(number))
	}

	var gammaRate, epsilonRate int
	for i := 0; i < len(values[0]); i++ {
		var sum int
		for j := 0; j < len(values); j++ {
			sum += int(values[j][i] - byte('0'))
		}

		mostCommon := sum / (len(values) / 2)
		leastCommon := mostCommon ^ 1
		gammaRate = gammaRate<<1 + mostCommon
		epsilonRate = epsilonRate<<1 + leastCommon
	}

	part1 = gammaRate * epsilonRate
	part2 = int(findO2rating(len(values[0]), numbers) * findCO2rating(len(values[0]), numbers))

	return
}
