package day1

import (
	"bufio"
	"math"
	"strconv"
	"strings"
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

func FuelRequirement(accountForFuelMass bool) int64 {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)
	var totalFuel int64

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
