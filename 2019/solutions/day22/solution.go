package day22

import (
	"math/big"
	"strconv"
	"strings"
)

type Transformation struct{ A, B, Size *big.Int }

func stack(size *big.Int) Transformation {
	return Transformation{
		A:    new(big.Int).Sub(size, big.NewInt(1)),
		B:    new(big.Int).Sub(size, big.NewInt(1)),
		Size: new(big.Int).Set(size),
	}
}

func cut(size, n *big.Int) Transformation {
	b := new(big.Int).Neg(n)
	b.Mod(b, size)

	return Transformation{
		A:    big.NewInt(1),
		B:    b,
		Size: new(big.Int).Set(size),
	}
}

func deal(size, n *big.Int) Transformation {
	return Transformation{
		A:    new(big.Int).Set(n),
		B:    big.NewInt(0),
		Size: new(big.Int).Set(size),
	}
}

func (t2 Transformation) apply(t1 Transformation) Transformation {
	a := new(big.Int).Mul(t2.A, t1.A)
	a.Mod(a, t2.Size)

	b := new(big.Int).Mul(t1.A, t2.B)
	b.Add(b, t1.B)
	b.Mod(b, t2.Size)

	return Transformation{
		A:    a,
		B:    b,
		Size: t2.Size,
	}
}

func parseTransformation(line string, size *big.Int) Transformation {
	words := strings.Split(line, " ")
	if words[0] == "cut" {
		n, _ := strconv.Atoi(words[1])
		return cut(size, big.NewInt(int64(n)))
	}

	if words[0] == "deal" && words[3] == "stack" {
		return stack(size)
	}

	n, _ := strconv.Atoi(words[len(words)-1])
	return deal(size, big.NewInt(int64(n)))
}

func shuffle2(size *big.Int) Transformation {
	var transformation Transformation
	lines := strings.Split(Input, "\n")
	for i, l := range lines {
		if i == 0 {
			transformation = parseTransformation(l, size)
			continue
		}

		transformation = transformation.apply(parseTransformation(l, size))
	}

	return transformation
}

func solvePart1() int {
	size := big.NewInt(10_007)
	t := shuffle2(size)
	answer := new(big.Int).Set(t.A)
	answer.Mul(answer, big.NewInt(2019))
	answer.Add(answer, t.B)
	answer.Mod(answer, big.NewInt(10_007))

	return int(answer.Int64())
}

func solvePart2() int {
	size := big.NewInt(119_315_717_514_047)
	shuffles := big.NewInt(101_741_582_076_661)
	finalPos := big.NewInt(2020)

	t := shuffle2(size)
	a, b := t.A, t.B

	a_final := new(big.Int).Exp(a, shuffles, size)
	a_minus_1 := new(big.Int).Sub(a, big.NewInt(1))
	inv_a_minus_1 := new(big.Int).ModInverse(a_minus_1, size)

	term1 := new(big.Int).Sub(a_final, big.NewInt(1))
	b_final := new(big.Int).Mul(b, term1)
	b_final.Mul(b_final, inv_a_minus_1)
	b_final.Mod(b_final, size)

	term2 := new(big.Int).Sub(finalPos, b_final)
	inv_a_final := new(big.Int).ModInverse(a_final, size)

	initialPos := new(big.Int).Mul(term2, inv_a_final)
	initialPos.Mod(initialPos, size)

	return int(initialPos.Int64())
}

func GetAnswers() (int, int) {
	return solvePart1(), solvePart2()
}
