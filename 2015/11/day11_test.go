package main

import "testing"

func Test_solve(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
		{"hxbxwxba", "hxbxxyzz"},
		{"hxbxxyzz", "hxcaabcc"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got, want := solve(tt.input), tt.want; got != want {
				t.Errorf("part1 got %v, want %v", got, want)
			}
		})
	}
}
