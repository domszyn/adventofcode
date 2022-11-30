package day6

type MessageData map[string]int

func parseInput() (letters []MessageData) {
	i := 0
	for _, v := range Input {
		if v == '\n' {
			i = 0
			continue
		}
		if len(letters) < i+1 {
			letters = append(letters, make(MessageData))
		}

		letters[i][string(v)]++
		i++
	}
	return
}

func findMostCommonLetter(message MessageData) string {
	letter, frequency := "", 0
	for l, f := range message {
		if f > frequency {
			letter, frequency = l, f
		}
	}
	return letter
}

func findLeastCommonLetter(message MessageData) string {
	letter, frequency := "", 1000
	for l, f := range message {
		if f < frequency {
			letter, frequency = l, f
		}
	}
	return letter
}

func Solve() (string, string) {
	letters := parseInput()
	message1, message2 := "", ""
	for _, v := range letters {
		message1 += findMostCommonLetter(v)
		message2 += findLeastCommonLetter(v)
	}
	return message1, message2
}
