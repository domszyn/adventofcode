package day6

import (
	"fmt"
	"strings"
)

type SpaceObject struct {
	Name       string
	Level      int
	Parent     *SpaceObject
	Satellites []*SpaceObject
}

func readInput() map[string]*SpaceObject {
	allOrbits := make(map[string]*SpaceObject)

	for _, s := range strings.Split(Input, "\n") {
		var parentName, satelliteName string
		fmt.Sscanf(s, "%3s)%3s", &parentName, &satelliteName)

		var parent, satellite *SpaceObject
		var found bool

		if satellite, found = allOrbits[satelliteName]; !found {
			satellite = &SpaceObject{Name: satelliteName, Satellites: []*SpaceObject{}}
		}

		if parent, found = allOrbits[parentName]; !found {
			parent = &SpaceObject{Name: parentName, Satellites: []*SpaceObject{}}
		}

		parent.Satellites = append(parent.Satellites, satellite)
		satellite.Parent = parent

		allOrbits[satelliteName] = satellite
		allOrbits[parentName] = parent
	}

	return allOrbits
}

func countOrbits(allOrbits map[string]*SpaceObject, name string, level int) int {
	orbits := level

	spaceObject := allOrbits[name]
	spaceObject.Level = level

	if spaceObject != nil {
		for _, satellite := range spaceObject.Satellites {
			orbits += countOrbits(allOrbits, satellite.Name, level+1)
		}
	}

	return orbits
}

func GetAnswers() (int, int) {
	allOrbits := readInput()
	count := countOrbits(allOrbits, "COM", 0)

	you := allOrbits["YOU"]
	san := allOrbits["SAN"]
	yourLevel := you.Level
	sanLevel := san.Level

	for san.Level > you.Level {
		san = san.Parent
	}

	for you.Level > san.Level {
		you = you.Parent
	}

	for you.Parent != san.Parent {
		you = you.Parent
		san = san.Parent
	}

	return count, yourLevel + sanLevel - 2*you.Level
}
