package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/domszyn/adventofcode/2019/solutions"
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
			break
		case "2":
			solution.SolveDay2()
			break
		case "3":
			solution.SolveDay3()
			break
		case "4":
			solution.SolveDay4()
			break
		case "5":
			solution.SolveDay5()
			break
		case "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25":
			fmt.Println("Solution is not ready yet")
			break
		case "quit":
			os.Exit(0)
			break
		default:
			fmt.Println("Type a number between 1 and 25 to get the answer and quit to exit")
			break
		}

	}
}
