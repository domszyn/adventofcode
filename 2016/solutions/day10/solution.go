package day10

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type BotSwarm map[int]*Bot
type OutputBins map[int]int

type Bot struct {
	Chips   []int
	Proceed func(f *Factory)
}

type Factory struct {
	Bots   BotSwarm
	Output OutputBins
}

func (f *Factory) CreateBot(index int) {
	if _, found := f.Bots[index]; !found {
		f.Bots[index] = &Bot{
			Chips: []int{},
		}
	}
}

func (f *Factory) OutputReady() bool {
	_, found0 := f.Output[0]
	_, found1 := f.Output[1]
	_, found2 := f.Output[2]
	return found0 && found1 && found2
}

func Solve() (part1, part2 int) {
	factory := Factory{
		Bots:   make(BotSwarm),
		Output: make(OutputBins),
	}
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		id, _ := strconv.Atoi(instruction[1])
		switch instruction[0] {
		case "value":
			idx, _ := strconv.Atoi(instruction[5])
			factory.CreateBot(idx)
			factory.Bots[idx].Chips = append(factory.Bots[idx].Chips, id)
		case "bot":
			factory.CreateBot(id)
			lowValue, _ := strconv.Atoi(instruction[6])
			highValue, _ := strconv.Atoi(instruction[11])
			lowTo := instruction[5]
			highTo := instruction[10]
			factory.Bots[id].Proceed = func(f *Factory) {
				chips := f.Bots[id].Chips
				if len(chips) < 2 {
					return
				}

				lo, hi := chips[0], chips[1]
				f.Bots[id].Chips = []int{}
				if lo > hi {
					lo, hi = hi, lo
				}

				if lo == 17 && hi == 61 {
					part1 = id
				}

				switch lowTo {
				case "bot":
					f.Bots[lowValue].Chips = append(f.Bots[lowValue].Chips, lo)
				case "output":
					f.Output[lowValue] = lo
				}

				switch highTo {
				case "bot":
					f.Bots[highValue].Chips = append(f.Bots[highValue].Chips, hi)
				case "output":
					f.Output[highValue] = hi
				}
			}
		}
	}

	for i := 0; ; i++ {
		for j := 0; j < len(factory.Bots); j++ {
			if len(factory.Bots[j].Chips) == 2 {
				fmt.Printf("%d\t", j)
				factory.Bots[j].Proceed(&factory)
			}
		}

		if len(factory.Output) == 21 {
			part2 = factory.Output[0] * factory.Output[1] * factory.Output[2]
			break
		}
	}

	return
}
