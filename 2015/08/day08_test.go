package main

import (
	_ "embed"
	"testing"
)

var example = `""
"abc"
"aaa\"aaa"
"\x27"`

//go:embed puzzle.txt
var puzzle string

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {example, 12},
		"puzzle":  {puzzle, 1333},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.input), tt.want; got != want {
				t.Errorf("got part1() = %v, want = %v", got, want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {example, 19},
		"puzzle":  {puzzle, 2046},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.input), tt.want; got != want {
				t.Errorf("got part2() = %v, want = %v", got, want)
			}
		})
	}
}
