package day08

import (
	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func splitAt(slice []int, n int) ([]int, []int) {
	return slice[:n], slice[n+1:]
}

type Map = [][]int
type Tree struct {
	Height int
}

func (tree Tree) IsHigherThan(trees []int) bool {
	for _, t := range trees {
		if t >= tree.Height {
			return false
		}
	}

	return true
}

func (tree Tree) CountVisible(trees []int) (count int) {
	for i := 0; i < len(trees); i++ {
		count++
		if trees[i] >= tree.Height {
			break
		}
	}

	return
}

func (tree Tree) CountVisibleInReverse(trees []int) (count int) {
	for i := len(trees) - 1; i >= 0; i-- {
		count++
		if trees[i] >= tree.Height {
			break
		}
	}

	return
}

func initMaps() (map1, map2 Map) {
	lines := utils.ReadInput("./solutions/day08/input.txt", mappers.ToString)
	height := len(lines)
	width := len(lines[0])
	map1 = make(Map, height)
	map2 = make(Map, width)

	for y := 0; y < len(lines); y++ {
		map1[y] = make([]int, width)

		for x := 0; x < len(lines[y]); x++ {
			map1[y][x] = int(lines[y][x] - '0')

			if len(map2[x]) != height {
				map2[x] = make([]int, height)
			}

			map2[x][y] = int(lines[y][x] - '0')
		}
	}

	return
}

func Solve() (part1 int, part2 int) {
	trees, trees90 := initMaps()

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			tree := Tree{trees[i][j]}
			left, right := splitAt(trees[i], j)
			up, down := splitAt(trees90[j], i)

			if tree.IsHigherThan(left) || tree.IsHigherThan(right) || tree.IsHigherThan(up) || tree.IsHigherThan(down) {
				part1++
			}

			if i > 0 && i < len(trees)-1 && j > 0 && j < len(trees90)-1 {
				cl := tree.CountVisibleInReverse(left)
				cr := tree.CountVisible(right)
				cu := tree.CountVisibleInReverse(up)
				cd := tree.CountVisible(down)

				score := cl * cr * cu * cd

				if score > part2 {
					part2 = score
				}
			}
		}
	}

	return
}
