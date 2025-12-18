package main

import (
	_ "embed"
	"testing"
)

var example = `turn on 0,0 through 999,999
toggle 0,0 through 999,0
turn off 499,499 through 500,500`

//go:embed puzzle.txt
var puzzle string

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"all":     {"turn on 0,0 through 999,999", 1000 * 1000},
		"example": {example, (1000 * 1000) - 1000 - 4},
		"puzzle":  {puzzle, 543903},
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
		"one":     {"turn on 0,0 through 0,0", 1},
		"all":     {"toggle 0,0 through 999,999", 2000000},
		"example": {example, (1000 * 1000) + 1000*2 - 4},
		"puzzle":  {puzzle, 14687245},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.input), tt.want; got != want {
				t.Errorf("got part2() = %v, want = %v", got, want)
			}
		})
	}
}
