package day13

import (
	"bufio"
	"fmt"
	"strings"
)

func parseHappiness(s string) (from, to string, happiness int) {
	var gainOrLose string
	fmt.Sscanf(s, "%s would %s %d happiness units by sitting next to %s.", &from, &gainOrLose, &happiness, &to)
	to = to[:len(to)-1]
	if gainOrLose == "lose" {
		happiness *= -1
	}
	return
}

type Happiness struct {
	From      string
	To        string
	Happiness int
}

func Solve(includeMyself bool) int {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	seats := make(map[string]int)

	var felicita []Happiness

	for scanner.Scan() {
		s := scanner.Text()
		from, to, happiness := parseHappiness(s)
		felicita = append(felicita, Happiness{
			From:      from,
			To:        to,
			Happiness: happiness,
		})
	}

	var counter int
	for _, f := range felicita {
		if _, ok := seats[f.From]; !ok {
			seats[f.From] = counter
			counter++
		}
		if _, ok := seats[f.To]; !ok {
			seats[f.To] = counter
			counter++
		}
	}

	if includeMyself {
		seats["Dmytro"] = counter
	}

	grid := make([][]int, len(seats))
	for i := 0; i < len(seats); i++ {
		grid[i] = make([]int, len(seats))
	}

	for _, f := range felicita {
		from := seats[f.From]
		to := seats[f.To]
		grid[from][to] = f.Happiness
	}

	initialElements := []int{0, 1, 2, 3, 4, 5, 6, 7}
	if includeMyself {
		initialElements = append(initialElements, 8)
	}

	permutations := permutate(initialElements)

	maxHappiness := 0
	for _, p := range permutations {
		happiness := calculateHappiness(p, grid)
		if happiness > maxHappiness {
			maxHappiness = happiness
		}
	}

	return maxHappiness
}

func permutate(elements []int) (permutations [][]int) {
	c := make([]int, len(elements))
	permutation := make([]int, len(elements))
	copy(permutation, elements)
	permutations = append(permutations, permutation)
	for i := 0; i < len(elements); {
		if c[i] < i {
			if i%2 == 0 {
				elements[0], elements[i] = elements[i], elements[0]
			} else {
				elements[c[i]], elements[i] = elements[i], elements[c[i]]
			}
			permutation = make([]int, len(elements))
			copy(permutation, elements)
			permutations = append(permutations, permutation)
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}

	return
}

func calculateHappiness(seats []int, felicita [][]int) (happiness int) {
	seats = append(seats, seats[0])
	for i := 0; i < len(seats)-1; i++ {
		happiness += felicita[seats[i]][seats[i+1]]
		happiness += felicita[seats[i+1]][seats[i]]
	}
	return
}
