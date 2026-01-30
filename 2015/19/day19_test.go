package main

import (
	_ "embed"
	"testing"
)

var example1 = `H => HO
H => OH
O => HH

HOH`

var example2 = `H => HO
H => OH
O => HH

HOHOHO`

//go:embed puzzle.txt
var puzzle string

func Test_part(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example 1": {example1, 4},
		"example 2": {example2, 7},
		"puzzle":    {puzzle, 518},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.input), tt.want; got != want {
				t.Errorf("part1 got %v, want %v", got, want)
			}
		})
	}
}
