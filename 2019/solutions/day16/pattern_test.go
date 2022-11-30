package day16

import (
	"fmt"
	"testing"
)

func TestCreatePattern(_ *testing.T) {
	pattern := CreatePattern(1, 10)
	fmt.Printf("%#v\n", pattern)
	pattern = CreatePattern(2, 10)
	fmt.Printf("%#v\n", pattern)
	pattern = CreatePattern(3, 10)
	fmt.Printf("%#v\n", pattern)
	pattern = CreatePattern(4, 10)
	fmt.Printf("%#v\n", pattern)
	pattern = CreatePattern(5, 10)
	fmt.Printf("%#v\n", pattern)
	pattern = CreatePattern(6, 10)
	fmt.Printf("%#v\n", pattern)
	pattern = CreatePattern(7, 10)
	fmt.Printf("%#v\n", pattern)
	pattern = CreatePattern(8, 10)
	fmt.Printf("%#v\n", pattern)
	pattern = CreatePattern(9, 10)
	fmt.Printf("%#v\n", pattern)
	pattern = CreatePattern(10, 10)
	fmt.Printf("%#v\n", pattern)
	pattern = CreatePattern(11, 10)
	fmt.Printf("%#v\n", pattern)
}
