package day8

import (
	"bufio"
	"strconv"
	"strings"
)

func Solve() (count int, count2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		unquote, _ := strconv.Unquote(s)
		quote := strconv.Quote(s)

		count += len(s) - len(unquote)
		count2 += len(quote) - len(s)
	}

	return
}
