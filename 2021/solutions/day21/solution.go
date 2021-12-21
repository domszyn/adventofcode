package day21

import (
	"fmt"

	tm "github.com/buger/goterm"
)

type Dice struct {
	Value  int
	Rolled int
}

func (d *Dice) Roll() int {
	d.Value = (3*(d.Rolled+2)-1)%100 + 1
	d.Rolled += 3
	return d.Value
}

type Player struct {
	Position int
	Score    int
}

func (p *Player) Move(roll int) Player {
	position := (p.Position+roll-1)%10 + 1
	return Player{position, p.Score + position}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Solve() (part1, part2 int) {
	players := [2]Player{{Position: 10}, {Position: 6}}
	dice := Dice{}

	for {
		players[0] = players[0].Move(dice.Roll())
		if players[0].Score >= 1000 {
			break
		}

		players[1] = players[1].Move(dice.Roll())
		if players[1].Score >= 1000 {
			break
		}
	}

	part1 = min(players[0].Score, players[1].Score) * dice.Rolled

	rolls := make(map[int]int)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				rolls[i+j+k+3]++
			}
		}
	}
	moves := make(map[[2]Player]int)
	moves[[2]Player{{Position: 10}, {Position: 6}}] = 1

	p1win, p2win := 0, 0
	for step := 1; len(moves) > 0; step++ {
		tm.MoveCursor(1, 1)
		tm.Println(fmt.Sprintf("Step %02d", step))
		nextMoves := make(map[[2]Player]int)
		for p, v := range moves {
			for r, mf := range rolls {
				nextMoves[[2]Player{p[0].Move(r), p[1]}] += v * mf
			}
		}
		moves = nextMoves

		for p, v := range moves {
			if p[0].Score > 20 {
				p1win += v
				delete(moves, p)
			}
		}

		nextMoves = make(map[[2]Player]int)
		for p, v := range moves {
			for r, mf := range rolls {
				nextMoves[[2]Player{p[0], p[1].Move(r)}] += v * mf
			}
		}
		moves = nextMoves

		for p, v := range moves {
			if p[1].Score > 20 {
				p2win += v
				delete(moves, p)
			}
		}
	}

	part2 = max(p1win, p2win)
	return
}
