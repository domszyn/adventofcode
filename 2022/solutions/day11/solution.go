package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

type Monkey struct {
	Items     []int
	Inspected int
	Divisible int
	IfTrue    int
	IfFalse   int
	Operand   string
	Change    string
}

type Operation = func(int) int
type Test = func(int) int

func (m *Monkey) Inspect(old int) (new int) {
	if m.Change == "old" {
		new = old
	} else {
		new, _ = strconv.Atoi(m.Change)
	}
	switch m.Operand {
	case "*":
		new *= old
	case "+":
		new += old
	}

	return
}

func getMonkeys() []*Monkey {
	lines := utils.ReadInput("./solutions/day11/input.txt", mappers.ToString)

	var monkeys []*Monkey

	var monkey *Monkey
	for i, line := range lines {
		switch i % 7 {
		case 0:
			monkey = &Monkey{}
		case 1:
			items := strings.Split(line[18:], ", ")
			for _, i := range items {
				item, _ := strconv.Atoi(i)
				monkey.Items = append(monkey.Items, item)
			}
		case 2:
			monkey.Operand = line[23:24]
			monkey.Change = line[25:]
		case 3:
			fmt.Sscanf(line[8:], "divisible by %d", &monkey.Divisible)
		case 4:
			fmt.Sscanf(line[13:], "throw to monkey %d", &monkey.IfTrue)
		case 5:
			fmt.Sscanf(line[14:], "throw to monkey %d", &monkey.IfFalse)
			monkeys = append(monkeys, monkey)
		}
	}

	return monkeys
}

func monkeyBusiness(monkeys []*Monkey, numRounds int, worryRelief func(int) int) int {
	for round := 0; round < numRounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			m := monkeys[i]
			for j := 0; j < len(m.Items); j++ {
				item := worryRelief(m.Inspect(m.Items[j]))

				if item%m.Divisible == 0 {
					monkeys[m.IfTrue].Items = append(monkeys[m.IfTrue].Items, item)
				} else {
					monkeys[m.IfFalse].Items = append(monkeys[m.IfFalse].Items, item)
				}
				m.Inspected++
			}
			m.Items = []int{}
		}
	}

	var inspections []int
	for i := 0; i < len(monkeys); i++ {
		inspections = append(inspections, monkeys[i].Inspected)
	}

	largest := utils.NLargest(2, inspections)

	return largest[0] * largest[1]
}

func Solve() (part1 int, part2 int) {
	monkeys := getMonkeys()
	part1 = monkeyBusiness(monkeys, 20, func(i int) int { return i / 3 })

	monkeys = getMonkeys()
	s := 1
	for _, m := range monkeys {
		s *= m.Divisible
	}
	part2 = monkeyBusiness(monkeys, 10000, func(i int) int { return i % s })
	return
}
