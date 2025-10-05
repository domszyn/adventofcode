package day20

import (
	"math"
	"strings"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

type Connection struct {
	Label       string
	LevelChange int8
	Location    toolbox.Point
}

type GridMap = map[toolbox.Point]bool
type DistanceMap = map[Connection]map[Connection]int

func NewDistanceMap(connections []Connection) DistanceMap {
	distances := make(DistanceMap)

	for _, conn := range connections {
		distances[conn] = make(map[Connection]int)
	}

	return distances
}

type SpaceMap struct {
	Grid        GridMap
	Connections []Connection
	Distances   DistanceMap
	Start, End  Connection
	Portals     map[Connection]Connection
}

func isLetter(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func getLevelChange(label string, min, max, point toolbox.Point) int8 {
	if label == "AA" || label == "ZZ" {
		return 0
	}

	if point.X == min.X || point.X == max.X || point.Y == min.Y || point.Y == max.Y {
		return -1
	}

	return 1
}

func ParseMap(allowWalkBetweenPortals bool) SpaceMap {
	var spaceMap SpaceMap
	lines := strings.Split(Input, "\n")
	grid := make(map[toolbox.Point]rune)

	min := toolbox.Point{X: math.MaxInt, Y: math.MaxInt}
	var max toolbox.Point

	for y, l := range lines {
		for x, c := range l {
			if c == ' ' {
				continue
			}

			if c == '.' {
				if x < min.X {
					min.X = x
				}
				if x > max.X {
					max.X = x
				}
				if y < min.Y {
					min.Y = y
				}
				if y > max.Y {
					max.Y = y
				}
			}

			grid[toolbox.Point{X: x, Y: y}] = c
		}
	}

	for k, v := range grid {
		if v == '.' || v == '#' {
			continue
		}

		var label string
		var point toolbox.Point

		if isLetter(v) && isLetter(grid[toolbox.Point{X: k.X, Y: k.Y - 1}]) && grid[toolbox.Point{X: k.X, Y: k.Y - 2}] == '.' {
			// up
			label = string([]rune{grid[toolbox.Point{X: k.X, Y: k.Y - 1}], v})
			point = toolbox.Point{X: k.X, Y: k.Y - 2}
		} else if isLetter(v) && isLetter(grid[toolbox.Point{X: k.X, Y: k.Y + 1}]) && grid[toolbox.Point{X: k.X, Y: k.Y + 2}] == '.' {
			// down
			label = string([]rune{v, grid[toolbox.Point{X: k.X, Y: k.Y + 1}]})
			point = toolbox.Point{X: k.X, Y: k.Y + 2}
		} else if isLetter(v) && isLetter(grid[toolbox.Point{X: k.X - 1, Y: k.Y}]) && grid[toolbox.Point{X: k.X - 2, Y: k.Y}] == '.' {
			// left
			label = string([]rune{grid[toolbox.Point{X: k.X - 1, Y: k.Y}], v})
			point = toolbox.Point{X: k.X - 2, Y: k.Y}
		} else if isLetter(v) && isLetter(grid[toolbox.Point{X: k.X + 1, Y: k.Y}]) && grid[toolbox.Point{X: k.X + 2, Y: k.Y}] == '.' {
			// right
			label = string([]rune{v, grid[toolbox.Point{X: k.X + 1, Y: k.Y}]})
			point = toolbox.Point{X: k.X + 2, Y: k.Y}
		} else {
			continue
		}

		levelChange := getLevelChange(label, min, max, point)
		connection := Connection{
			Label:       label,
			LevelChange: levelChange,
			Location:    point,
		}
		spaceMap.Connections = append(spaceMap.Connections, connection)

		switch label {
		case "AA":
			spaceMap.Start = connection
		case "ZZ":
			spaceMap.End = connection
		}
	}

	spaceMap.Grid = make(GridMap)

	for k, v := range grid {
		if v == '.' {
			spaceMap.Grid[k] = true
		}
	}

	spaceMap.Distances = NewDistanceMap(spaceMap.Connections)

	locationToConnection := make(map[toolbox.Point]Connection)
	for _, c := range spaceMap.Connections {
		locationToConnection[c.Location] = c
	}

	for _, from := range spaceMap.Connections {
		queue := toolbox.NewQueue[toolbox.Point]()
		cost := make(map[toolbox.Point]int)

		startLoc := from.Location
		queue.Enqueue(startLoc)
		cost[startLoc] = 0

		for queue.Size() > 0 {
			currentLoc, _ := queue.Dequeue()
			currentCost := cost[currentLoc]

			if destConn, isPortal := locationToConnection[currentLoc]; isPortal && currentLoc != startLoc {
				if dist, found := spaceMap.Distances[from][destConn]; !found || currentCost < dist {
					spaceMap.Distances[from][destConn] = currentCost
				}
			}

			for _, n := range currentLoc.GetNeighbors() {
				if _, visited := cost[n]; visited || !spaceMap.Grid[n] {
					continue
				}

				cost[n] = currentCost + 1
				queue.Enqueue(n)
			}
		}
	}

	spaceMap.Portals = make(map[Connection]Connection)

	for _, c1 := range spaceMap.Connections {
		for _, c2 := range spaceMap.Connections {
			if c1.Label == c2.Label && c1.Location != c2.Location {
				if allowWalkBetweenPortals {
					spaceMap.Distances[c1][c2] = 1
				}
				spaceMap.Portals[c1] = c2
			}
		}
	}

	return spaceMap
}
