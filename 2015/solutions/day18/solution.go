package day18

import (
	"bufio"
	"strings"
)

type Grid map[int]map[int]bool

func parseInput(input string) Grid {
	grid := make(Grid)
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	row := 0
	for scanner.Scan() {
		rowText := scanner.Text()

		grid[row] = make(map[int]bool, len(rowText))
		for i, char := range rowText {
			switch char {
			case '.':
				grid[row][i] = false
			case '#':
				grid[row][i] = true
			}
		}
		row++
	}

	return grid
}

func countNeighbors(grid Grid, row, column int) (count int) {
	if prevRow, ok := grid[row-1]; ok {
		if prevRow[column-1] {
			count++
		}
		if prevRow[column] {
			count++
		}
		if prevRow[column+1] {
			count++
		}
	}
	if grid[row][column-1] {
		count++
	}
	if grid[row][column+1] {
		count++
	}
	if nextRow, ok := grid[row+1]; ok {
		if nextRow[column-1] {
			count++
		}
		if nextRow[column] {
			count++
		}
		if nextRow[column+1] {
			count++
		}
	}
	return
}

func countLights(grid Grid) (count int) {
	for _, row := range grid {
		for _, v := range row {
			if v {
				count++
			}
		}
	}

	return
}

func switchLights(grid Grid) Grid {
	size := len(grid)
	nextGrid := make(Grid)
	for i := 0; i < size; i++ {
		nextGrid[i] = make(map[int]bool)
		for j := 0; j < size; j++ {
			neighbors := countNeighbors(grid, i, j)

			if grid[i][j] {
				nextGrid[i][j] = neighbors == 2 || neighbors == 3
			} else {
				nextGrid[i][j] = neighbors == 3
			}
		}
	}

	return nextGrid
}

func Solve() (part1, part2 int) {
	grid := parseInput(Input)

	for i := 0; i < 100; i++ {
		grid = switchLights(grid)
	}

	part1 = countLights(grid)

	grid = parseInput(Input)
	grid[0][0] = true
	grid[0][len(grid)-1] = true
	grid[len(grid)-1][0] = true
	grid[len(grid)-1][len(grid)-1] = true

	for i := 0; i < 100; i++ {
		grid = switchLights(grid)
		grid[0][0] = true
		grid[0][len(grid)-1] = true
		grid[len(grid)-1][0] = true
		grid[len(grid)-1][len(grid)-1] = true
	}

	part2 = countLights(grid)

	return
}
