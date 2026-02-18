package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var inventory = `Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3`

type Item struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

func parse(input string) ([]Item, []Item, []Item) {
	weapons, armors, rings := make([]Item, 0), make([]Item, 0), make([]Item, 0)
	var items *[]Item
	var space bool

	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}

		fields := strings.Fields(line)
		if strings.HasSuffix(fields[0], ":") {
			switch fields[0] {
			case "Weapons:":
				items = &weapons
				space = false
			case "Armor:":
				items = &armors
				space = false
			case "Rings:":
				items = &rings
				space = true
			}
			continue
		}

		name := fields[0]
		if space {
			name += " " + fields[1]
		}
		cost, _ := strconv.Atoi(fields[len(fields)-3])
		damage, _ := strconv.Atoi(fields[len(fields)-2])
		armor, _ := strconv.Atoi(fields[len(fields)-1])

		*items = append(*items, Item{name, cost, damage, armor})
	}
	return weapons, armors, rings
}

type Stats struct {
	Hitpoints int
	Damage    int
	Armor     int
}

func battle(player Stats, monster Stats) bool {
	for player.Hitpoints > 0 && monster.Hitpoints > 0 {
		monster.Hitpoints -= player.Damage - monster.Armor
		if monster.Hitpoints <= 0 {
			break
		}
		player.Hitpoints -= monster.Damage - player.Armor
	}
	return player.Hitpoints > 0
}

func optimize(input string, init int, update func(bool, int, int) int) int {
	weapons, armors, rings := parse(inventory)

	var monster Stats
	fmt.Sscanf(input, "Hit Points: %d\nDamage: %d\nArmor: %d",
		&monster.Hitpoints, &monster.Damage, &monster.Armor)

	best := init
	for i := range weapons {
		// No armor, no rings
		weapon := weapons[i]
		cost := weapon.Cost

		outcome := battle(Stats{100, weapon.Damage, 0}, monster)
		best = update(outcome, cost, best)

		for j := range armors {

			// Weapon and armor, no rings
			armor := armors[j]
			cost = weapon.Cost + armor.Cost

			outcome := battle(Stats{100, weapon.Damage, armor.Armor}, monster)
			best = update(outcome, cost, best)

			for k := range rings {

				// Weapon, armor, one ring.
				ring1 := rings[k]
				cost := weapon.Cost + armor.Cost + ring1.Cost

				outcome := battle(Stats{100, weapon.Damage + ring1.Damage, armor.Armor + ring1.Armor}, monster)
				best = update(outcome, cost, best)

				for m := k + 1; m < len(rings); m++ {

					// Weapon, armor, two rings.
					ring2 := rings[m]
					cost := weapon.Cost + armor.Cost + ring1.Cost + ring2.Cost

					outcome := battle(
						Stats{100,
							weapon.Damage + ring1.Damage + ring2.Damage,
							armor.Armor + ring1.Armor + ring2.Armor},
						monster)
					best = update(outcome, cost, best)
				}
			}
		}
	}
	return best
}

func part1(input string) int {
	return optimize(input, math.MaxInt, func(win bool, cost int, best int) int {
		if !win {
			return best
		}
		return min(cost, best)
	})
}

func part2(input string) int {
	return optimize(input, 0, func(win bool, cost int, best int) int {
		if win {
			return best
		}
		return max(cost, best)
	})
}
