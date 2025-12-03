package main

import "testing"

var example = `987654321111111
811111111111119
234234234234278
818181911112111`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: example,
			want:  357,
		},
		"puzzle": {
			input: puzzle,
			want:  17405,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.input), tt.want; got != want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
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
			want:  3121910778619,
		},
		"puzzle": {
			input: puzzle,
			want:  171990312704598,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.input), tt.want; got != want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
