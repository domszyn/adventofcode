package day4

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type Room struct {
	Name     string
	SectorID int
	Checksum string
}

type Letter struct {
	Name  rune
	Count int
}

type ByFrequency []Letter

func (f ByFrequency) Len() int      { return len(f) }
func (f ByFrequency) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
func (f ByFrequency) Less(i, j int) bool {
	if f[i].Count == f[j].Count {
		return f[i].Name < f[j].Name
	}

	return f[i].Count > f[j].Count
}

func (f ByFrequency) Checksum() string {
	checksum := ""
	for i := 0; i < 5; i++ {
		checksum += string(f[i].Name)
	}
	return checksum
}

func (r Room) IsReal() bool {
	lettersMap := make(map[rune]int)
	for _, r := range r.Name {
		if r == '-' {
			continue
		}
		lettersMap[r]++
	}

	letters := []Letter{}
	for k, v := range lettersMap {
		letters = append(letters, Letter{Name: k, Count: v})
	}

	sort.Sort(ByFrequency(letters))

	return r.Checksum == ByFrequency(letters).Checksum()
}

func shiftString(input string) (output string) {
	for _, r := range input {
		if r == '-' {
			output += "-"
			continue
		}

		nextLetter := r + 1
		if nextLetter > 'z' {
			nextLetter = 'a'
		}
		output += string(nextLetter)
	}
	return
}

func (r Room) DecryptName() (decryptedName string) {
	decryptedName = r.Name
	for i := 0; i < r.SectorID; i++ {
		decryptedName = shiftString(decryptedName)
	}
	return
}

func parseInput() (rooms []Room) {
	r := Room{}
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		r.Name = strings.Join(parts[:len(parts)-1], "-")
		fmt.Sscanf(parts[len(parts)-1], "%d%s", &r.SectorID, &r.Checksum)
		r.Checksum = r.Checksum[1 : len(r.Checksum)-1]
		rooms = append(rooms, r)
	}
	return
}

func Solve() (part1 int, part2 int) {
	for _, r := range parseInput() {
		if r.IsReal() {
			part1 += r.SectorID
			if r.DecryptName() == "northpole-object-storage" {
				part2 = r.SectorID
			}
		}
	}

	return
}
