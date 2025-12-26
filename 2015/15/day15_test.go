package main

import (
	_ "embed"
	"testing"
)

var example = `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`

//go:embed puzzle.txt
var puzzle string

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {example, 62842880},
		"puzzle":  {puzzle, 18965440},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.input), tt.want; got != want {
				t.Errorf("part1 got %v, want %v", got, want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {example, 57600000},
		"puzzle":  {puzzle, 15862900},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.input), tt.want; got != want {
				t.Errorf("part2 got %v, want %v", got, want)
			}
		})
	}
}
