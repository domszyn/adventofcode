package day21

import (
	"math"
	"strconv"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

func add(to chan int, from1, from2 chan int) {
	to <- <-from1 + <-from2
}

func sub(to chan int, from1, from2 chan int) {
	to <- <-from1 - <-from2
}

func mul(to chan int, from1, from2 chan int) {
	to <- <-from1 * <-from2
}

func div(to chan int, from1, from2 chan int) {
	a := <-from1
	b := <-from2
	if b == 0 {
		to <- 0
	} else {
		to <- a / b
	}
}

type Tuple struct{ Left, Right int }

func cmp(to chan Tuple, from1, from2 chan int) {
	to <- Tuple{<-from1, <-from2}
}

func solve(lines []string, human []int) (res int, cmpResults []Tuple) {
	nRounds := len(human)
	if nRounds < 1 {
		nRounds = 1
	}

	monkeys := make(map[string]chan int)

	for _, l := range lines {
		m := l[:4]
		if n, err := strconv.Atoi(l[6:]); err == nil {
			monkeys[m] = make(chan int, nRounds)
			for i := 0; i < nRounds; i++ {
				if m == "humn" && nRounds > 1 {
					monkeys[m] <- human[i]
				} else {
					monkeys[m] <- n
				}
			}
			close(monkeys[m])
		} else {
			monkeys[m] = make(chan int, nRounds)
		}
	}

	cmpChan := make(chan Tuple, nRounds)

	for i := 0; i < nRounds; i++ {
		for _, l := range lines {
			val := l[6:]
			if _, err := strconv.Atoi(val); err != nil {
				m := l[:4]
				left := val[:4]
				right := val[7:]

				if m == "root" && nRounds > 1 {
					go cmp(cmpChan, monkeys[left], monkeys[right])
				} else {
					switch val[5] {
					case '+':
						go add(monkeys[m], monkeys[left], monkeys[right])
					case '-':
						go sub(monkeys[m], monkeys[left], monkeys[right])
					case '*':
						go mul(monkeys[m], monkeys[left], monkeys[right])
					case '/':
						go div(monkeys[m], monkeys[left], monkeys[right])
					}
				}
			}
		}

		if nRounds > 1 {
			cmpResults = append(cmpResults, <-cmpChan)
		} else {
			res = <-monkeys["root"]
		}
	}

	return
}

func SolvePart1(lines []string) int {
	res, _ := solve(lines, []int{})
	return res
}

func SolvePart2(lines []string) int {
	humanInputs := make([]int, 19)
	humanInputs[0] = 1
	for i := 1; i < len(humanInputs); i++ {
		humanInputs[i] = humanInputs[i-1] * 10
	}

	base := 0
	for round := 0; ; round++ {
		_, res := solve(lines, humanInputs)
		for i, v := range res {
			if v.Left > v.Right {
				base = humanInputs[i]
			}

			if v.Left == v.Right {
				return humanInputs[i]
			}
		}

		humanInputs = []int{}
		mag := int(math.Floor(math.Log10(float64(base)))) - round
		diff := int(math.Pow10(mag))
		for i := 1; i < 10; i++ {
			humanInputs = append(humanInputs, base+(i*diff))
		}
	}
}

func Solve() (int, int) {
	lines := utils.ReadInput("./solutions/day21/input.txt", mappers.ToString)

	return SolvePart1(lines), SolvePart2(lines)
}
