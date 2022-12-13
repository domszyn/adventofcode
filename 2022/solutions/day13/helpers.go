package day13

import (
	"strconv"
)

func parsePacket(s string) Packet {
	if s == "[]" {
		return CompoundPacket{}
	}
	contents := s[1 : len(s)-1]

	var elements []Packet
	for len(contents) > 0 {
		switch contents[0] {
		case '[':
			opened := 1
			for i := 1; i < len(s); i++ {
				if contents[i] == '[' {
					opened++
				}
				if contents[i] == ']' {
					opened--
				}

				if opened == 0 {
					elements = append(elements, parsePacket(contents[:i+1]))
					contents = contents[i+1:]
					break
				}
			}
		case ',':
			contents = contents[1:]
		default:
			noComa := true

			for i := 1; i < len(contents); i++ {
				if contents[i] == ',' {
					x, _ := strconv.Atoi(contents[:i])
					elements = append(elements, IntegerPacket{x})
					contents = contents[i+1:]
					noComa = false
					break
				}
			}

			if noComa {
				x, _ := strconv.Atoi(contents)
				elements = append(elements, IntegerPacket{x})
				contents = ""
			}
		}
	}

	return CompoundPacket{elements}
}

func isDivider(p Packet, x int) bool {
	if cp, ok := p.(CompoundPacket); ok {
		if len(cp.Values) != 1 {
			return false
		}

		if cp2, ok2 := cp.Values[0].(CompoundPacket); ok2 {
			if len(cp2.Values) != 1 {
				return false
			}

			if ip, ok3 := cp2.Values[0].(IntegerPacket); ok3 {
				return ip.Value == x
			}
		}
	}
	return false
}
