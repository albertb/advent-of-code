package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2024/day04/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	contents, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Part 1 count:", part1(string(contents)))
	fmt.Println("Part 2 count:", part2(string(contents)))
}

func makeGrid(puzzle string) [][]rune {
	reader := strings.NewReader(puzzle)
	scanner := bufio.NewScanner(reader)

	grid := [][]rune{}
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune{})
		for _, letter := range line {
			grid[lineNumber] = append(grid[lineNumber], letter)
		}
		lineNumber++
	}

	return grid
}

func part1(puzzle string) int {
	grid := makeGrid(puzzle)

	count := 0
	for x := range len(grid) {
		for y := range len(grid[x]) {
			count += search(grid, x, y, 0, 0, 0)
		}
	}
	return count
}

func search(grid [][]rune, x, y, dx, dy, n int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
		return 0
	}

	if n == 0 {
		if grid[x][y] != 'X' {
			return 0
		}

		count := 0
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if search(grid, x+dx, y+dy, dx, dy, n+1) > 0 {
					//fmt.Println("FOUND at", x, ",", y, "with", dx, ",", dy)
					count++
				}
			}
		}
		return count
	}

	if n == 1 {
		if grid[x][y] != 'M' {
			return 0
		}
		return search(grid, x+dx, y+dy, dx, dy, n+1)
	}

	if n == 2 {
		if grid[x][y] != 'A' {
			return 0
		}
		return search(grid, x+dx, y+dy, dx, dy, n+1)

	}

	if n == 3 && grid[x][y] == 'S' {
		return 1
	}

	return 0
}

func part2(puzzle string) int {
	grid := makeGrid(puzzle)

	search := func(grid [][]rune, x, y int) bool {
		if x < 1 || y < 1 || x > len(grid)-2 || y > len(grid[x])-2 {
			return false
		}

		if grid[x][y] != 'A' {
			return false
		}

		tl := grid[x-1][y-1]
		br := grid[x+1][y+1]
		if !((tl == 'M' && br == 'S') || (tl == 'S' && br == 'M')) {
			return false
		}

		tr := grid[x-1][y+1]
		bl := grid[x+1][y-1]
		if !((tr == 'M' && bl == 'S') || (tr == 'S' && bl == 'M')) {
			return false
		}

		return true
	}

	count := 0
	for x := range len(grid) {
		for y := range len(grid[x]) {
			if search(grid, x, y) {
				count++
			}
		}
	}
	return count
}
