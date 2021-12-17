package day17

type Location struct{ X, Y int }

func (l *Location) WithinTarget() bool {
	return l.X >= 209 && l.X <= 238 && l.Y >= -86 && l.Y <= -59
}

func (l *Location) DeeperThanTarget() bool {
	return l.Y < -86
}

func (l *Location) NotCloseEnough() bool {
	return l.X < 209
}

func Solve() (part1, part2 int) {
	// target area: x=209..238, y=-86..-59
	probes := make(map[Location]int)

	velX := 0
	for {
		if velX*(velX+1)/2 > 209 {
			break
		}
		velX++
	}

	for y := 85; y >= -86; y-- {
		for x := velX; x <= 238; x++ {
			vel := Location{X: x, Y: y}
			loc := Location{X: 0, Y: 0}
			initialVel := vel
			probes[initialVel] = 0

			for {
				loc.X += vel.X
				loc.Y += vel.Y

				if loc.Y > probes[initialVel] {
					probes[initialVel] = loc.Y
				}

				if vel.X > 0 {
					vel.X--
				}

				vel.Y--

				if loc.WithinTarget() || loc.DeeperThanTarget() || (loc.NotCloseEnough() && vel.X == 0) {
					break
				}
			}

			if !loc.WithinTarget() {
				probes[initialVel] = -1
			}
		}
	}

	maxY := 0

	for _, v := range probes {
		if v > maxY {
			maxY = v
		}

		if v != -1 {
			part2++
		}
	}

	part1 = maxY

	return
}
