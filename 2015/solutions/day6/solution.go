package day6

import (
	"bufio"
	"fmt"
	"strings"
)

func parseInstruction(s string) (instruction string, x1, y1, x2, y2 int) {
	if n, _ := fmt.Sscanf(s, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2); n == 4 {
		instruction = "turn on"
	} else if n, _ := fmt.Sscanf(s, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2); n == 4 {
		instruction = "turn off"
	} else if n, _ := fmt.Sscanf(s, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2); n == 4 {
		instruction = "toggle"
	}
	return
}

func countLights(grid [][]bool) (lightsOn int) {
	for _, row := range grid {
		for _, light := range row {
			if light {
				lightsOn++
			}
		}
	}
	return
}

func brightness(grid [][]int) (b int) {
	for _, row := range grid {
		for _, light := range row {
			b += light
		}
	}
	return
}

func SolvePart1() int {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	grid := make([][]bool, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]bool, 1000)
	}

	for scanner.Scan() {
		s := scanner.Text()
		instruction, x1, y1, x2, y2 := parseInstruction(s)

		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				switch instruction {
				case "turn on":
					grid[i][j] = true
				case "turn off":
					grid[i][j] = false
				case "toggle":
					grid[i][j] = !grid[i][j]
				}
			}
		}

		fmt.Printf("%s -> lights on %d\n", s, countLights(grid))
	}

	return countLights(grid)
}

func SolvePart2() int {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}

	for scanner.Scan() {
		s := scanner.Text()
		instruction, x1, y1, x2, y2 := parseInstruction(s)

		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				switch instruction {
				case "turn on":
					grid[i][j]++
				case "turn off":
					grid[i][j]--
					if grid[i][j] < 0 {
						grid[i][j] = 0
					}
				case "toggle":
					grid[i][j] += 2
				}
			}
		}

		fmt.Printf("%s -> total brightness %d\n", s, brightness(grid))
	}

	return brightness(grid)
}
