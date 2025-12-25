package main

import (
	_ "embed"
	"testing"
)

var example = `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

//go:embed puzzle.txt
var puzzle string

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		time  int
		want  int
	}{
		"example": {example, 1000, 1120},
		"puzzle":  {puzzle, 2503, 2640},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.input, tt.time), tt.want; got != want {
				t.Errorf("part1 got %v, want %v", got, want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		input string
		time  int
		want  int
	}{
		"ex1":    {example, 1, 1},
		"ex2":    {example, 140, 139},
		"ex3":    {example, 1000, 689},
		"puzzle": {puzzle, 2503, 1102},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.input, tt.time), tt.want; got != want {
				t.Errorf("part2 got %v, want %v", got, want)
			}
		})
	}
}
