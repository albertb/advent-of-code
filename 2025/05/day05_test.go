package main

import "testing"

var example = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: example,
			want:  3,
		},
		"puzzle": {
			input: puzzle,
			want:  638,
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

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: example,
			want:  14,
		},
		"puzzle": {
			input: puzzle,
			want:  352946349407338,
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
