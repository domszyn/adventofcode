package day11

import (
	"bufio"
	"fmt"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

type Point struct{ x, y int }

func initAnimation(grid [][]byte) {
	tm.Clear()
	tm.MoveCursor(9, 2)
	tm.Println(tm.Bold(tm.Color(fmt.Sprintf("STEP 001"), tm.WHITE)))
	printBox()
	tm.MoveCursor(40, 20)
	tm.Flush()
	printGrid(grid)
}

func printStep(step int) {
	tm.MoveCursor(9, 2)
	tm.Println(tm.Bold(tm.Color(fmt.Sprintf("STEP %03d", step+1), tm.WHITE)))
	tm.MoveCursor(40, 20)
	tm.Flush()
}

func printBox() {
	for x := 0; x < 22; x++ {
		for y := 0; y < 12; y++ {
			tm.MoveCursor(x+2, y+3)
			if x == 0 && y == 0 {
				tm.Println(tm.Bold(tm.Color("╔", tm.WHITE)))
			} else if x == 21 && y == 0 {
				tm.Println(tm.Bold(tm.Color("╗", tm.WHITE)))
			} else if x == 0 && y == 11 {
				tm.Println(tm.Bold(tm.Color("╚", tm.WHITE)))
			} else if x == 21 && y == 11 {
				tm.Println(tm.Bold(tm.Color("╝", tm.WHITE)))
			} else if x == 0 || x == 21 {
				tm.Println(tm.Bold(tm.Color("║", tm.WHITE)))
			} else if y == 0 || y == 11 {
				tm.Println(tm.Bold(tm.Color("══", tm.WHITE)))
			}
			tm.Flush()
		}
	}
	tm.Flush()
}

func printCell(x, y int, grid [][]byte) {
	tm.MoveCursor(x*2+3, y+4)
	switch grid[y][x] {
	case 0:
		tm.Println(" ")
	case 1:
		tm.Println(tm.Color("ᐧ", tm.MAGENTA))
	case 2:
		tm.Println(tm.Bold(tm.Color("·", tm.MAGENTA)))
	case 3:
		tm.Println(tm.Bold(tm.Color("•", tm.BLUE)))
	case 4:
		tm.Println(tm.Color("○", tm.GREEN))
	case 5:
		tm.Println(tm.Bold(tm.Color("⧲", tm.GREEN)))
	case 6:
		tm.Println(tm.Color("●", tm.YELLOW))
	case 7:
		tm.Println(tm.Color("⦿", tm.RED))
	case 8:
		tm.Println(tm.Bold(tm.Color("◉", tm.RED)))
	case 9:
		tm.Println("🐙")
	default:
		tm.Println("💥")
	}
	tm.MoveCursor(40, 20)
	tm.Flush()
	time.Sleep(time.Millisecond * 17)
}

func printGrid(grid [][]byte) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			printCell(x, y, grid)
		}
	}
}

func inc(x, y int, grid [][]byte, flashed map[Point]bool, animate bool) {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid) {
		return
	}

	if !flashed[Point{x: x, y: y}] {
		grid[y][x]++

		if animate {
			printCell(x, y, grid)
		}

		if grid[y][x] > 9 {
			flashed[Point{x: x, y: y}] = true
			inc(x-1, y-1, grid, flashed, animate)
			inc(x, y-1, grid, flashed, animate)
			inc(x+1, y-1, grid, flashed, animate)
			inc(x-1, y, grid, flashed, animate)
			inc(x+1, y, grid, flashed, animate)
			inc(x-1, y+1, grid, flashed, animate)
			inc(x, y+1, grid, flashed, animate)
			inc(x+1, y+1, grid, flashed, animate)
			grid[y][x] = 0

			if animate {
				printCell(x, y, grid)
			}
		}
	}
}

func Solve(animate bool) (part1, part2 int) {
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

	if animate {
		initAnimation(grid)
	}

	for step := 0; step < 310; step++ {
		if animate {
			printStep(step)
		}

		flashed := make(map[Point]bool)

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid); x++ {
				inc(x, y, grid, flashed, animate)
			}
		}

		if step < 100 {
			part1 += len(flashed)
		}

		if len(flashed) == len(grid)*len(grid) && part2 == 0 {
			part2 = step + 1
			if !animate {
				break
			}
		}

		flashed = make(map[Point]bool)
	}

	return
}
