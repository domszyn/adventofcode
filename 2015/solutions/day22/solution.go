package day22

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	HitPoints int
	Mana      int
	Armor     int
	Spells    []Spell
}

type Boss struct {
	HitPoints int
	Damage    int
}

type Spell struct {
	Name      string
	Cost      int
	Damage    int
	HitPoints int
	Effect    string
}

type Effect struct {
	Duration  int
	Remaining int
	Active    bool
	Proc      func(p *Player, b *Boss)
}

func (p *Player) AvailableSpells(effects map[string]*Effect) (spells []Spell, bail bool) {
	spells = []Spell{}
	for _, spell := range p.Spells {
		if spell.Cost < p.Mana && !effects[spell.Effect].Active {
			spells = append(spells, spell)
		}
	}
	bail = len(spells) == 0
	return
}

func NewEffect(duration int, proc func(p *Player, b *Boss)) *Effect {
	ef := Effect{duration, 0, false, proc}
	return &ef
}

func (e *Effect) String() string {
	if e.Active {
		return fmt.Sprintf("ACTIVE at %d/%d", e.Remaining, e.Duration)
	}

	return fmt.Sprintf("INACTIVE")
}

func (e *Effect) Apply(p *Player, b *Boss) {
	e.Proc(p, b)
	e.Remaining--
	if e.Remaining == 0 {
		e.Active = false
	}
}

func ApplyEffects(p *Player, b *Boss, fx map[string]*Effect) {
	for _, e := range fx {
		if e.Active {
			e.Apply(p, b)
		}
	}
}

func ApplySpell(s Spell, p *Player, b *Boss, fx map[string]*Effect) int {
	p.Mana -= s.Cost
	p.HitPoints += s.HitPoints
	b.HitPoints -= s.Damage
	if s.Effect != "None" {
		if fx[s.Effect].Active {
			panic("trying to activate an already-active status; something has gone awry")
		}
		fx[s.Effect].Active = true
		fx[s.Effect].Remaining = fx[s.Effect].Duration
	}
	return s.Cost
}

func Round(p *Player, b *Boss, fx map[string]*Effect) (continueFight bool, bail bool, manaSpent int) {
	// player's turn

	if p.HitPoints--; p.HitPoints == 0 {
		return false, false, manaSpent
	}

	ApplyEffects(p, b, fx)

	validSpells, bail := p.AvailableSpells(fx)
	if bail {
		return false, true, manaSpent
	}

	choice := validSpells[rand.Intn(len(validSpells))]

	manaSpent = ApplySpell(choice, p, b, fx)

	p.Armor = 0

	// boss's turn
	ApplyEffects(p, b, fx)
	if b.HitPoints > 0 {
		dmg := b.Damage - p.Armor
		if dmg < 1 {
			dmg = 1
		}
		p.HitPoints -= dmg
	}

	p.Armor = 0

	if b.HitPoints <= 0 || p.HitPoints <= 0 {
		return false, false, manaSpent
	}
	return true, false, manaSpent
}

func Fight(p Player, b Boss) (playerWon bool, manaCost int) {
	fx := map[string]*Effect{
		"None":     NewEffect(0, func(p *Player, b *Boss) {}),
		"Poison":   NewEffect(6, func(p *Player, b *Boss) { b.HitPoints -= 3 }),
		"Shield":   NewEffect(6, func(p *Player, b *Boss) { p.Armor = 7 }),
		"Recharge": NewEffect(5, func(p *Player, b *Boss) { p.Mana += 101 }),
	}

	continueBattle := true
	manaSpentThatTurn := 0
	bail := false

	for continueBattle && !bail {
		continueBattle, bail, manaSpentThatTurn = Round(&p, &b, fx)
		manaCost = manaCost + manaSpentThatTurn
	}
	return p.HitPoints > 0 && !bail, manaCost
}

func SolvePart1() int {
	spells := []Spell{
		{"Magic Missile", 53, 4, 0, "None"},
		{"Drain", 73, 2, 2, "None"},
		{"Shield", 113, 0, 0, "Shield"},
		{"Poison", 173, 0, 0, "Poison"},
		{"Recharge", 229, 0, 0, "Recharge"},
	}
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	player := Player{HitPoints: 50, Armor: 0, Mana: 500, Spells: spells}
	boss := Boss{HitPoints: 58, Damage: 9}
	bestCost := 999999999
	costs := make(map[int]int)
	for {
		won, cost := Fight(player, boss)
		if won && (cost < bestCost) {
			bestCost = cost
			costs[cost]++
			fmt.Println(cost)
		}

		if costs[cost] > 1 {
			return bestCost
		}
	}
}
