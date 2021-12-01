package day7

import (
	"bufio"
	"strconv"
	"strings"
)

type Circuit map[string]Wire

type Wire struct {
	Label string
	Input string
	Value *uint16
}

func parseWire(s string) Wire {
	parts := strings.Split(s, " -> ")
	wire := Wire{
		Label: parts[1],
		Input: parts[0],
	}
	if val, err := strconv.ParseUint(wire.Input, 10, 16); err == nil {
		value := uint16(val)
		wire.Value = &value
	}
	return wire
}

func (c Circuit) TryGetWireValue(label string) {
	wire := c[label]
	if wire.Value != nil {
		return
	}

	tokens := strings.Split(wire.Input, " ")
	if len(tokens) == 1 && c[tokens[0]].Value != nil {
		c[label] = Wire{
			Label: wire.Label,
			Input: wire.Input,
			Value: c[tokens[0]].Value,
		}
	} else if len(tokens) == 2 && tokens[0] == "NOT" && c[tokens[1]].Value != nil {
		val := ^(*c[tokens[1]].Value)
		c[label] = Wire{
			Label: wire.Label,
			Input: wire.Input,
			Value: &val,
		}
	} else if len(tokens) == 3 {
		var left, right, val uint16
		op := tokens[1]

		if val, err := strconv.Atoi(tokens[0]); err == nil {
			left = uint16(val)
		} else if c[tokens[0]].Value != nil {
			left = *c[tokens[0]].Value
		} else {
			return
		}

		if val, err := strconv.Atoi(tokens[2]); err == nil {
			right = uint16(val)
		} else if c[tokens[2]].Value != nil {
			right = *c[tokens[2]].Value
		} else {
			return
		}

		switch op {
		case "LSHIFT":
			val = left << right
		case "RSHIFT":
			val = left >> right
		case "AND":
			val = left & right
		case "OR":
			val = left | right
		}

		c[label] = Wire{
			Label: wire.Label,
			Input: wire.Input,
			Value: &val,
		}
	}
}

func Solve(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	circuit := make(Circuit)

	for scanner.Scan() {
		s := scanner.Text()
		wire := parseWire(s)
		circuit[wire.Label] = wire
	}

	for circuit["a"].Value == nil {
		for label := range circuit {
			circuit.TryGetWireValue(label)
		}

		var resolved int
		for _, val := range circuit {
			if val.Value != nil {
				resolved++
			}
		}
	}

	return int(*circuit["a"].Value)
}
