package day15

import (
	"bufio"
	"fmt"
	"strings"
)

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

func parseIngredient(s string) (i Ingredient) {
	fmt.Sscanf(s,
		"%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
		&i.Name, &i.Capacity, &i.Durability, &i.Flavor, &i.Texture, &i.Calories)
	return
}

func Solve() (int, int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	var ingredients []Ingredient

	for scanner.Scan() {
		s := scanner.Text()
		ingredient := parseIngredient(s)
		ingredients = append(ingredients, ingredient)
	}

	recipes := mixIngredients(len(ingredients), 100)

	maxScore := 0
	maxScore500Cal := 0
	for _, r := range recipes {
		score := score(ingredients, r)
		calories := totalCalories(ingredients, r)
		if score > maxScore {
			maxScore = score
		}

		if calories == 500 && score > maxScore500Cal {
			maxScore500Cal = score
		}
	}

	return maxScore, maxScore500Cal
}

func mixIngredients(numIngredients, total int) (mixes [][]int) {
	switch numIngredients {
	case 1:
		mixes = make([][]int, 0, 1)
	case 2:
		mixes = make([][]int, 0, 101)
	case 3:
		mixes = make([][]int, 0, 5151)
	case 4:
		mixes = make([][]int, 0, 176851)
	}
	if numIngredients == 1 {
		mixes = append(mixes, []int{total})
		return
	}

	for i := 0; i <= total; i++ {
		submixes := mixIngredients(numIngredients-1, total-i)
		for _, sm := range submixes {
			mix := make([]int, 0, numIngredients)
			mix = append(mix, i)
			mix = append(mix, sm...)
			mixes = append(mixes, mix)
		}
	}

	return
}

func sum(ingredients []int) (sum int) {
	for _, i := range ingredients {
		sum += i
	}
	return
}

func score(ingredients []Ingredient, recipe []int) int {
	var mix Ingredient
	for i, ingredient := range ingredients {
		mix.Capacity += recipe[i] * ingredient.Capacity
		mix.Durability += recipe[i] * ingredient.Durability
		mix.Flavor += recipe[i] * ingredient.Flavor
		mix.Texture += recipe[i] * ingredient.Texture
		mix.Calories += recipe[i] * ingredient.Calories
	}

	if mix.Capacity < 0 {
		mix.Capacity = 0
	}

	if mix.Durability < 0 {
		mix.Durability = 0
	}

	if mix.Flavor < 0 {
		mix.Flavor = 0
	}

	if mix.Texture < 0 {
		mix.Texture = 0
	}

	if mix.Calories < 0 {
		mix.Calories = 0
	}

	return mix.Capacity * mix.Durability * mix.Flavor * mix.Texture
}

func totalCalories(ingredients []Ingredient, recipe []int) (calories int) {
	for i, ingredient := range ingredients {
		calories += recipe[i] * ingredient.Calories
	}

	return
}
