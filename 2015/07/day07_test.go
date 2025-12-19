package main

import (
	_ "embed"
	"testing"
)

//go:embed puzzle.txt
var puzzle string

var example = `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
NOT h -> a`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {example, 123},
		"puzzle":  {puzzle, 16076},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.input), tt.want; got != want {
				t.Errorf("got part1() = %v, want = %v", got, want)
			}
		})
	}
}
