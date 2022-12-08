package day08

import (
	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func treesLeft(trees [][]byte, x, y int) []byte {
	if x == 0 {
		return []byte{}
	}

	return trees[y][:x]
}

func treesRight(trees [][]byte, x, y int) []byte {
	if x == len(trees[y])-1 {
		return []byte{}
	}

	return trees[y][x+1:]
}

func treesUp(trees [][]byte, x, y int) (tt []byte) {
	for i := 0; i < y; i++ {
		tt = append(tt, trees[i][x])
	}

	return
}

func treesDown(trees [][]byte, x, y int) (tt []byte) {
	for i := y + 1; i < len(trees); i++ {
		tt = append(tt, trees[i][x])
	}

	return
}

func allLower(trees []byte, n byte) bool {
	for _, t := range trees {
		if t >= n {
			return false
		}
	}

	return true
}

func reverse(items []byte) (res []byte) {
	for i := len(items) - 1; i >= 0; i-- {
		res = append(res, items[i])
	}

	return
}

func countVisible(trees []byte, n byte) (count int) {
	for _, t := range trees {
		count++
		if t >= n {
			break
		}
	}

	return
}

func Solve() (part1 int, part2 int) {
	lines := utils.ReadInput("./solutions/day08/input.txt", mappers.ToString)

	trees := make([][]byte, len(lines))

	for i, v := range lines {
		trees[i] = make([]byte, len(lines[i]))

		for j := 0; j < len(v); j++ {
			trees[i][j] = v[j] - byte('0')
		}
	}

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			t := trees[i][j]
			tl := treesLeft(trees, j, i)
			tr := treesRight(trees, j, i)
			tu := treesUp(trees, j, i)
			td := treesDown(trees, j, i)

			if allLower(tl, t) || allLower(tr, t) || allLower(tu, t) || allLower(td, t) {
				part1++
			}

			cl := countVisible(reverse(tl), t)
			cr := countVisible(tr, t)
			cu := countVisible(reverse(tu), t)
			cd := countVisible(td, t)

			score := cl * cr * cu * cd

			if score > part2 {
				part2 = score
			}
		}
	}

	return
}
