package main

import (
	"fmt"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		input  string
		rounds int
		want   int
	}{
		{"1", 1, len("11")},
		{"1", 2, len("21")},
		{"1", 3, len("1211")},
		{"1", 4, len("312211")},
		{"3113322113", 40, 329356},
		{"3113322113", 50, 4666278},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s @ %d", tt.input, tt.rounds), func(t *testing.T) {
			if got, want := solve(tt.input, tt.rounds), tt.want; got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
