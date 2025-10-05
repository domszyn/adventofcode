package day17

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

const (
	_ = iota
	Up
	Right
	Down
	Left
)

type Point struct{ X, Y, Orientation int }

func solvePart1() (int, map[Point]rune) {
	program := toolbox.LoadProgram(Input)
	input := make(chan int)
	close(input)
	output := make(chan int)
	done := make(chan bool, 1)
	var chars []rune

	go program.IntCode(input, output, done)
readLoop:
	for {
		select {
		case nextChar := <-output:
			chars = append(chars, rune(nextChar))
		case <-done:
			break readLoop
		}
	}

	gridString := strings.TrimSpace(string(chars))
	lines := strings.Split(gridString, "\n")
	var totalAlignment int

	cells := make(map[Point]rune)
	for y, line := range lines {
		for x, cell := range line {
			loc := Point{X: x, Y: y}
			switch cell {
			case '^':
				loc.Orientation = Up
			case 'v':
				loc.Orientation = Down
			case '>':
				loc.Orientation = Right
			case '<':
				loc.Orientation = Left
			}
			cells[loc] = cell
		}
	}

	for loc := range cells {
		if cells[loc] == '#' &&
			cells[Point{X: loc.X, Y: loc.Y - 1}] == '#' &&
			cells[Point{X: loc.X, Y: loc.Y + 1}] == '#' &&
			cells[Point{X: loc.X - 1, Y: loc.Y}] == '#' &&
			cells[Point{X: loc.X + 1, Y: loc.Y}] == '#' {
			totalAlignment += loc.X * loc.Y
		}
	}

	return totalAlignment, cells
}

func getStartLoc(cells map[Point]rune) Point {
	for loc := range cells {
		if loc.Orientation > 0 {
			return loc
		}
	}

	return Point{-1, -1, 0}
}

var delta = map[int]Point{
	Up:    {Y: -1},
	Down:  {Y: 1},
	Left:  {X: -1},
	Right: {X: 1},
}

func findPath(cells map[Point]rune) string {
	var sb strings.Builder
	var moves int
	currentLoc := getStartLoc(cells)

	for {
		nextLoc := Point{
			X: currentLoc.X + delta[currentLoc.Orientation].X,
			Y: currentLoc.Y + delta[currentLoc.Orientation].Y,
		}

		if cells[nextLoc] == '#' {
			moves++
			currentLoc.X, currentLoc.Y = nextLoc.X, nextLoc.Y
			continue
		}

		// need to turn
		if moves > 0 {
			sb.WriteString(strconv.Itoa(moves))
			sb.WriteString(",")
			moves = 0
		}

		rightTurnDir := (currentLoc.Orientation + 1) % 4
		if rightTurnDir == 0 {
			rightTurnDir = 4
		}

		rightPos := Point{
			X: currentLoc.X + delta[rightTurnDir].X,
			Y: currentLoc.Y + delta[rightTurnDir].Y,
		}

		if cells[rightPos] == '#' {
			currentLoc.Orientation = rightTurnDir
			sb.WriteString("R,")
		} else {
			leftTurnDir := (currentLoc.Orientation - 1) % 4
			if leftTurnDir == 0 {
				leftTurnDir = 4
			}

			lefPos := Point{
				X: currentLoc.X + delta[leftTurnDir].X,
				Y: currentLoc.Y + delta[leftTurnDir].Y,
			}

			if cells[lefPos] == '#' {
				currentLoc.Orientation = leftTurnDir
				sb.WriteString("L,")
			} else {
				break
			}
		}
	}

	return strings.TrimSuffix(sb.String(), ",")
}

func parsePathWithSequences(moves, a, b, c []string) (string, bool) {
	var sequence []string
	var pos int

	for pos < len(moves) {
		switch {
		case len(moves[pos:]) >= len(a) && slices.Equal(moves[pos:pos+len(a)], a):
			sequence = append(sequence, "A")
			pos += len(a)
		case len(moves[pos:]) >= len(b) && slices.Equal(moves[pos:pos+len(b)], b):
			sequence = append(sequence, "B")
			pos += len(b)
		case len(moves[pos:]) >= len(c) && slices.Equal(moves[pos:pos+len(c)], c):
			sequence = append(sequence, "C")
			pos += len(c)
		default:
			return "", false
		}
	}

	return strings.Join(sequence, ","), true
}

func compressPath(path string) (string, bool) {
	moves := strings.Split(path, ",")

	for lenA := 2; lenA <= 10; lenA++ {
		seqA := moves[:lenA]
		if len(strings.Join(seqA, ",")) > 20 {
			break
		}

		startB := 0
		for len(moves[startB:]) >= lenA && slices.Equal(moves[startB:startB+lenA], seqA) {
			startB += lenA
		}

		if startB >= len(moves) {
			continue
		}

		for lenB := 2; lenB <= 10; lenB++ {
			if startB+lenB > len(moves) {
				continue
			}

			seqB := moves[startB : startB+lenB]

			if slices.Equal(seqA, seqB) {
				continue
			}

			if len(strings.Join(seqB, ",")) > 20 {
				break
			}

			startC, pos := -1, 0
			for pos < len(moves) {
				if len(moves[pos:]) >= lenA && slices.Equal(moves[pos:pos+lenA], seqA) {
					pos += lenA
				} else if len(moves[pos:]) >= lenB && slices.Equal(moves[pos:pos+lenB], seqB) {
					pos += lenB
				} else {
					startC = pos
					break
				}
			}

			if startC == -1 {
				continue
			}

			endC := startC + 2
			for ; endC < len(moves); endC++ {
				if len(moves[endC:]) >= lenA && slices.Equal(moves[endC:endC+lenA], seqA) ||
					len(moves[endC:]) >= lenB && slices.Equal(moves[endC:endC+lenB], seqB) {
					break
				}
			}

			if endC > len(moves) {
				continue
			}

			seqC := moves[startC:endC]
			if len(strings.Join(seqC, ",")) > 20 || slices.Equal(seqA, seqC) || slices.Equal(seqB, seqC) {
				continue
			}

			path, ok := parsePathWithSequences(moves, seqA, seqB, seqC)

			if ok {
				return strings.Join([]string{
					path,
					strings.Join(seqA, ","),
					strings.Join(seqB, ","),
					strings.Join(seqC, ","),
				}, "\n"), ok
			}
		}
	}

	return "", false
}

func solvePart2(cells map[Point]rune) int {
	fullPath, ok := compressPath(findPath(cells))
	if !ok {
		return -1
	}

	finalInput := fullPath + "\nn\n"

	program := toolbox.LoadProgram(Input).Patch(
		[]toolbox.Replacement{{Position: 0, Value: 2}},
	)
	input := make(chan int)
	output := make(chan int)
	done := make(chan bool, 1)

	go program.IntCode(input, output, done)

	var lastOutput int
	var debugOutput strings.Builder
	inputIdx := 0

loop:
	for {
		var inputChan chan int
		var nextInput int

		if inputIdx < len(finalInput) {
			inputChan = input
			nextInput = int(finalInput[inputIdx])
		}

		select {
		case inputChan <- nextInput:
			debugOutput.WriteRune(rune(nextInput))
			inputIdx++
		case val := <-output:
			lastOutput = val
			debugOutput.WriteRune(rune(val))
		case <-done:
			break loop
		}
	}

	fmt.Println(debugOutput.String())
	return lastOutput
}

func GetAnswers() (int, int) {
	totalAlignment, cells := solvePart1()
	return totalAlignment, solvePart2(cells)
}
