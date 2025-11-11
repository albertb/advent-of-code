package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "easy",
			input: `
Button A: X+5, Y+2
Button B: X+4, Y+1
Prize: X=14, Y=5`,
			want: 7,
		},
		{
			name: "first",
			input: `
Button A: X+94, Y+34,
Button B: X+22, Y+67
Prize: X=8400, Y=5400`,
			want: 280,
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
