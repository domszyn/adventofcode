package day16

import (
	"math"
)

func getDecimal(bits []byte) (result int) {
	for _, bit := range bits {
		result <<= 1
		result += int(bit)
	}

	return
}

func sum(slice []int) (result int) {
	for _, v := range slice {
		result += v
	}
	return
}

func product(slice []int) (result int) {
	result = 1
	for _, v := range slice {
		result *= v
	}
	return
}

func min(slice []int) (result int) {
	result = slice[0]
	for _, v := range slice[1:] {
		if v < result {
			result = v
		}
	}
	return
}

func max(slice []int) (result int) {
	result = slice[0]
	for _, v := range slice[1:] {
		if v > result {
			result = v
		}
	}
	return
}

func readPacket(bits []byte, maxPackets int, nestingLevel int) (versions []int, numbers []int, i int) {
	for i = 0; i < len(bits)-7 && len(numbers) < maxPackets; {
		version := getDecimal(bits[i : i+3])
		versions = append(versions, version)
		i += 3
		typeID := getDecimal(bits[i : i+3])
		i += 3
		if typeID == 4 {
			number := 0
			for {
				group := bits[i : i+5]
				i += 5

				number <<= 4
				number += getDecimal(group[1:])
				if group[0] == 0 {
					break
				}
			}
			numbers = append(numbers, number)
		} else {
			lengthType := bits[i]
			i++
			var subpacketVersions []int
			var subpacketValues []int
			if lengthType == 0 {
				totalLength := getDecimal(bits[i : i+15])
				i += 15
				subpacketVersions, subpacketValues, _ = readPacket(bits[i:i+totalLength], math.MaxInt64, nestingLevel+1)
				i += totalLength
			} else {
				subPacketsLength := getDecimal(bits[i : i+11])
				i += 11
				var offset int
				subpacketVersions, subpacketValues, offset = readPacket(bits[i:], subPacketsLength, nestingLevel+1)
				i += offset
			}

			switch typeID {
			case 0:
				numbers = append(numbers, sum(subpacketValues))
			case 1:
				numbers = append(numbers, product(subpacketValues))
			case 2:
				numbers = append(numbers, min(subpacketValues))
			case 3:
				numbers = append(numbers, max(subpacketValues))
			case 5:
				if subpacketValues[0] > subpacketValues[1] {
					numbers = append(numbers, 1)
				} else {
					numbers = append(numbers, 0)
				}
			case 6:
				if subpacketValues[0] < subpacketValues[1] {
					numbers = append(numbers, 1)
				} else {
					numbers = append(numbers, 0)
				}
			case 7:
				if subpacketValues[0] == subpacketValues[1] {
					numbers = append(numbers, 1)
				} else {
					numbers = append(numbers, 0)
				}
			}

			versions = append(versions, subpacketVersions...)
		}
	}

	return
}

func Solve() (part1, part2 int) {
	var bits []byte
	for _, b := range Input {
		switch rune(b) {
		case '0':
			bits = append(bits, 0, 0, 0, 0)
		case '1':
			bits = append(bits, 0, 0, 0, 1)
		case '2':
			bits = append(bits, 0, 0, 1, 0)
		case '3':
			bits = append(bits, 0, 0, 1, 1)
		case '4':
			bits = append(bits, 0, 1, 0, 0)
		case '5':
			bits = append(bits, 0, 1, 0, 1)
		case '6':
			bits = append(bits, 0, 1, 1, 0)
		case '7':
			bits = append(bits, 0, 1, 1, 1)
		case '8':
			bits = append(bits, 1, 0, 0, 0)
		case '9':
			bits = append(bits, 1, 0, 0, 1)
		case 'A':
			bits = append(bits, 1, 0, 1, 0)
		case 'B':
			bits = append(bits, 1, 0, 1, 1)
		case 'C':
			bits = append(bits, 1, 1, 0, 0)
		case 'D':
			bits = append(bits, 1, 1, 0, 1)
		case 'E':
			bits = append(bits, 1, 1, 1, 0)
		case 'F':
			bits = append(bits, 1, 1, 1, 1)
		}
	}
	versions, numbers, _ := readPacket(bits, math.MaxInt64, 0)

	for _, v := range versions {
		part1 += v
	}

	part2 = numbers[0]

	return
}
