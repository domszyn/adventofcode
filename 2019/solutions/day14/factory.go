package day14

import (
	"math"
	"strings"
)

type Factory struct {
	Inventory map[string]int
	Needed    map[string]int
	Reactions map[string]Chemical
}

func InitFactory(input string) Factory {
	var factory Factory
	factory.Inventory = make(map[string]int)
	factory.Needed = make(map[string]int)
	factory.Reactions = make(map[string]Chemical)

	for reaction := range strings.SplitSeq(input, "\n") {
		chemical := ParseChemical(reaction)
		factory.Reactions[chemical.Name] = chemical
	}

	return factory
}

func (f Factory) CalculateOre() int {
	for {
		var chemical string
		for n := range f.Needed {
			if n != "ORE" {
				chemical = n
				break
			}
		}

		if chemical == "" {
			break
		}

		needed := f.Needed[chemical]
		currentSurplus := f.Inventory[chemical]

		if currentSurplus >= needed {
			f.Inventory[chemical] = currentSurplus - needed
			if f.Inventory[chemical] == 0 {
				delete(f.Inventory, chemical)
			}
			delete(f.Needed, chemical)
			continue
		}

		netNeeded := needed - currentSurplus
		reaction := f.Reactions[chemical]
		numReactions := int(math.Ceil(float64(netNeeded) / float64(reaction.Quantity)))

		for _, input := range reaction.Inputs {
			f.Needed[input.Name] += numReactions * input.Quantity
		}

		amountProduced := numReactions * reaction.Quantity
		f.Inventory[chemical] = currentSurplus + amountProduced - needed
		if f.Inventory[chemical] == 0 {
			delete(f.Inventory, chemical)
		}
		delete(f.Needed, chemical)
	}

	return f.Needed["ORE"]
}
