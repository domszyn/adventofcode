package day7

import (
	"bufio"
	"strings"
)

type AddressType string

const (
	Hypernet = AddressType("Hypernet")
	Supernet = AddressType("Supernet")
)

type AddressPart struct {
	Address string
	Type    AddressType
}

func parseAddress(address string) (parts []AddressPart) {
	for {
		if i := strings.Index(address, "["); i >= 0 {
			parts = append(parts, AddressPart{Address: address[:i], Type: Supernet})
			closeBracketIdx := strings.Index(address, "]")
			parts = append(parts, AddressPart{Address: address[i+1 : closeBracketIdx], Type: Hypernet})
			address = address[closeBracketIdx+1:]
		} else {
			parts = append(parts, AddressPart{Address: address, Type: Supernet})
			return
		}
	}
}

func containsABBA(seq string) bool {
	if len(seq) < 4 {
		return false
	}

	for i := 0; i < len(seq)-3; i++ {
		if seq[i] == seq[i+3] &&
			seq[i+1] == seq[i+2] &&
			seq[i] != seq[i+1] {
			return true
		}
	}

	return false
}

func supportsTLS(address string) (res bool) {
	parts := parseAddress(address)
	for _, p := range parts {
		if p.Type == Supernet && containsABBA(p.Address) {
			res = true
			break
		}
	}

	for _, p := range parts {
		if p.Type == Hypernet && containsABBA(p.Address) {
			res = false
			break
		}
	}

	return
}

func findABA(parts []AddressPart) (abas []string) {
	for _, v := range parts {
		if v.Type == Hypernet {
			continue
		}

		seq := v.Address
		for i := 0; i < len(seq)-2; i++ {
			if seq[i] == seq[i+2] && seq[i] != seq[i+1] {
				abas = append(abas, seq[i:i+3])
			}
		}
	}

	return
}

func supportsSSL(address string) bool {
	parts := parseAddress(address)
	abas := findABA(parts)

	for _, v := range abas {
		for _, x := range parts {
			if x.Type == Hypernet && strings.Index(x.Address, string([]byte{v[1], v[0], v[1]})) >= 0 {
				return true
			}
		}
	}

	return false
}

func Solve() (int, int) {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	scanner.Split(bufio.ScanLines)
	supportTLS := 0
	supportSSL := 0
	for scanner.Scan() {
		address := scanner.Text()
		if supportsTLS(address) {
			supportTLS++
		}
		if supportsSSL(address) {
			supportSSL++
		}
	}

	return supportTLS, supportSSL
}
