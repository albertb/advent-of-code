package main

import (
	_ "embed"
	"testing"
)

var example = `.#.#.#
...##.
#....#
..#...
#.#..#
####..`

//go:embed puzzle.txt
var puzzle string

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		grid  string
		steps int
		want  int
	}{
		"example": {example, 4, 4},
		"puzzle":  {puzzle, 100, 821},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.grid, tt.steps), tt.want; got != want {
				t.Errorf("part1 got %v, want %v", got, want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		grid  string
		steps int
		want  int
	}{
		"example": {example, 5, 17},
		"puzzle":  {puzzle, 100, 886},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.grid, tt.steps), tt.want; got != want {
				t.Errorf("part2 got %v, want %v", got, want)
			}
		})
	}
}
