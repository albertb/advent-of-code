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
		"1":      {"ugknbfddgicrmopn", 1},
		"2":      {"aaa", 1},
		"3":      {"jchzalrnumimnmhp", 0},
		"4":      {"haegwjzuvuyypxyu", 0},
		"5":      {"dvszwmarrgswjxmb", 0},
		"badstr": {"pqaaa", 0},
		"puzzle": {puzzle, 258},
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
		"1":      {"qjhvhtzxzqqjkmpb", 1},
		"2":      {"xxyxx", 1},
		"3":      {"uurcxstgmygtbstg", 0},
		"4":      {"ieodomkazucvgmuy", 0},
		"puzzle": {puzzle, 53},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.input), tt.want; got != want {
				t.Errorf("got part2() = %v, want = %v", got, want)
			}
		})
	}
}
