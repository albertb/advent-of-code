package main

import (
	_ "embed"
	"testing"
)

//go:embed puzzle.txt
var puzzle string

func Test_part1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{`[1,2,3]`, 6},
		{`"a":2,"b":4}`, 6},
		{`[[[3]]]`, 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`[]`, 0},
		{`{}`, 0},
		{puzzle, 191164},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got, want := part1(tt.input), tt.want; got != want {
				t.Errorf("part1 got %v, want %v", got, want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{`[1,2,3]`, 6},
		{`{"a":2,"b":4}`, 6},
		{`[[[3]]]`, 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`[]`, 0},
		{`{}`, 0},
		{`{"a":1, "b":{"c":"red", "d":2}}`, 1},
		{`{"a":1, "b":{"c":"red", "d":2}, "d":[5]}`, 6},
		{`[1, {"a":"red","b":2}, 3]`, 4},
		{puzzle, 87842},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got, want := part2(tt.input), tt.want; got != want {
				t.Errorf("part2 got %v, want %v", got, want)
			}
		})
	}
}
