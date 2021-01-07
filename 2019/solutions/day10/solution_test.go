package day10

import "testing"

func TestDetection(_ *testing.T) {
	var af = [][]rune{
		[]rune(".#..#"),
		[]rune("....."),
		[]rune("#####"),
		[]rune("....#"),
		[]rune("...##"),
	}
	CountAsteroids(af)
}
