package main

import (
	_ "embed"
	"testing"
)

//go:embed puzzle.txt
var puzzle string

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: "2x3x4",
			want:  58,
		},
		"puzzle": {
			input: puzzle,
			want:  1598415,
		},
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
		"example1": {
			input: "2x3x4",
			want:  34,
		},
		"example2": {
			input: "1x1x10",
			want:  14,
		},
		"puzzle": {
			input: puzzle,
			want:  3812909,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.input), tt.want; got != want {
				t.Errorf("got part2() = %v, want = %v", got, want)
			}
		})
	}
}
