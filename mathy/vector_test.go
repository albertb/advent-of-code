package mathy

import "testing"

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		vec  Vec
		n    int
		want Vec
	}{
		{
			name: "90 clockwise",
			vec:  Vec{3, 2},
			n:    1,
			want: Vec{-2, 3},
		},
		{
			name: "90 counter-clockwise",
			vec:  Vec{3, 2},
			n:    -1,
			want: Vec{2, -3},
		},
		{
			name: "360",
			vec:  Vec{123, 456},
			n:    4,
			want: Vec{123, 456},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vec.Rotate90(tt.n); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
