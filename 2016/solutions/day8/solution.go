package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	tm "github.com/buger/goterm"
)

type Screen [][]bool

func (s *Screen) getColumn(i int) (column []bool) {
	for _, row := range *s {
		column = append(column, row[i])
	}
	return
}

func (s *Screen) setColumn(i int, column []bool) {
	for r := range *s {
		(*s)[r][i] = column[r]
	}
}

func (s *Screen) rect(x, y int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			(*s)[j][i] = true
		}
	}
}

func (s Screen) countLitPixels() (count int) {
	for _, row := range s {
		for _, pixel := range row {
			if pixel {
				count++
			}
		}
	}
	return
}

func (s *Screen) rotateRow(y, offset int) {
	(*s)[y] = rotate((*s)[y], offset)
}

func (s *Screen) rotateColumn(x, offset int) {
	s.setColumn(x, rotate(s.getColumn(x), offset))
}

func (s Screen) print() {
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			tm.MoveCursor(x+1, y+5)
			if s[y][x] {
				tm.Print("#")
			} else {
				tm.Print(".")
			}
			tm.Flush()
		}
	}
}

func rotate(array []bool, offset int) (result []bool) {
	offset = offset % len(array)
	if offset > 0 {
		result = append(result, array[len(array)-offset:]...)
		result = append(result, array[:len(array)-offset]...)
	}
	return
}

func makeScreen(width, height int) (screen Screen) {
	screen = make(Screen, height)
	for i := 0; i < height; i++ {
		screen[i] = make([]bool, width)
	}
	return
}

func Solve() int {
	screen := makeScreen(50, 6)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		screen.print()
		reader.ReadString('\n')
		operation := strings.Split(scanner.Text(), " ")
		switch operation[0] {
		case "rect":
			var x, y int
			fmt.Sscanf(operation[1], "%dx%d", &x, &y)
			screen.rect(x, y)
		case "rotate":
			if operation[1] == "row" {
				var y int
				fmt.Sscanf(operation[2], "y=%d", &y)
				offset, _ := strconv.Atoi(operation[4])
				screen.rotateRow(y, offset)
			} else if operation[1] == "column" {
				var x int
				fmt.Sscanf(operation[2], "x=%d", &x)
				offset, _ := strconv.Atoi(operation[4])
				screen.rotateColumn(x, offset)
			}
		}
	}

	return screen.countLitPixels()
}
