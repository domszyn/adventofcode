package day16

import (
	"testing"
)

func TestFFT(t *testing.T) {
	if output := FFT("12345678", 4, 0); output != "01029498" {
		t.Errorf("Expected 01029498, got %s", output)
	}

	// if output := FFT("80871224585914546619083218645595", 100, 0); output != "24176176" {
	// 	t.Errorf("Expected 24176176, got %s", output)
	// }

	// if output := FFT("19617804207202209144916044189917", 100, 0); output != "73745418" {
	// 	t.Errorf("Expected 73745418, got %s", output)
	// }

	// if output := FFT("69317163492948606335995924319873", 100, 0); output != "52432133" {
	// 	t.Errorf("Expected 52432133, got %s", output)
	// }

	// if output := SolvePart1(); output != "96136976" {
	// 	t.Errorf("Expected 96136976 , got %s", output)
	// }

	// if output := SolvePart2("03036732577212944063491565474664"); output != "84462026" {
	// 	t.Errorf("Expected 84462026 , got %s", output)
	// }
}
