package main

import (
	_ "embed"
	"testing"
)

var example1a = `H => HO
H => OH
O => HH

HOH`

var example1b = `H => HO
H => OH
O => HH

HOHOHO`

//go:embed puzzle.txt
var puzzle string

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example 1a": {example1a, 4},
		"example 1b": {example1b, 7},
		"puzzle":     {puzzle, 518},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.input), tt.want; got != want {
				t.Errorf("part1 got %v, want %v", got, want)
			}
		})
	}
}

var example2a = `e => H
e => O
H => HO
H => OH
O => HH

HOH`

var example2b = `e => H
e => O
H => HO
H => OH
O => HH

HOHOHO`

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example 2a": {example2a, 3},
		"example 2b": {example2b, 6},
		"puzzle":     {puzzle, 200},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.input), tt.want; got != want {
				t.Errorf("part2 got %v, want %v", got, want)
			}
		})
	}
}
