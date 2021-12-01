package day2

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func parseDimensions(dimensions string) (length, width, height int) {
	parts := strings.SplitN(dimensions, "x", 3)
	length, _ = strconv.Atoi(parts[0])
	width, _ = strconv.Atoi(parts[1])
	height, _ = strconv.Atoi(parts[2])
	return
}

func CalculatePaperAndRibbon() (paper, ribbon int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		length, width, height := parseDimensions(scanner.Text())
		sides := []int{length * width, width * height, height * length}
		perimeters := []int{length + width, width + height, height + length}
		minSide := math.MaxInt16
		minPerimeter := math.MaxInt16
		for i := 0; i < 3; i++ {
			side := sides[i]
			perimeter := perimeters[i]
			paper += 2 * side
			if side < minSide {
				minSide = side
			}
			if perimeter < minPerimeter {
				minPerimeter = perimeter
			}
		}
		paper += minSide
		ribbon += 2*minPerimeter + length*width*height
	}

	return
}
