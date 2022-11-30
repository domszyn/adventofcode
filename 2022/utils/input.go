package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadInput[T any](mapFn func(string) T) (result []T) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		result = append(result, mapFn(s))
	}

	return
}
