package main

import "testing"

var example = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: example,
			want:  4277556,
		},
		"puzzle": {
			input: puzzle,
			want:  3525371263915,
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
			want:  3263827,
		},
		"puzzle": {
			input: puzzle,
			want:  6846480843636,
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
