package main

import (
	"fmt"
	"image"
	"log"
	"strings"
)

type State struct {
	on   map[image.Point]struct{}
	grid image.Rectangle
}

func parse(input string) State {
	state := State{
		on: make(map[image.Point]struct{}),
	}

	var y int
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		for x, sym := range line {
			switch sym {
			case '#':
				state.on[image.Pt(x, y)] = struct{}{}
			case '.':
				// Ignored.
			default:
				log.Fatalln("unrecognized symbol at", x, y, ":", string(sym))
			}
		}
		state.grid.Max.X = max(state.grid.Max.X, len(line)-1)
		state.grid.Max.Y = max(state.grid.Max.Y, y)
		y++
	}
	return state
}

func run(state State, steps int, corners bool) int {
	last := state

	if corners {
		x, y := state.grid.Max.X, state.grid.Max.Y
		for _, corner := range []image.Point{{0, 0}, {0, y}, {x, 0}, {x, y}} {
			last.on[corner] = struct{}{}
		}
	}

	for range steps {
		next := State{
			on:   make(map[image.Point]struct{}),
			grid: last.grid,
		}

		for y := range state.grid.Max.Y + 1 {
		light:
			for x := range state.grid.Max.X + 1 {
				loc := image.Pt(x, y)

				if corners {
					x, y := state.grid.Max.X, state.grid.Max.Y
					for _, corner := range []image.Point{{0, 0}, {0, y}, {x, 0}, {x, y}} {
						if loc.Eq(corner) {
							next.on[loc] = struct{}{}
							continue light
						}
					}
				}

				var lit int
				for _, off := range []image.Point{
					{-1, -1}, {0, -1}, {1, -1},
					{-1, 0} /*{0,0}*/, {1, 0},
					{-1, 1}, {0, 1}, {1, 1}} {
					if _, ok := last.on[loc.Add(off)]; ok {
						lit++
					}
				}

				if _, ok := last.on[loc]; ok {
					// A light which is on stays on when 2 or 3 neighbours are on.
					if lit == 2 || lit == 3 {
						next.on[loc] = struct{}{}
					}
				} else {
					// A light which is off turns on if exactly 3 neighbours are on.
					if lit == 3 {
						next.on[loc] = struct{}{}
					}
				}
			}
			fmt.Println()
		}
		last = next
	}
	return len(last.on)
}

func part1(grid string, steps int) int {
	state := parse(grid)
	return run(state, steps, false)
}

func part2(grid string, steps int) int {
	state := parse(grid)
	return run(state, steps, true)
}
