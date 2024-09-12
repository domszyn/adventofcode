package day14

import (
	"fmt"
	"strings"
)

type Chemical struct {
	Name     string
	Quantity int
	Inputs   []Chemical
}

func parseIngredient(ingredient string) (quantity int, name string) {
	fmt.Sscanf(ingredient, "%d %s", &quantity, &name)
	return
}

func ParseChemical(reaction string) Chemical {
	var chemical Chemical
	parts := strings.Split(reaction, " => ")
	input, output := parts[0], parts[1]
	chemical.Quantity, chemical.Name = parseIngredient(output)
	inputs := strings.Split(input, ", ")

	for _, i := range inputs {
		var input Chemical
		input.Quantity, input.Name = parseIngredient(i)
		chemical.Inputs = append(chemical.Inputs, input)
	}

	return chemical
}
