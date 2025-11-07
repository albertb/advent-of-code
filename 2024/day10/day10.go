package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:", part1(input))
	//fmt.Println("Part 2:", part2(input))
}

func parse(input string) [][]int {
	var grid [][]int
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}

		row := []int{}
		for digit := range strings.SplitSeq(line, "") {
			if digit == "." {
				row = append(row, -1)
			} else {
				height, err := strconv.ParseInt(digit, 10, 16)
				if err != nil {
					log.Fatalln(err)
				}
				row = append(row, int(height))
			}
		}
		grid = append(grid, row)
	}
	return grid
}

type Coord struct {
	x, y int
}

func route(grid [][]int, x, y, height int, tops map[Coord]struct{}) {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]) {
		// Out of bounds, give up.
		return
	}

	if grid[y][x] != height {
		// Unexpected height, we cannot reach the top from here.
		return
	}

	if height == 9 {
		// We reached the top, add this top coordinate.
		tops[Coord{x, y}] = struct{}{}
		return
	}

	// Route to each direction: left, up, right, down, looking for the next height.
	for _, delta := range [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
		route(grid, x+delta[0], y+delta[1], height+1, tops)
	}
}

func part1(input string) int {
	var sum int

	grid := parse(input)
	for y, row := range grid {
		for x, height := range row {
			if height == 0 {
				tops := map[Coord]struct{}{}
				route(grid, x, y, 0, tops)
				sum += len(tops)
			}
		}
	}
	return sum
}
