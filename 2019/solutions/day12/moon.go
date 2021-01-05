package day12

import "math"

type Moon struct {
	X, Y, Z, Vx, Vy, Vz int64
}

func (m *Moon) potentialEnergy() int64 {
	return int64(math.Abs(float64(m.X)) + math.Abs(float64(m.Y)) + math.Abs(float64(m.Z)))
}

func (m *Moon) kineticEnergy() int64 {
	return int64(math.Abs(float64(m.Vx)) + math.Abs(float64(m.Vy)) + math.Abs(float64(m.Vz)))
}

func (m *Moon) TotalEnergy() int64 {
	return m.potentialEnergy() * m.kineticEnergy()
}

func (m *Moon) Move() {
	m.X += m.Vx
	m.Y += m.Vy
	m.Z += m.Vz
}

func SigNum(x int64) int64 {
	return (x >> 63) | int64(uint64(-x)>>63)
}

func diff(a, b Moon) (dx, dy, dz int64) {
	dx = SigNum(a.X - b.X)
	dy = SigNum(a.Y - b.Y)
	dz = SigNum(a.Z - b.Z)
	return
}

func (m *Moon) Accelerate(m2 *Moon) {
	dx, dy, dz := diff(*m2, *m)
	m.Vx += dx
	m2.Vx -= dx
	m.Vy += dy
	m2.Vy -= dy
	m.Vz += dz
	m2.Vz -= dz
}

type MoonSlice []Moon

func (moons MoonSlice) Move() {
	for i := 0; i < len(moons); i++ {
		for j := i + 1; j < len(moons); j++ {
			moons[i].Accelerate(&moons[j])
		}

		moons[i].Move()
	}
}

func (moons MoonSlice) TotalEnergy() (totalEnergy int64) {
	for i := 0; i < len(moons); i++ {
		totalEnergy += moons[i].TotalEnergy()
	}
	return
}
