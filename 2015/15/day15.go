package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Ingredient struct {
	capacity, durability, flavor, texture, calories int
}

func parse(input string) map[string]Ingredient {
	ingredients := make(map[string]Ingredient)
	for line := range strings.Lines(input) {
		if len(line) < 1 {
			continue
		}
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ":", "")
		fields := strings.Fields(line)

		var g Ingredient
		fmt.Sscanf(fields[2], "%d", &g.capacity)
		fmt.Sscanf(fields[4], "%d", &g.durability)
		fmt.Sscanf(fields[6], "%d", &g.flavor)
		fmt.Sscanf(fields[8], "%d", &g.texture)
		fmt.Sscanf(fields[10], "%d", &g.calories)
		ingredients[fields[0]] = g
	}
	return ingredients
}

func cook(ingredients map[string]Ingredient, meal bool) int {
	names := slices.Collect(maps.Keys(ingredients))
	current := make(map[string]int)

	var best int
	var backtrack func(index, remaining int)
	backtrack = func(index, remaining int) {
		if index == len(names) {
			var cap, dur, fla, tex, cal int
			for name, count := range current {
				cap += count * ingredients[name].capacity
				dur += count * ingredients[name].durability
				fla += count * ingredients[name].flavor
				tex += count * ingredients[name].texture
				cal += count * ingredients[name].calories
			}
			if !meal || cal == 500 {
				score := max(0, cap) * max(0, dur) * max(0, fla) * max(0, tex)
				best = max(best, score)
			}
			return
		}

		key := names[index]
		for count := 0; count <= remaining; count++ {
			current[key] = count
			backtrack(index+1, remaining-count)
		}
		delete(current, key)
	}
	backtrack(0, 100)
	return best
}

func part1(input string) int {
	ingredients := parse(input)
	return cook(ingredients, false)
}

func part2(input string) int {
	ingredients := parse(input)
	return cook(ingredients, true)
}
