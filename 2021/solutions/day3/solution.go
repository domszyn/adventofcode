package day3

import (
	"bufio"
	"strconv"
	"strings"
)

func splitByFirstBit(midbase uint16, values []uint16) (ones, zeroes []uint16) {
	for _, v := range values {
		if v >= midbase {
			ones = append(ones, v%midbase)
		} else {
			zeroes = append(zeroes, v)
		}
	}

	return
}

type FilterFn = func(ones, zeroes []uint16) ([]uint16, uint)

func findRating(base int, values []uint16, filter FilterFn) (rating uint) {
	for i := base - 1; i >= 0; i-- {
		rating <<= 1
		if len(values) == 1 {
			continue
		}

		var midbase uint16 = 1 << i
		var offset uint
		ones, zeroes := splitByFirstBit(midbase, values)
		values, offset = filter(ones, zeroes)
		rating += offset
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

	base := len(values[0])
	var gammaRate, epsilonRate int
	for i := 0; i < base; i++ {
		var sum int
		for _, v := range values {
			sum += int(v[i] - byte('0'))
		}

		mostCommon := sum / (len(values) / 2)
		leastCommon := mostCommon ^ 1
		gammaRate = gammaRate<<1 + mostCommon
		epsilonRate = epsilonRate<<1 + leastCommon
	}

	part1 = gammaRate * epsilonRate
	o2rating := findRating(base, numbers, func(ones, zeroes []uint16) ([]uint16, uint) {
		if len(ones) >= len(zeroes) {
			return ones, 1
		}
		return zeroes, 0
	})
	co2rating := findRating(base, numbers, func(ones, zeroes []uint16) ([]uint16, uint) {
		if len(ones) >= len(zeroes) {
			return zeroes, 0
		}
		return ones, 1
	})
	part2 = int(o2rating * co2rating)

	return
}
