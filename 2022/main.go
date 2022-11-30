package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/domszyn/adventofcode/2022/solutions"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Advent of Code 2022")

	solution := solutions.Solution{}
	for {
		fmt.Print("<*> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		switch text {
		case "1":
			solution.SolveDay01()
		case "2":
			solution.SolveDay02()
		case "3":
			solution.SolveDay03()
		case "4":
			solution.SolveDay04()
		case "5":
			solution.SolveDay05()
		case "6":
			solution.SolveDay06()
		case "7":
			solution.SolveDay07()
		case "8":
			solution.SolveDay08()
		case "9":
			solution.SolveDay09()
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
		case "20":
			solution.SolveDay20()
		case "21":
			solution.SolveDay21()
		case "22":
			solution.SolveDay22()
		case "23":
			solution.SolveDay23()
		case "24":
			solution.SolveDay24()
		case "25":
			solution.SolveDay25()
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Type a number between 1 and 25 to get the answer and quit to exit")
		}
	}
}
