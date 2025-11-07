package main

import "testing"

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `
0123
1234
8765
9876`,
			want: 1,
		},
		{
			name: "example2",
			input: `
...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`,

			want: 2,
		},
		{
			name: "example3",
			input: `
..90..9
...1.98
...2..7
6543456
765.987
876....
987....`,
			want: 4,
		},
		{
			name: "example4",
			input: `
10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01`,
			want: 3,
		},
		{
			name: "example5",
			input: `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			want: 36,
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
