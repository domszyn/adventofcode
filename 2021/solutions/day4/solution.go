package day4

import (
	"bufio"
	"fmt"
	"strings"
)

type BingoBoard struct {
	Numbers [][]int
}

func (board *BingoBoard) call(n int) bool {
	for i, row := range board.Numbers {
		for j, value := range row {
			if n == value {
				board.Numbers[i][j] = 0
				return true
			}
		}
	}

	return false
}

func (board BingoBoard) sumUnmarked() (sum int) {
	for _, row := range board.Numbers {
		for _, number := range row {
			if number > 0 {
				sum += number - 1
			}
		}
	}

	return
}

func (board BingoBoard) hasWonRow(row int) bool {
	for _, number := range board.Numbers[row] {
		if number > 0 {
			return false
		}
	}
	return true
}

func (board BingoBoard) hasWonColumn(column int) bool {
	for _, row := range board.Numbers {
		if row[column] > 0 {
			return false
		}
	}
	return true
}

func (board BingoBoard) hasWonAnyRowOrColumn() bool {
	for i := range board.Numbers {
		if board.hasWonRow(i) || board.hasWonColumn(i) {
			return true
		}
	}
	return false
}

func (board BingoBoard) hasWonDiagonal() bool {
	for i, row := range board.Numbers {
		if row[i] > 0 {
			return false
		}
	}
	return true
}

func (board BingoBoard) hasWonOtherDiagonal() bool {
	for i, row := range board.Numbers {
		if row[len(board.Numbers)-i-1] > 0 {
			return false
		}
	}
	return true
}

func (board BingoBoard) hasWon() bool {
	return board.hasWonAnyRowOrColumn() ||
		board.hasWonDiagonal() ||
		board.hasWonOtherDiagonal()
}

func makeBoard() BingoBoard {
	return BingoBoard{
		Numbers: make([][]int, 5),
	}
}

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var boards []BingoBoard
	var i int
	var board BingoBoard
	for scanner.Scan() {
		s := scanner.Text()
		switch i {
		case 0:
			board = makeBoard()
		case 5:
			boards = append(boards, board)
			i = 0
			continue
		}

		board.Numbers[i] = make([]int, 5)
		var n1, n2, n3, n4, n5 int
		fmt.Sscanf(s, "%d %d %d %d %d", &n1, &n2, &n3, &n4, &n5)
		board.Numbers[i] = []int{n1 + 1, n2 + 1, n3 + 1, n4 + 1, n5 + 1}
		i++
	}

	boardsWon := make(map[int]bool)
	for _, n := range Numbers {
		for bi, board := range boards {
			dabbed := board.call(n + 1)
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
