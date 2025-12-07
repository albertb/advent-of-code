package main

import "testing"

var example = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: example,
			want:  21,
		},
		"puzzle": {
			input: puzzle,
			want:  1609,
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

var simple = `
..S..
.....
..^..
.....
.^.^.
.....`

var simple2 = `
...S...
.......
...^...
.......
..^.^..
.......
.^.^.^.
.......
`

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"simple": {
			input: simple,
			want:  4,
		},
		"simple2": {
			input: simple2,
			want:  8,
		},
		"example": {
			input: example,
			want:  40,
		},
		"puzzle": {
			input: puzzle,
			want:  12472142047197,
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
