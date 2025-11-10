package main

import "testing"

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "first",
			input: `
AAAA
BBCD
BBCC
EEEC`,
			want: 140,
		},
		{
			name: "second",
			input: `
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			want: 772,
		},
		{
			name: "larger",
			input: `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
			want: 1930,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
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
			name: "first",
			input: `
AAAA
BBCD
BBCC
EEEC`,
			want: 80,
		},
		{
			name: "second",
			input: `
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			want: 436,
		},
		{
			name: "e-shaped",
			input: `
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			want: 236,
		},
		{
			name: "two b",
			input: `
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`,
			want: 368,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
