package main

import (
	"testing"
)

var example = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`

func Test_Part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: example,
			want:  50,
		},
		"puzzle": {
			input: puzzle,
			want:  4748769124,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.input), tt.want; got != want {
				t.Errorf("part1() = %v, want = %v", got, want)
			}
		})
	}
}

func Test_Part2(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: example,
			want:  24,
		},
		"puzzle": {
			input: puzzle,
			want:  1525991432,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.input), tt.want; got != want {
				t.Errorf("part2() = %v, want = %v", got, want)
			}
		})
	}
}
