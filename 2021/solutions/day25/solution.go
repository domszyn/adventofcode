package day25

import (
	"bufio"
	"strings"
)

func canMoveEast(rows [][]string) bool {
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[0]); x++ {
			posEast := (x + 1) % len(rows[0])
			if rows[y][x] == ">" && rows[y][posEast] == "." {
				return true
			}
		}
	}

	return false
}

func canMoveSouth(rows [][]string) bool {
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[0]); x++ {
			posSouth := (y + 1) % len(rows)
			if rows[y][x] == "v" && rows[posSouth][x] == "." {
				return true
			}
		}
	}

	return false
}

func moveEast(rows [][]string) (result [][]string) {
	for y := 0; y < len(rows); y++ {
		row := make([]string, len(rows[0]))
		for x := 0; x < len(rows[0]); x++ {
			posEast := (x + 1) % len(rows[0])
			if rows[y][x] == ">" && rows[y][posEast] == "." {
				row[x] = "."
				row[posEast] = ">"
				x++
			} else {
				row[x] = rows[y][x]
			}
		}
		result = append(result, row)
	}

	return
}

func moveSouth(rows [][]string) (result [][]string) {
	for y := 0; y < len(rows); y++ {
		row := make([]string, len(rows[0]))
		for x := 0; x < len(rows[0]); x++ {
			row[x] = rows[y][x]
		}
		result = append(result, row)
	}

	for x := 0; x < len(rows[0]); x++ {
		for y := 0; y < len(rows); y++ {
			posSouth := (y + 1) % len(rows)
			if rows[y][x] == "v" && rows[posSouth][x] == "." {
				result[y][x] = "."
				result[posSouth][x] = "v"
				y++
			}
		}
	}

	return
}

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var rows [][]string
	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for _, v := range line {
			row = append(row, string(v))
		}
		rows = append(rows, row)
	}

	for step := 1; ; step++ {
		if !canMoveEast(rows) && !canMoveSouth(rows) {
			part1 = step
			break
		}

		rows = moveEast(rows)
		rows = moveSouth(rows)
	}

	return
}
