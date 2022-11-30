package day20

import (
	"bufio"
	"strconv"
	"strings"
)

func getPixelIndex(image []string, x, y int, inifiniteLit bool) (int64, error) {
	indexStr := ""
	for yy := y - 1; yy < y+2; yy++ {
		if yy < 0 || yy >= len(image) {
			if inifiniteLit {
				indexStr += "111"
			} else {
				indexStr += "000"
			}
			continue
		}
		for xx := x - 1; xx < x+2; xx++ {
			if xx < 0 || xx >= len(image[0]) {
				if inifiniteLit {
					indexStr += "1"
				} else {
					indexStr += "0"
				}
				continue
			}
			if image[yy][xx] == '#' {
				indexStr += "1"
			} else {
				indexStr += "0"
			}
		}
	}

	return strconv.ParseInt(indexStr, 2, 64)
}

func Solve() (part1, part2 int) {
	scanner := bufio.NewScanner(strings.NewReader(InputImage))
	scanner.Split(bufio.ScanLines)

	var image []string
	for scanner.Scan() {
		s := scanner.Text()
		image = append(image, s)
	}

	infiniteLit := false
	for i := 0; i < 2; i++ {
		nextImage := make([]string, 0, len(image)+2)
		for y := 0; y < len(image)+2; y++ {
			nextImage = append(nextImage, "")
			for x := 0; x < len(image[0])+2; x++ {
				idx, _ := getPixelIndex(image, x-1, y-1, infiniteLit)
				nextImage[y] += string(ImageEnhancementAlgorithm[idx])
			}
		}
		image = nextImage
		infiniteLit = !infiniteLit
	}

	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image); x++ {
			if image[y][x] == '#' {
				part1++
			}
		}
	}

	for i := 0; i < 48; i++ {
		nextImage := make([]string, 0, len(image)+2)
		for y := 0; y < len(image)+2; y++ {
			nextImage = append(nextImage, "")
			for x := 0; x < len(image[0])+2; x++ {
				idx, _ := getPixelIndex(image, x-1, y-1, infiniteLit)
				nextImage[y] += string(ImageEnhancementAlgorithm[idx])
			}
		}
		image = nextImage
		infiniteLit = !infiniteLit
	}

	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image); x++ {
			if image[y][x] == '#' {
				part2++
			}
		}
	}

	return
}
