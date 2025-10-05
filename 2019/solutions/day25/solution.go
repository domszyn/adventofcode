package day25

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

func solve() {
	reader := bufio.NewReader(os.Stdin)

	program := toolbox.LoadProgram(Input)
	input := make(chan int)
	output := make(chan int)
	done := make(chan bool, 1)

	go program.IntCode(input, output, done)

	inputIdx := 0
	var finalInput string
	var prompt []rune
	// commands := toolbox.NewQueue(
	// 	"north\n",
	// 	"north\n",
	// 	"east\n",
	// 	"west\n",
	// )

loop:
	for {
		if strings.HasSuffix(string(prompt), "Command?") {
			fmt.Println(string(prompt))
			// cmd, _ := commands.Dequeue()
			cmd, _ := reader.ReadString('\n')
			finalInput = cmd + "\n"
			prompt = []rune{}
			inputIdx = 0
		}
		var inputChan chan int
		var nextInput int

		if inputIdx < len(finalInput) {
			inputChan = input
			nextInput = int(finalInput[inputIdx])
		}

		select {
		case inputChan <- nextInput:
			inputIdx++
		case val := <-output:
			prompt = append(prompt, rune(val))
		case <-done:
			break loop
		}
	}
}

func GetAnswers() (int, int) {
	solve()
	return -1, -1
}
