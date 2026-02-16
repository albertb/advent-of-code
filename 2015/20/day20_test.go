package main

import "testing"

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"house1": {10, 1},
		"house2": {30, 2},
		"house3": {40, 3},
		"house4": {70, 4},
		"puzzle": {36000000, 831600},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.input), tt.want; got != want {
				t.Errorf("part1 got %v, want %v", got, want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"house1": {11, 1},
		"house2": {11 + 22, 2},
		"puzzle": {36000000, 884520},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.input), tt.want; got != want {
				t.Errorf("part1 got %v, want %v", got, want)
			}
		})
	}
}
