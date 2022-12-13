package day13

import (
	"github.com/domszyn/adventofcode/2022/utils"
)

const (
	RightOrder     = 1
	IncorrectOrder = -1
	Undecided      = 0
)

type Packet interface {
	Compare(Packet) int
}

type IntegerPacket struct {
	Value int
}

type CompoundPacket struct {
	Values []Packet
}

func (p IntegerPacket) Compare(other Packet) int {
	if integer, ok := other.(IntegerPacket); ok {
		if p.Value < integer.Value {
			return RightOrder
		}

		if p.Value > integer.Value {
			return IncorrectOrder
		}

		return Undecided
	}

	return CompoundPacket{[]Packet{p}}.Compare(other)
}

func (p CompoundPacket) Compare(other Packet) int {
	if _, ok := other.(IntegerPacket); ok {
		return p.Compare(CompoundPacket{[]Packet{other}})
	}

	op := other.(CompoundPacket)

	maxLen := utils.Max([]int{len(p.Values), len(op.Values)})
	for i := 0; i < maxLen; i++ {
		if i >= len(p.Values) {
			return RightOrder
		}

		if i >= len(op.Values) {
			return IncorrectOrder
		}

		if rightOrder := p.Values[i].Compare(op.Values[i]); rightOrder != Undecided {
			return rightOrder
		}
	}

	return Undecided
}
