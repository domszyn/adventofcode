package day4

import (
	"bufio"
	"fmt"
	"strings"
)

type BingoNumber struct {
	Value  int
	Called bool
}

type BingoBoard struct {
	Numbers [][]BingoNumber
}

func (board *BingoBoard) call(n int) bool {
	for i := 0; i < len(board.Numbers); i++ {
		for j := 0; j < len(board.Numbers); j++ {
			if board.Numbers[i][j].Value == n {
				board.Numbers[i][j].Called = true
				return true
			}
		}
	}

	return false
}

func (board BingoBoard) sumUnmarked() (sum int) {
	for _, row := range board.Numbers {
		for _, number := range row {
			if !number.Called {
				sum += number.Value
			}
		}
	}

	return
}

func (board BingoBoard) hasWonRow(row int) bool {
	won := true
	for j := 0; j < len(board.Numbers); j++ {
		if !board.Numbers[row][j].Called {
			won = false
			break
		}
	}
	return won
}

func (board BingoBoard) hasWonAnyRow() bool {
	for j := 0; j < len(board.Numbers); j++ {
		if board.hasWonRow(j) {
			return true
		}
	}
	return false
}

func (board BingoBoard) hasWonColumn(column int) bool {
	won := true
	for j := 0; j < len(board.Numbers); j++ {
		if !board.Numbers[j][column].Called {
			won = false
			break
		}
	}
	return won
}

func (board BingoBoard) hasWonAnyColumn() bool {
	for j := 0; j < len(board.Numbers); j++ {
		if board.hasWonColumn(j) {
			return true
		}
	}
	return false
}

func (board BingoBoard) hasWonDiagonal() bool {
	won := true
	for j := 0; j < len(board.Numbers); j++ {
		if !board.Numbers[j][j].Called {
			won = false
			break
		}
	}
	return won
}

func (board BingoBoard) hasWonOtherDiagonal() bool {
	won := true
	for j := 0; j < len(board.Numbers); j++ {
		if !board.Numbers[j][len(board.Numbers)-j-1].Called {
			won = false
			break
		}
	}
	return won
}

func (board BingoBoard) hasWon() bool {
	return board.hasWonAnyColumn() ||
		board.hasWonAnyRow() ||
		board.hasWonDiagonal() ||
		board.hasWonOtherDiagonal()
}

func Solve() (part1, part2 int) {
	numbers := []int{63, 23, 2, 65, 55, 94, 38, 20, 22, 39, 5, 98, 9, 60, 80, 45, 99, 68, 12, 3, 6, 34, 64, 10, 70, 69, 95, 96, 83, 81, 32, 30, 42, 73, 52, 48, 92, 28, 37, 35, 54, 7, 50, 21, 74, 36, 91, 97, 13, 71, 86, 53, 46, 58, 76, 77, 14, 88, 78, 1, 33, 51, 89, 26, 27, 31, 82, 44, 61, 62, 75, 66, 11, 93, 49, 43, 85, 0, 87, 40, 24, 29, 15, 59, 16, 67, 19, 72, 57, 41, 8, 79, 56, 4, 18, 17, 84, 90, 47, 25}

	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var boards []BingoBoard
	var i int
	board := BingoBoard{
		Numbers: make([][]BingoNumber, 5),
	}
	for scanner.Scan() {
		s := scanner.Text()
		if i == 5 {
			boards = append(boards, board)
			board = BingoBoard{
				Numbers: make([][]BingoNumber, 5),
			}
			i = 0
			continue
		}

		board.Numbers[i] = make([]BingoNumber, 5)
		var n1, n2, n3, n4, n5 int
		fmt.Sscanf(s, "%d %d %d %d %d", &n1, &n2, &n3, &n4, &n5)
		board.Numbers[i] = []BingoNumber{
			{Value: n1, Called: false},
			{Value: n2, Called: false},
			{Value: n3, Called: false},
			{Value: n4, Called: false},
			{Value: n5, Called: false},
		}
		i++
	}

	boardsWon := make(map[int]bool)
	for _, n := range numbers {
		for bi, board := range boards {
			dabbed := board.call(n)
			if dabbed && board.hasWon() {
				if part1 == 0 {
					part1 = n * board.sumUnmarked()
				}
				boardsWon[bi] = true
				if len(boardsWon) == len(boards) {
					part2 = n * board.sumUnmarked()
					return
				}
			}
		}
	}

	return
}
