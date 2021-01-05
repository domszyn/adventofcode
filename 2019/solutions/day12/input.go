package day12

import (
	"fmt"
	"strings"
)

const Input = `<x=3, y=-6, z=6>
<x=10, y=7, z=-9>
<x=-3, y=-7, z=9>
<x=-8, y=0, z=4>`

func calculateMoonPositions() MoonSlice {
	var moons []Moon

	for _, p := range strings.Split(Input, "\n") {
		moon := Moon{}
		fmt.Sscanf(p, "<x=%d, y=%d, z=%d>", &moon.X, &moon.Y, &moon.Z)
		moons = append(moons, moon)
	}

	return moons
}
