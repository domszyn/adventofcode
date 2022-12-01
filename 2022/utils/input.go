package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadInput[T any](file string, mapFn func(string) T) (result []T) {
	f, err := os.Open(file)
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
