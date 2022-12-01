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

func Max(numbers []int) (max int) {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}

	return
}

func Sum(numbers []int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return
}
