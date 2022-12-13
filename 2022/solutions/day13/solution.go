package day13

import (
	"sort"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func Solve() (part1 int, part2 int) {
	lines := utils.ReadInput("./solutions/day13/input.txt", mappers.ToString)

	var allPackets Packets

	var pairs [][2]Packet
	for i := 0; i < len(lines); i += 3 {
		p1 := parsePacket(lines[i])
		p2 := parsePacket(lines[i+1])
		pairs = append(pairs, [2]Packet{p1, p2})

		allPackets = append(allPackets, p1, p2)
	}

	allPackets = append(allPackets, parsePacket("[[2]]"), parsePacket("[[6]]"))

	for i := 0; i < len(pairs); i++ {
		if pairs[i][0].Compare(pairs[i][1]) == RightOrder {
			part1 += i + 1
		}
	}

	sort.Sort(allPackets)
	part2 = 1
	for i := 0; i < len(allPackets); i++ {
		if isDivider(allPackets[i], 2) || isDivider(allPackets[i], 6) {
			part2 *= (i + 1)
		}
	}

	return part1, part2
}
