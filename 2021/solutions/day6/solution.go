package day6

func countSpawn(days int, counts map[int]int) (count int) {
	if days < 0 {
		return 1
	}

	if cachedCount, found := counts[days]; found {
		return cachedCount
	}

	for days > 0 {
		if days > 9 {
			count += countSpawn(days-9, counts)
		}
		count++
		days -= 7
	}

	return
}

func Solve() (part1, part2 int) {
	counts := make(map[int]int)

	for i := 0; i < 256; i++ {
		counts[i] = countSpawn(i, counts)
	}

	for _, fish := range Input {
		part1 += counts[80-fish]
		part2 += counts[256-fish]
	}

	return
}
