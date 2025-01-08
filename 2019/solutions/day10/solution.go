package day10

import "fmt"

const (
	EmptySpace = '.'
	Asteriod   = '#'
)

type AsteroidField = [][]rune

func Print(af AsteroidField) {
	s := ""
	for _, row := range af {
		s += "\n" + string(row)
	}
	fmt.Println(s)
}

func detectAsteroid(af AsteroidField, x, y, dx, dy int) int {
	if y+dy < 0 || y+dy >= len(af) {
		return 0
	}

	if x+dx < 0 || x+dx >= len(af[0]) {
		return 0
	}

	if af[y+dy][x+dx] == Asteriod {
		return 1
	}

	return detectAsteroid(af, x+dx, y+dy, dx, dy)
}

func CountAsteroids(af AsteroidField) [][]int {
	res := make([][]int, len(af))
	for y := 0; y < len(af); y++ {
		res[y] = make([]int, len(af[0]))
		for x := 0; x < len(af[0]); x++ {
			if af[y][x] == Asteriod {
				res[y][x] += detectAsteroid(af, x, y, 0, -1)
				res[y][x] += detectAsteroid(af, x, y, 0, 1)
				res[y][x] += detectAsteroid(af, x, y, -1, 0)
				res[y][x] += detectAsteroid(af, x, y, 1, 0)
				res[y][x] += detectAsteroid(af, x, y, 1, 1)
				res[y][x] += detectAsteroid(af, x, y, 1, -1)
				res[y][x] += detectAsteroid(af, x, y, -1, 1)
				res[y][x] += detectAsteroid(af, x, y, -1, -1)

				for dy := 1; dy < len(af); dy++ {
					for dx := 1; dx < len(af[0]); dx++ {
						if dx == dy {
							continue
						}
						res[y][x] += detectAsteroid(af, x, y, -dx, -dy)
						res[y][x] += detectAsteroid(af, x, y, -dx, dy)
						res[y][x] += detectAsteroid(af, x, y, dx, -dy)
						res[y][x] += detectAsteroid(af, x, y, dx, dy)
					}
				}
			}
		}
	}
	fmt.Printf("%#v", res)
	return res
}

func GetAnswers() (int, int) {
	return 0, 0
}
