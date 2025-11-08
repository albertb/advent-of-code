package main

import "testing"

func Test_part1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		blinks int
		want   int
	}{
		{
			name:   "example",
			input:  "0 1 10 99 999",
			blinks: 1,
			want:   7,
		},
		{
			name:   "longer",
			input:  "125 17",
			blinks: 6,
			want:   22,
		},
		{
			name:   "longer25",
			input:  "125 17",
			blinks: 25,
			want:   55312,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, tt.blinks); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}

			if got := part2(tt.input, tt.blinks); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
