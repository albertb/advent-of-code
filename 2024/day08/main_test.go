package main

import "testing"

var simple = `
....
.a..
..a.
....`

var example = `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

var simple2 = `
T.........
...T......
.T........
..........
..........
..........
..........
..........
..........
..........
`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "simple",
			input: simple,
			want:  2,
		},
		{
			name:  "example",
			input: example,
			want:  14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part(tt.input, false); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "simple2",
			input: simple2,
			want:  6,
		},
		{
			name:  "example",
			input: example,
			want:  34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part(tt.input, true); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
