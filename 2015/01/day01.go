package main

import (
	_ "embed"
	"fmt"
)

func part1(input string) int {
	floor := 0
	for _, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor
}

func part2(input string) int {
	floor := 0
	for i, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return -1
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("part 1:", part1(puzzle))
	fmt.Println("part 2:", part2(puzzle))
}
