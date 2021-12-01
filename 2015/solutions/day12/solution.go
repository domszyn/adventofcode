package day12

import (
	"encoding/json"
	"io"
	"strings"
)

func SolvePart1() float64 {
	decoder := json.NewDecoder(strings.NewReader(Input))

	var sum float64
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}

		if number, ok := t.(float64); ok {
			sum += number
		}
	}

	return sum
}

func SolvePart2() float64 {
	var data []interface{}
	// decoder := json.NewDecoder(strings.NewReader(Input))

	err := json.Unmarshal([]byte(Input), &data)
	if err != nil {
		return -1
	}

	return processArray(data)
}

func processArray(data []interface{}) (sum float64) {
	for _, dataElement := range data {
		sum += processElement(dataElement)
	}

	return
}

func processElement(element interface{}) float64 {
	if object, ok := element.(map[string]interface{}); ok {
		return processObject(object)
	}
	if array, ok := element.([]interface{}); ok {
		return processArray(array)
	}
	if number, ok := element.(float64); ok {
		return number
	}
	return 0
}

func processObject(data map[string]interface{}) (sum float64) {
	for _, value := range data {
		if value == "red" {
			return 0
		}
	}

	for _, value := range data {
		sum += processElement(value)
	}

	return
}
