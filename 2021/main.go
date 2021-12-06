package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/domszyn/adventofcode/2021/solutions"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Advent of Code 2021")

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
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Type a number between 1 and 25 to get the answer and quit to exit")
		}
	}
}
