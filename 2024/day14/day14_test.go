package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name  string
		input string
		space Vec
		iters int
		want  int
	}{
		{
			name:  "simple",
			input: `p=2,4 v=2,-3`,
			space: Vec{11, 7},
			iters: 5,
			want:  0,
		},
		{
			name: "example",
			input: `
			p=0,4 v=3,-3
			p=6,3 v=-1,-3
			p=10,3 v=-1,2
			p=2,0 v=2,-1
			p=0,0 v=1,3
			p=3,0 v=-2,-2
			p=7,6 v=-1,-3
			p=3,0 v=-1,-2
			p=9,3 v=2,3
			p=7,3 v=-1,2
			p=2,4 v=2,-3
			p=9,5 v=-3,-3`,
			space: Vec{11, 7},
			iters: 100,
			want:  12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, tt.space, tt.iters); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
