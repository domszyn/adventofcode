package day16

import (
	"bufio"
	"fmt"
	"strings"
)

type Sue struct {
	Number      int
	Children    int
	Cats        int
	Samoyeds    int
	Pomeranians int
	Akitas      int
	Vizslas     int
	Goldfish    int
	Trees       int
	Cars        int
	Perfumes    int
}

func parseSue(s string) (i Sue) {
	parts := strings.SplitN(s, ": ", 2)
	fmt.Sscanf(parts[0], "Sue %d", &i.Number)
	i.Children = -1
	i.Cats = -1
	i.Samoyeds = -1
	i.Pomeranians = -1
	i.Akitas = -1
	i.Vizslas = -1
	i.Goldfish = -1
	i.Trees = -1
	i.Cars = -1
	i.Perfumes = -1
	compounds := strings.Split(parts[1], ", ")
	var name string
	var quantity int
	for _, c := range compounds {
		fmt.Sscanf(c, "%s %d", &name, &quantity)
		name = name[:len(name)-1]
		switch name {
		case "children":
			i.Children = quantity
		case "cats":
			i.Cats = quantity
		case "samoyeds":
			i.Samoyeds = quantity
		case "pomeranians":
			i.Pomeranians = quantity
		case "akitas":
			i.Akitas = quantity
		case "vizslas":
			i.Vizslas = quantity
		case "goldfish":
			i.Goldfish = quantity
		case "trees":
			i.Trees = quantity
		case "cars":
			i.Cars = quantity
		case "perfumes":
			i.Perfumes = quantity
		}
	}
	return
}

func SolvePart1() int {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		aunt := parseSue(s)

		if (aunt.Children == 3 || aunt.Children == -1) &&
			(aunt.Cats == 7 || aunt.Cats == -1) &&
			(aunt.Samoyeds == 2 || aunt.Samoyeds == -1) &&
			(aunt.Pomeranians == 3 || aunt.Pomeranians == -1) &&
			(aunt.Akitas == 0 || aunt.Akitas == -1) &&
			(aunt.Vizslas == 0 || aunt.Vizslas == -1) &&
			(aunt.Goldfish == 5 || aunt.Goldfish == -1) &&
			(aunt.Trees == 3 || aunt.Trees == -1) &&
			(aunt.Cars == 2 || aunt.Cars == -1) &&
			(aunt.Perfumes == 1 || aunt.Perfumes == -1) {
			return aunt.Number
		}
	}

	return 0
}

func SolvePart2() int {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		aunt := parseSue(s)

		if (aunt.Children == 3 || aunt.Children == -1) &&
			(aunt.Cats > 7 || aunt.Cats == -1) &&
			(aunt.Samoyeds == 2 || aunt.Samoyeds == -1) &&
			(aunt.Pomeranians < 3 || aunt.Pomeranians == -1) &&
			(aunt.Akitas == 0 || aunt.Akitas == -1) &&
			(aunt.Vizslas == 0 || aunt.Vizslas == -1) &&
			(aunt.Goldfish < 5 || aunt.Goldfish == -1) &&
			(aunt.Trees > 3 || aunt.Trees == -1) &&
			(aunt.Cars == 2 || aunt.Cars == -1) &&
			(aunt.Perfumes == 1 || aunt.Perfumes == -1) {
			return aunt.Number
		}
	}

	return 0
}
