package day22

import (
	"bufio"
	"fmt"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type Cube struct {
	MinX, MinY, MinZ int
	MaxX, MaxY, MaxZ int
}

func (c Cube) Overlaps(c2 Cube) bool {
	return !(c.MaxX < c2.MinX || c.MinX > c2.MaxX ||
		c.MaxY < c2.MinY || c.MinY > c2.MaxY ||
		c.MaxZ < c2.MinZ || c.MinZ > c2.MaxZ)
}

func (c Cube) GetOverlap(c2 Cube) Cube {
	return Cube{
		max(c.MinX, c2.MinX), max(c.MinY, c2.MinY), max(c.MinZ, c2.MinZ),
		min(c.MaxX, c2.MaxX), min(c.MaxY, c2.MaxY), min(c.MaxZ, c2.MaxZ),
	}
}

func (c Cube) GetCount() int {
	return (c.MaxX - c.MinX + 1) * (c.MaxY - c.MinY + 1) * (c.MaxZ - c.MinZ + 1)
}

type CommandCube struct {
	Cube Cube
	On   bool
}

func (cc CommandCube) Overlaps(cc2 CommandCube) bool {
	return cc.Cube.Overlaps(cc2.Cube)
}

func (cc CommandCube) GetOverlap(cc2 CommandCube) Cube {
	return cc.Cube.GetOverlap(cc2.Cube)
}

func (cc CommandCube) Count() int {
	return cc.Cube.GetCount()
}

func getCommands(input string) (commands []CommandCube) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s[strings.Index(s, " ")+1:], ",")
		var minX, maxX, minY, maxY, minZ, maxZ int
		fmt.Sscanf(parts[0], "x=%d..%d", &minX, &maxX)
		fmt.Sscanf(parts[1], "y=%d..%d", &minY, &maxY)
		fmt.Sscanf(parts[2], "z=%d..%d", &minZ, &maxZ)

		var overlaps []CommandCube
		command := CommandCube{
			On:   s[1] == 'n',
			Cube: Cube{minX, minY, minZ, maxX, maxY, maxZ},
		}

		for _, prevCommand := range commands {
			if command.Overlaps(prevCommand) {
				overlap := command.GetOverlap(prevCommand)
				overlaps = append(overlaps, CommandCube{overlap, !prevCommand.On})
			}
		}

		commands = append(commands, overlaps...)
		if command.On {
			commands = append(commands, command)
		}
	}

	return
}

func Solve() (part1, part2 int) {
	for _, command := range getCommands(Input) {
		if command.On {
			part1 += command.Count()
		} else {
			part1 -= command.Count()
		}
	}

	for _, command := range getCommands(FullInput) {
		if command.On {
			part2 += command.Count()
		} else {
			part2 -= command.Count()
		}
	}

	return
}
