package day12

import (
	"bufio"
	"regexp"
	"strings"
)

const (
	SMALL = iota
	BIG
)

type Cave struct {
	Name      string
	Visited   int
	Type      int
	Connected []*Cave
}

func isBigCave(c string) bool {
	matched, _ := regexp.MatchString("^[A-Z]+$", c)
	return matched
}

func smallCaveVisitedTwice(caves map[string]*Cave) bool {
	for _, cave := range caves {
		if cave.Type == SMALL && cave.Visited == 2 {
			return true
		}
	}

	return false
}

func findPaths(caves map[string]*Cave, name string, maxVisits int) (count int) {
	for _, cc := range caves[name].Connected {
		if len(cc.Connected) == 0 {
			count++
			continue
		}

		if smallCaveVisitedTwice(caves) {
			maxVisits = 1
		}

		if cc.Type == BIG || (cc.Type == SMALL && cc.Visited < maxVisits) {
			cc.Visited++
			count += findPaths(caves, cc.Name, maxVisits)
			cc.Visited--
		}
	}

	return
}

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	caves := make(map[string]*Cave)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "-")
		from, to := s[0], s[1]
		if from == "end" || to == "start" {
			from, to = to, from
		}
		if _, found := caves[from]; !found {
			cave := Cave{
				Name: from,
			}

			if isBigCave(from) {
				cave.Type = BIG
			}

			caves[from] = &cave
		}

		if _, found := caves[to]; !found {
			cave := Cave{
				Name: to,
			}

			if isBigCave(to) {
				cave.Type = BIG
			}

			caves[to] = &cave
		}

		caves[from].Connected = append(caves[from].Connected, caves[to])
		if from != "start" && to != "end" {
			caves[to].Connected = append(caves[to].Connected, caves[from])
		}
	}

	return findPaths(caves, "start", 1), findPaths(caves, "start", 2)
}
