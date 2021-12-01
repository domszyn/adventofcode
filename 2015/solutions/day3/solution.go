package day3

type House struct {
	X int
	Y int
}

func CountHouses() int {
	currentHouse := House{X: 0, Y: 0}
	visitedHouses := make(map[House]bool)
	visitedHouses[currentHouse] = true

	for _, move := range Input {
		switch move {
		case '>':
			currentHouse.X++
		case '<':
			currentHouse.X--
		case '^':
			currentHouse.Y++
		case 'v':
			currentHouse.Y--
		}
		visitedHouses[currentHouse] = true
	}

	return len(visitedHouses)
}

func CountHousesWithRoboSanta() int {
	santa := House{X: 0, Y: 0}
	roboSanta := House{X: 0, Y: 0}
	visitedHouses := make(map[House]bool)
	visitedHouses[santa] = true

	for i, move := range Input {
		var loc *House

		if i%2 == 0 {
			loc = &santa
		} else {
			loc = &roboSanta
		}

		switch move {
		case '>':
			loc.X++
		case '<':
			loc.X--
		case '^':
			loc.Y++
		case 'v':
			loc.Y--
		}
		visitedHouses[*loc] = true
	}

	return len(visitedHouses)
}
