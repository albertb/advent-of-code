package main

import (
	_ "embed"
	"testing"
)

var example = `20
15
10
5
5`

//go:embed puzzle.txt
var puzzle string

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		capacities string
		quanity    int
		want       int
	}{
		"example": {example, 25, 4},
		"puzzle":  {puzzle, 150, 1638},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.capacities, tt.quanity), tt.want; got != want {
				t.Errorf("part1 got %v, want %v", got, want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		capacities string
		quanity    int
		want       int
	}{
		"example": {example, 25, 3},
		"puzzle":  {puzzle, 150, 17},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.capacities, tt.quanity), tt.want; got != want {
				t.Errorf("part2 got %v, want %v", got, want)
			}
		})
	}
}
