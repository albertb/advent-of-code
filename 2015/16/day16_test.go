package main

import (
	_ "embed"
	"testing"
)

//go:embed puzzle.txt
var puzzle string

func Test_part1(t *testing.T) {
	got1, got2 := solve(puzzle)
	want1, want2 := 103, 405
	if got1 != want1 {
		t.Errorf("part1 got %v, want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("part2 got %v, want %v", got2, want2)
	}
}
