package day18

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type SnailFish struct {
	Left       *SnailFish
	Right      *SnailFish
	LeftValue  int
	RightValue int
}

func (sf *SnailFish) Magnitude() int {
	magnitude := 3*sf.LeftValue + 2*sf.RightValue

	if sf.Left != nil {
		magnitude += 3 * sf.Left.Magnitude()
	}
	if sf.Right != nil {
		magnitude += 2 * sf.Right.Magnitude()
	}

	return magnitude
}

func parseSnailFish(fish string) *SnailFish {
	fish = fish[1 : len(fish)-1]
	splitIndex := 0
	level := 0
	for i, c := range fish {
		switch c {
		case '[':
			level++
		case ']':
			level--
		case ',':
			if level == 0 {
				splitIndex = i
				break
			}
		}
	}

	leftPart, rightPart := fish[:splitIndex], fish[splitIndex+1:]

	result := &SnailFish{}

	if val, err := strconv.Atoi(leftPart); err != nil {
		result.Left = parseSnailFish(leftPart)
	} else {
		result.LeftValue = val
	}

	if val, err := strconv.Atoi(rightPart); err != nil {
		result.Right = parseSnailFish(rightPart)
	} else {
		result.RightValue = val
	}

	return result
}

func getNestingLevel(fish string) (level int) {
	return strings.Count(fish, "[") - strings.Count(fish, "]")
}

func explode(fish string) (string, bool) {
	pairRegex, _ := regexp.Compile("\\[\\d+,\\d+\\]")
	numberRegex, _ := regexp.Compile("\\d+")

	locations := pairRegex.FindAllStringIndex(fish, -1)
	if locations == nil {
		return fish, false
	}

	for _, loc := range locations {
		if getNestingLevel(fish[:loc[0]]) != 4 {
			continue
		}

		pair := strings.Split(fish[loc[0]+1:loc[1]-1], ",")
		left, _ := strconv.Atoi(pair[0])
		right, _ := strconv.Atoi(pair[1])

		leftFish, rightFish := fish[:loc[0]], fish[loc[1]:]

		if numbersLeft := numberRegex.FindAllStringIndex(leftFish, -1); numbersLeft != nil {
			firstLeft := numbersLeft[len(numbersLeft)-1]
			val, _ := strconv.Atoi(fish[firstLeft[0]:firstLeft[1]])

			leftFish = fmt.Sprintf("%s%d%s", fish[:firstLeft[0]], val+left, fish[firstLeft[1]:loc[0]])
		}

		if firstRight := numberRegex.FindStringIndex(rightFish); firstRight != nil {
			val, _ := strconv.Atoi(rightFish[firstRight[0]:firstRight[1]])
			rightFish = fmt.Sprintf("%s%d%s", rightFish[:firstRight[0]], val+right, rightFish[firstRight[1]:])
		}

		return fmt.Sprintf("%s0%s", leftFish, rightFish), true
	}

	return fish, false
}

func split(fish string) (string, bool) {
	numberRegex, _ := regexp.Compile("\\d{2}")

	loc := numberRegex.FindStringIndex(fish)
	if loc == nil {
		return fish, false
	}

	val, _ := strconv.Atoi(fish[loc[0]:loc[1]])
	left := int(val / 2)
	right := val - left

	return fmt.Sprintf("%s[%d,%d]%s", fish[:loc[0]], left, right, fish[loc[1]:]), true
}

func add(a, b string) string {
	return reduce(fmt.Sprintf("[%s,%s]", a, b))
}

func reduce(fish string) string {
	for {
		exploded, splitSuccess := false, false
		if fish, exploded = explode(fish); exploded {
			continue
		}

		if fish, splitSuccess = split(fish); splitSuccess {
			continue
		}

		if !exploded && !splitSuccess {
			return fish
		}
	}
}

func Solve() (magnitude, maxMagnitude int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	fish := ""
	var rows []string
	for scanner.Scan() {
		s := scanner.Text()
		if len(fish) == 0 {
			fish = s
		} else {
			fish = add(fish, s)
		}
		rows = append(rows, s)
	}

	magnitude = parseSnailFish(fish).Magnitude()

	for i := 0; i < len(rows); i++ {
		for j := i + 1; j < len(rows); j++ {
			magnitude := parseSnailFish(add(rows[i], rows[j])).Magnitude()
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
			}

			magnitude = parseSnailFish(add(rows[j], rows[i])).Magnitude()
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
			}
		}
	}
	return
}
