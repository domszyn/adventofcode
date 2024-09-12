package day14

func GetInventory(factory Factory) map[string]int {
	done := false
	factory.Inventory["FUEL"] = 0
	for !done {
		done = true

		for ingredient, quantity := range factory.Inventory {
			if ingredient == "ORE" || (ingredient == "FUEL" && quantity > 0) {
				continue
			}
			if quantity < 0 || ingredient == "FUEL" {
				reaction := factory.Reactions[ingredient]
				for _, input := range reaction.Inputs {
					factory.Inventory[input.Name] -= input.Quantity
				}
				factory.Inventory[ingredient] += reaction.Quantity
				done = false
			}
		}
	}

	return factory.Inventory
}

func SolvePart1(input string) int {
	factory := InitFactory(input)
	return -GetInventory(factory)["ORE"]
}

func SolvePart2(input string) int {
	factory := InitFactory(input)
	// inventoryAfterOneFuel := GetInventory(factory)
	factory.Inventory["ORE"] += 1000000000000
	for factory.Inventory["ORE"] > 0 {
		GetInventory(factory)
		for ingredient, quantity := range factory.Inventory {
			if ingredient == "ORE" {
				continue
			}
			if quantity < 0 || ingredient == "FUEL" {
				reaction := factory.Reactions[ingredient]
				for _, input := range reaction.Inputs {
					factory.Inventory[input.Name] -= input.Quantity
				}
				factory.Inventory[ingredient] += reaction.Quantity
			}
		}
	}

	return factory.Inventory["FUEL"]
}

func GetAnswers() (int, int) {
	return SolvePart1(Input), -1
}
