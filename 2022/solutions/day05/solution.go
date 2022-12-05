package day05

import (
	"fmt"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func getStacks() []string {
	return []string{
		"QSWCZVFT",
		"QRB",
		"BZTQPMS",
		"DVFRQH",
		"JGLDBSTP",
		"WRTZ",
		"HQMNSFRJ",
		"RNFHW",
		"JZTQPRB",
	}
}

func parseMoveInstruction(move string) (n, from, to int) {
	fmt.Sscanf(move, "move %d from %d to %d", &n, &from, &to)
	from--
	to--
	return
}

func Solve() (part1 string, part2 string) {
	moves := utils.ReadInput("./solutions/day05/moves.txt", mappers.ToString)
	stacks1 := getStacks()
	stacks2 := getStacks()

	for _, move := range moves {

		n, from, to := parseMoveInstruction(move)

		removed1 := stacks1[from][len(stacks1[from])-n:]
		removed2 := stacks2[from][len(stacks2[from])-n:]
		stacks1[from] = stacks1[from][:len(stacks1[from])-n]
		stacks2[from] = stacks2[from][:len(stacks2[from])-n]

		for i := len(removed1) - 1; i >= 0; i-- {
			stacks1[to] += string(removed1[i])
		}

		stacks2[to] += removed2
	}

	for _, s := range stacks1 {
		part1 += string(s[len(s)-1])
	}

	for _, s := range stacks2 {
		part2 += string(s[len(s)-1])
	}

	return
}
