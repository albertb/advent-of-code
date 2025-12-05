package main

import (
	"cmp"
	_ "embed"
	"log"
	"slices"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

type Interval struct {
	lower, upper int
}

func (v Interval) Contains(n int) bool {
	return v.lower <= n && n <= v.upper
}

func (v Interval) Len() int {
	return v.upper - v.lower + 1
}

type Inventory struct {
	fresh       []Interval
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
			inventory.fresh = append(inventory.fresh, Interval{lower, upper})
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
	slices.SortFunc(inv.fresh, func(a, b Interval) int {
		return cmp.Compare(a.lower, b.lower)
	})

	// The interval that we're trying to grow by merging.
	interval := inv.fresh[0]

	count := 0
	for _, fresh := range inv.fresh {
		if fresh.lower <= interval.upper {
			// If these intervals overlap, combine them.
			interval.upper = mathy.Max(interval.upper, fresh.upper)
		} else {
			// Otherwise, add up this interval's length, and continue with the next one.
			count += interval.Len()
			interval = fresh
		}
	}
	// Add up the last interval's length.
	count += interval.Len()

	return count
}

//go:embed puzzle.txt
var puzzle string
