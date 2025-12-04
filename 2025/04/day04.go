package main

import (
	_ "embed"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

type Floor struct {
	rolls  map[mathy.Vec]struct{}
	bounds mathy.Bounds
}

func parse(input string) Floor {
	var floor Floor
	floor.rolls = map[mathy.Vec]struct{}{}

	y := 0
	for line := range strings.SplitSeq(input, "\n") {
		x := 0
		for _, sym := range line {
			if sym == '@' {
				floor.rolls[mathy.Vec{X: x, Y: y}] = struct{}{}
			}
			if x > floor.bounds.X {
				floor.bounds.X = x
			}
			x++
		}
		if y > floor.bounds.Y {
			floor.bounds.Y = y
		}
		y++
	}
	return floor
}

func freeNeighbors(roll mathy.Vec, floor Floor) bool {
	neighbors := 0
	for _, diff := range []mathy.Vec{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0} /*{0, 0}*/, {1, 0},
		{-1, 1}, {0, 1}, {1, 1}} {
		spot := roll.Plus(diff)
		if !floor.bounds.Contains(spot) {
			continue
		}
		if _, ok := floor.rolls[spot]; ok {
			neighbors++
			if neighbors > 3 {
				return false
			}
		}
	}
	return true
}

func part1(input string) int {
	floor := parse(input)

	count := 0
	for roll := range floor.rolls {
		if freeNeighbors(roll, floor) {
			count++
		}
	}
	return count
}

func part2(input string) int {
	floor := parse(input)

	removed := 0
	for {
		done := true
		for roll := range floor.rolls {
			if freeNeighbors(roll, floor) {
				removed++
				delete(floor.rolls, roll)
				done = false
			}
		}
		if done {
			break
		}
	}
	return removed
}

//go:embed puzzle.txt
var puzzle string
