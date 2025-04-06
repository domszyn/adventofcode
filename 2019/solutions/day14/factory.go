package day14

import "strings"

type Factory struct {
	Inventory map[string]int
	Reactions map[string]Chemical
}

func InitFactory(input string) Factory {
	var factory Factory
	factory.Inventory = make(map[string]int)
	factory.Reactions = make(map[string]Chemical)

	for _, reaction := range strings.Split(input, "\n") {
		chemical := ParseChemical(reaction)
		factory.Reactions[chemical.Name] = chemical
	}

	return factory
}
