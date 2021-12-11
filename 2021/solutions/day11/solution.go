package day11

import (
	"bufio"
	"strings"
)

type Point struct{ x, y int }

func inc(x, y int, grid [][]byte, flashed map[Point]bool) {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid) {
		return
	}

	if !flashed[Point{x: x, y: y}] {
		grid[y][x]++
		if grid[y][x] > 9 {
			flashed[Point{x: x, y: y}] = true
			inc(x-1, y-1, grid, flashed)
			inc(x, y-1, grid, flashed)
			inc(x+1, y-1, grid, flashed)
			inc(x-1, y, grid, flashed)
			inc(x+1, y, grid, flashed)
			inc(x-1, y+1, grid, flashed)
			inc(x, y+1, grid, flashed)
			inc(x+1, y+1, grid, flashed)
			grid[y][x] = 0
		}
	}
}

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var grid [][]byte
	for scanner.Scan() {
		s := scanner.Text()
		grid = append(grid, make([]byte, len(s)))

		for i, c := range s {
			grid[len(grid)-1][i] = byte(c) - '0'
		}
	}

	for step := 0; step < 1000; step++ {
		flashed := make(map[Point]bool)

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid); x++ {
				inc(x, y, grid, flashed)
			}
		}

		if step < 100 {
			part1 += len(flashed)
		}

		if len(flashed) == len(grid)*len(grid) {
			part2 = step + 1
			break
		}

		flashed = make(map[Point]bool)
	}

	return
}
