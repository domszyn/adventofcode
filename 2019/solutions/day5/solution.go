package day5

import (
	"strconv"
	"strings"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

func readInput() []int {
	var ints []int
	for _, s := range strings.Split(Input, ",") {
		number, _ := strconv.Atoi(s)
		ints = append(ints, number)
	}

	return ints
}

func GetAnswers() (int, int) {
	diagnosticCode1 := toolbox.IntCodeWithInput(readInput(), []toolbox.Replacement{}, 1)
	diagnosticCode5 := toolbox.IntCodeWithInput(readInput(), []toolbox.Replacement{}, 5)
	return diagnosticCode1, diagnosticCode5

}
