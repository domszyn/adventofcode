package mappers

import (
	"log"
	"strconv"
)

func ToString(s string) string {
	return s
}

func ToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
