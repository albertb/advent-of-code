package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"1":      {"abcdef", 609043},
		"2":      {"pqrstuv", 1048970},
		"puzzle": {"iwrupvqb", 346386},
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
		"puzzle": {"iwrupvqb", 9958218},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.input), tt.want; got != want {
				t.Errorf("got part2() = %v, want = %v", got, want)
			}
		})
	}
}
