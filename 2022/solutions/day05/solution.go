package day05

import (
	"fmt"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func getStacks() (stacks []string) {
	stackLines := utils.ReadInput("./solutions/day05/stacks.txt", mappers.ToString)

	for i := len(stackLines) - 1; i >= 0; i-- {
		line := stackLines[i]

		for j := 0; j < len(line); j += 4 {
			elem := string(line[j+1])

			if elem == " " {
				continue
			}

			if (j+1)/4 >= len(stacks) {
				stacks = append(stacks, elem)
			} else {
				stacks[(j+1)/4] += elem
			}
		}
	}

	return
}

type Crane struct {
	Stacks []string
}

func parseMoveInstruction(move string) (n, from, to int) {
	fmt.Sscanf(move, "move %d from %d to %d", &n, &from, &to)
	return
}

func (c *Crane) Load(stacks []string) {
	c.Stacks = make([]string, len(stacks))
	copy(c.Stacks, stacks)
}

func (c *Crane) MoveOneByOne(nCrates int, from, to int) {
	var removedCrates string
	splitAt := len(c.Stacks[from-1]) - nCrates
	removedCrates, c.Stacks[from-1] = c.Stacks[from-1][splitAt:], c.Stacks[from-1][:splitAt]

	for i := len(removedCrates) - 1; i >= 0; i-- {
		c.Stacks[to-1] += string(removedCrates[i])
	}
}

func (c *Crane) MoveAll(nCrates int, from, to int) {
	var removedCrates string
	splitAt := len(c.Stacks[from-1]) - nCrates
	removedCrates, c.Stacks[from-1] = c.Stacks[from-1][splitAt:], c.Stacks[from-1][:splitAt]

	c.Stacks[to-1] += removedCrates
}

func (c *Crane) ReadStackTops() (result string) {
	for _, s := range c.Stacks {
		result += string(s[len(s)-1])
	}

	return
}

func Solve() (string, string) {
	moves := utils.ReadInput("./solutions/day05/moves.txt", mappers.ToString)
	stacks := getStacks()

	var crane1, crane2 Crane

	crane1.Load(stacks)
	crane2.Load(stacks)

	for _, move := range moves {
		n, from, to := parseMoveInstruction(move)

		crane1.MoveOneByOne(n, from, to)
		crane2.MoveAll(n, from, to)
	}

	return crane1.ReadStackTops(), crane2.ReadStackTops()
}
