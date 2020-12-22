package solutions

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func calcModuleFuel(mass int64) int64 {
	return int64(math.Floor(float64(mass)/3)) - 2
}

func calcAdditionalFuel(mass int64) int64 {
	fuel := calcModuleFuel(mass)
	if fuel <= 0 {
		return 0
	}

	return fuel + calcAdditionalFuel(fuel)
}

func fuelRequirement(accountForFuelMass bool) int64 {
	file, err := os.Open("./input/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var totalFuel int64 = 0

	for scanner.Scan() {
		mass, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		if accountForFuelMass {
			totalFuel += calcAdditionalFuel(mass)
		} else {
			totalFuel += calcModuleFuel(mass)
		}
	}

	return totalFuel
}

func (s *Solution) SolveDay1() {
	start := time.Now()
	s.Part1 = fuelRequirement(false)
	s.Part2 = fuelRequirement(true)
	s.ExecTime = time.Since(start)
}
