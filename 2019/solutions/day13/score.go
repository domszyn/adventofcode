package day13

import (
	"fmt"

	tm "github.com/buger/goterm"
)

type Score struct {
	Score, Hiscore int
}

func (s *Score) Set(score int) {
	s.Score = score
	if score > s.Hiscore {
		s.Hiscore = score
	}
}

func (s Score) Endgame() bool {
	return s.Score > 10000 && s.Score < 10552
}

func (s Score) Print() {
	tm.MoveCursor(1, 2)
	tm.Println(tm.Bold(fmt.Sprintf("HiScore %s", tm.Color(fmt.Sprintf("%05d", s.Hiscore), tm.RED))))
	tm.Flush()
	tm.MoveCursor(1, 3)
	tm.Println(tm.Bold(fmt.Sprintf("Score   %s", tm.Color(fmt.Sprintf("%05d", s.Score), tm.MAGENTA))))
	tm.Flush()
}

func printRound(round int) {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Println(tm.Bold(fmt.Sprintf("Round   %s", tm.Color(fmt.Sprintf("%05d", round+1), tm.CYAN))))
	tm.Flush()
}

func printTile(x, y, id int) {
	tm.MoveCursor(x+1, y+4)
	switch id {
	case EmptyTile:
		tm.Print(" ")
	case Wall:
		tm.Print(tm.Bold("#"))
	case Block:
		tm.Print(tm.Bold(tm.Color("*", tm.YELLOW)))
	case Paddle:
		tm.Print(tm.Bold(tm.Color("=", tm.GREEN)))
	case Ball:
		tm.Print(tm.Bold(tm.Color("O", tm.BLUE)))
	}
	tm.Flush()
}

func printGameOver() {
	for i, v := range "GAME OVER!" {
		tm.MoveCursor(20+i, i%3+1)
		tm.Print(tm.Color(tm.Bold(string(v)), tm.RED))
	}
	tm.Flush()
}

func printWin() {
	tm.MoveCursor(20, 1)
	tm.Print(tm.Color(tm.Bold("YOU WON!!!"), tm.GREEN))
	tm.Flush()
}
