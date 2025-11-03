package main

import "testing"

func Test_part1(t *testing.T) {
	var simple = `
XMAS
M..A
A..M
SAMX`

	var example = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "simple",
			input: simple,
			want:  4,
		},

		{
			name:  "example",
			input: example,
			want:  18,
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
	var simple = `
M.S
.A.
M.S`

	var example = `
.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........
`
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "simple",
			input: simple,
			want:  1,
		},
		{
			name:  "example",
			input: example,
			want:  9,
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
