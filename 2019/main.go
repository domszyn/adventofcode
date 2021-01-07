package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/domszyn/adventofcode/2019/solutions"
	"github.com/domszyn/adventofcode/2019/solutions/day13"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Advent of Code 2019")

	solution := solutions.Solution{}
	for {
		fmt.Print("<*> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		switch text {
		case "1":
			solution.SolveDay1()
		case "2":
			solution.SolveDay2()
		case "3":
			solution.SolveDay3()
		case "4":
			solution.SolveDay4()
		case "5":
			solution.SolveDay5()
		case "6":
			solution.SolveDay6()
		case "7":
			solution.SolveDay7()
		case "8":
			solution.SolveDay8()
		case "9":
			solution.SolveDay9()
		case "11":
			solution.SolveDay11()
		case "12":
			solution.SolveDay12()
		case "13":
			solution.SolveDay13()
		case "13 animate":
			day13.SolvePart2(true, false)
		case "13 fastwin":
			day13.SolvePart2(true, true)
		case "10", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25":
			fmt.Println("Solution is not ready yet")
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Type a number between 1 and 25 to get the answer and quit to exit")
		}

	}
}
