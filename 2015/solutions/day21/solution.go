package day21

import "math"

type ShopItem struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

type Weapon = ShopItem
type Armor = ShopItem
type Ring = ShopItem

var weapons = []Weapon{
	{
		Name:   "Dagger",
		Cost:   8,
		Damage: 4,
		Armor:  0,
	},
	{
		Name:   "Shortsword",
		Cost:   10,
		Damage: 5,
		Armor:  0,
	},
	{
		Name:   "Warhammer",
		Cost:   25,
		Damage: 6,
		Armor:  0,
	},
	{
		Name:   "Longsword",
		Cost:   40,
		Damage: 7,
		Armor:  0,
	},
	{
		Name:   "Greataxe",
		Cost:   74,
		Damage: 8,
		Armor:  0,
	},
}

var armor = []*Armor{
	nil,
	{
		Name:   "Leather",
		Cost:   13,
		Damage: 0,
		Armor:  1,
	},
	{
		Name:   "Chainmail",
		Cost:   31,
		Damage: 0,
		Armor:  2,
	},
	{
		Name:   "Splintmail",
		Cost:   53,
		Damage: 0,
		Armor:  3,
	},
	{
		Name:   "Bandedmail",
		Cost:   75,
		Damage: 0,
		Armor:  4,
	},
	{
		Name:   "Platemail",
		Cost:   102,
		Damage: 0,
		Armor:  5,
	},
}

var rings = []*Ring{
	nil,
	{
		Name:   "Damage +1",
		Cost:   25,
		Damage: 1,
		Armor:  0,
	},
	{
		Name:   "Damage +2",
		Cost:   50,
		Damage: 2,
		Armor:  0,
	},
	{
		Name:   "Damage +3",
		Cost:   100,
		Damage: 3,
		Armor:  0,
	},
	{
		Name:   "Defense +1",
		Cost:   20,
		Damage: 0,
		Armor:  1,
	},
	{
		Name:   "Defense +2",
		Cost:   40,
		Damage: 0,
		Armor:  2,
	},
	{
		Name:   "Defense +3",
		Cost:   80,
		Damage: 0,
		Armor:  3,
	},
}

type Player struct {
	HitPoints int
	Weapon    Weapon
	Armor     *Armor
	LeftRing  *Ring
	RightRing *Ring
}

func (p *Player) GetDamage() (damage int) {
	damage += p.Weapon.Damage
	if p.Armor != nil {
		damage += p.Armor.Damage
	}
	if p.LeftRing != nil {
		damage += p.LeftRing.Damage
	}
	if p.RightRing != nil {
		damage += p.RightRing.Damage
	}
	return
}

func (p *Player) GetArmor() (armor int) {
	armor += p.Weapon.Armor
	if p.Armor != nil {
		armor += p.Armor.Armor
	}
	if p.LeftRing != nil {
		armor += p.LeftRing.Armor
	}
	if p.RightRing != nil {
		armor += p.RightRing.Armor
	}
	return
}

func (p *Player) GetCost() (cost int) {
	cost += p.Weapon.Cost
	if p.Armor != nil {
		cost += p.Armor.Cost
	}
	if p.LeftRing != nil {
		cost += p.LeftRing.Cost
	}
	if p.RightRing != nil {
		cost += p.RightRing.Cost
	}
	return
}

func createBoss() Player {
	return Player{
		HitPoints: 104,
		Weapon:    weapons[4],
		Armor:     armor[1],
	}
}

func attackBoss(player Player, boss Player) bool {
	for {
		playerDamage := player.GetDamage() - boss.GetArmor()
		if playerDamage < 1 {
			playerDamage = 1
		}
		if boss.HitPoints -= playerDamage; boss.HitPoints <= 0 {
			return true
		}

		bossDamage := boss.GetDamage() - player.GetArmor()
		if player.HitPoints -= bossDamage; player.HitPoints <= 0 {
			return false
		}
	}
}

func Solve() (int, int) {
	minCost := math.MaxInt64
	maxCost := 0
	for i := 0; i < len(weapons); i++ {
		for j := 0; j < len(armor); j++ {
			for k1 := 0; k1 < len(rings); k1++ {
				for k2 := k1; k2 < len(rings); k2++ {
					if k2 > 0 && k1 == k2 {
						continue
					}

					player := Player{
						HitPoints: 100,
						Weapon:    weapons[i],
						Armor:     armor[j],
						LeftRing:  rings[k1],
						RightRing: rings[k2],
					}

					cost := player.GetCost()
					if attackBoss(player, createBoss()) {
						if cost < minCost {
							minCost = cost
						}
					} else {
						if cost > maxCost {
							maxCost = cost
						}
					}
				}
			}
		}
	}

	return minCost, maxCost
}
