package day12

import "github.com/domszyn/adventofcode/2019/toolbox"

func SolvePart1() int64 {
	moons := calculateMoonPositions()

	for step := 0; step < 1000; step++ {
		moons.Move()
	}

	return moons.TotalEnergy()
}

func equal(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func findSteps(positions []int64) int64 {
	velocities := make([]int64, len(positions))
	endVelocities := make([]int64, len(positions))

	var steps int64 = 0

	for {
		velocityChange := velocityDiff(positions)
		for i := 0; i < len(positions); i++ {
			velocities[i] += velocityChange[i]
		}
		for i := 0; i < len(positions); i++ {
			positions[i] += velocities[i]
		}

		steps++
		if equal(velocities, endVelocities) {
			break
		}
	}

	return steps * 2
}

type ProjectFn func(Moon) int64

func project(moons MoonSlice, fn ProjectFn) []int64 {
	res := make([]int64, len(moons))
	for i := 0; i < len(moons); i++ {
		res[i] = fn(moons[i])
	}

	return res
}

func velocityDiff(positions []int64) []int64 {
	res := make([]int64, len(positions))
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			diff := SigNum(positions[j] - positions[i])
			res[i] += diff
			res[j] -= diff
		}
	}
	return res
}

func SolvePart2() int64 {
	moons := calculateMoonPositions()

	x := findSteps(project(moons, func(m Moon) int64 { return m.X }))
	y := findSteps(project(moons, func(m Moon) int64 { return m.Y }))
	z := findSteps(project(moons, func(m Moon) int64 { return m.Z }))

	return toolbox.LcmN(x, y, z)
}

func GetAnswers() (int64, int64) {
	return SolvePart1(), SolvePart2()
}
