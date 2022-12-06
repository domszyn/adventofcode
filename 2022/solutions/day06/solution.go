package day06

import (
	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func findMarker(buffer string, nDifferent int) int {
	for i := 0; i < len(buffer)-nDifferent; i++ {
		markerCandidate := buffer[i : i+nDifferent]
		notCandidate := false

		for i := 0; i < nDifferent; i++ {
			for j := i + 1; j < nDifferent; j++ {
				if markerCandidate[i] == markerCandidate[j] {
					notCandidate = true
					break
				}
			}
		}

		if notCandidate {
			continue
		}

		return (i + nDifferent)
	}

	return 0
}

func Solve() (int, int) {
	buffer := utils.ReadInput("./solutions/day06/input.txt", mappers.ToString)[0]

	return findMarker(buffer, 4), findMarker(buffer, 14)
}
