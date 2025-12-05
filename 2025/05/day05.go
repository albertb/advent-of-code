package main

import (
	"cmp"
	_ "embed"
	"log"
	"slices"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

type Range struct {
	lower, upper int
}

func (r Range) Contains(v int) bool {
	return r.lower <= v && v <= r.upper
}

func (r Range) Len() int {
	return r.upper - r.lower + 1
}

type Inventory struct {
	fresh       []Range
	ingredients []int
}

func parse(input string) Inventory {
	var inventory Inventory
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		if strings.Contains(line, "-") {
			bounds := strings.Split(line, "-")
			if len(bounds) != 2 {
				log.Fatalln("unexpected fresh bounds:", line)
			}
			lower := mathy.MustParseInt(bounds[0])
			upper := mathy.MustParseInt(bounds[1])
			inventory.fresh = append(inventory.fresh, Range{lower, upper})
		} else {
			ingredient := mathy.MustParseInt(line)
			inventory.ingredients = append(inventory.ingredients, ingredient)
		}
	}
	return inventory
}

func part1(input string) int {
	inv := parse(input)

	count := 0
	for _, ingredient := range inv.ingredients {
		for _, fresh := range inv.fresh {
			if fresh.Contains(ingredient) {
				count++
				break
			}
		}
	}
	return count
}

func part2(input string) int {
	inv := parse(input)

	// Sort the ranges by their lower bound.
	slices.SortFunc(inv.fresh, func(a, b Range) int {
		return cmp.Compare(a.lower, b.lower)
	})

	rg := inv.fresh[0]

	count := 0
	for _, r := range inv.fresh {
		if r.lower <= rg.upper {
			// If these ranges overlap, combine them.
			rg.upper = mathy.Max(rg.upper, r.upper)
		} else {
			// Otherwise, add up this range's length, and continue with the next range.
			count += rg.Len()
			rg = r
		}
	}
	// Add up the last range's length.
	count += rg.Len()

	return count
}

//go:embed puzzle.txt
var puzzle string
