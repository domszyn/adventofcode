package day20

const Presents = 29000000

func SolvePart1() (part1 int) {
	houses := make([]int, Presents/10)
	for elf := 1; elf < Presents/10; elf++ {
		for h := elf; h < Presents/10; h += elf {
			houses[h] += elf * 10
		}
	}

	for i, v := range houses {
		part1 = i
		if v > Presents {
			break
		}
	}

	return
}

func SolvePart2() (part2 int) {
	houses := make([]int, Presents/10)
	for elf := 1; elf < Presents/10; elf++ {
		for h, hh := elf, 0; h < Presents/10 && hh < 50; h, hh = h+elf, hh+1 {
			houses[h] += elf * 11
		}
	}

	for i, v := range houses {
		part2 = i
		if v > Presents {
			break
		}
	}

	return
}
