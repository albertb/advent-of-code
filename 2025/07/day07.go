package main

import (
	_ "embed"
	"log"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

type Manifold struct {
	start     mathy.Vec
	splitters map[mathy.Vec]struct{}
	bounds    mathy.Bounds
}

func parse(input string) Manifold {
	manifold := Manifold{
		splitters: make(map[mathy.Vec]struct{}),
	}

	y := 0
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		for x, symbol := range line {
			switch symbol {
			case 'S':
				manifold.start = mathy.Vec{x, y}
			case '^':
				manifold.splitters[mathy.Vec{x, y}] = struct{}{}
			case '.':
				// Ignored.
			default:
				log.Fatalln("Bad input:", symbol)
			}
			manifold.bounds.Bound(mathy.Vec{x, y})
		}
		y++
	}
	return manifold
}

func part1(input string) int {
	manifold := parse(input)

	splits := map[mathy.Vec]struct{}{}
	beams := []mathy.Vec{manifold.start}
	visited := map[mathy.Vec]struct{}{}

	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]

		// Don't revisit the same beam location.
		if _, ok := visited[beam]; ok {
			continue
		}
		visited[beam] = struct{}{}

		down := beam.Plus(mathy.Vec{0, 1})
		if _, ok := manifold.splitters[down]; ok {
			// We hit a splitter.
			splits[down] = struct{}{}

			// Add the left beam if possible.
			left := down.Plus(mathy.Vec{-1, 0})
			if manifold.bounds.Contains(left) {
				beams = append(beams, left)
			}

			// Add the right beam if possible.
			right := down.Plus(mathy.Vec{1, 0})
			if manifold.bounds.Contains(right) {
				beams = append(beams, right)
			}

		} else if manifold.bounds.Contains(down) {
			// We hit free space, keep going down.
			beams = append(beams, down)
		}
	}

	return len(splits)
}

func part2(input string) int {
	manifold := parse(input)

	// Count the number of paths that can reach each location.
	paths := map[mathy.Vec]int{
		manifold.start: 1,
	}

	for y := 1; y <= manifold.bounds.Y; y++ {
		for x := 0; x <= manifold.bounds.X; x++ {
			loc := mathy.Vec{x, y}

			beams := paths[loc.Plus(mathy.Vec{0, -1})]
			if beams == 0 {
				// No beam reaches this location.
				continue
			}

			if _, ok := manifold.splitters[loc]; ok {
				// There's a splitter here.
				left := loc.Plus(mathy.Vec{-1, 0})
				paths[left] += beams

				right := loc.Plus(mathy.Vec{1, 0})
				paths[right] += beams
			} else {
				// No splitter, keep going down.
				paths[loc] += beams
			}
		}
	}

	// Sum up the number of paths that can reach the last row.
	count := 0
	for x := 0; x <= manifold.bounds.X; x++ {
		count += paths[mathy.Vec{x, manifold.bounds.Y}]
	}
	return count
}

//go:embed puzzle.txt
var puzzle string
