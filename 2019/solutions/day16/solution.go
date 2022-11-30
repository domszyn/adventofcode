package day16

import (
	"strconv"
	"strings"
)

func SolvePart1() string {
	return FFT(Input, 100, 0)
}

func SolvePart2() string {
	var b strings.Builder
	b.Grow(len(Input) * 10000)
	for i := 0; i < 10000; i++ {
		b.WriteString(Input)
	}
	messageOffset, _ := strconv.Atoi(Input[:7])
	return FFT2(b.String(), 100, messageOffset)
}

func GetAnswers() (string, string) {
	return SolvePart1(), SolvePart2()
}
