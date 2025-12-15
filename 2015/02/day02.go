package main

import (
	"fmt"
	"slices"
	"strings"
)

func part1(input string) int {
	total := 0
	for line := range strings.SplitSeq(input, "\n") {
		var l, w, h int
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		areas := []int{2 * l * w, 2 * w * h, 2 * h * l}
		extra := slices.Min(areas) / 2
		total += areas[0] + areas[1] + areas[2] + extra
	}
	return total
}

func part2(input string) int {
	total := 0
	for line := range strings.SplitSeq(input, "\n") {
		var l, w, h int
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		sides := []int{l, w, h}
		slices.Sort(sides)
		perimeter := sides[0]*2 + sides[1]*2
		bow := l * w * h
		total += perimeter + bow
	}
	return total
}
