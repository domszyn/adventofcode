package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/domszyn/adventofcode/2015/solutions"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Advent of Code 2015")

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
		case "10":
			solution.SolveDay10()
		case "11":
			solution.SolveDay11()
		case "12":
			solution.SolveDay12()
		case "13":
			solution.SolveDay13()
		case "14":
			solution.SolveDay14()
		case "15":
			solution.SolveDay15()
		case "16":
			solution.SolveDay16()
		case "17":
			solution.SolveDay17()
		case "18":
			solution.SolveDay18()
		case "19":
			solution.SolveDay19()
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Type a number between 1 and 25 to get the answer and quit to exit")
		}
	}
}
