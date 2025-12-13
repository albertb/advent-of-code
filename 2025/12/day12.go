package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// Every present in puzzle.txt is 3x3. Every region either a) has enough room
// fit presents without overlap OR b) too small to fit presents even with
// overlap.
func part1(input string) int {
	var count int
	for line := range strings.SplitSeq(input, "\n") {
		if !strings.Contains(line, "x") {
			continue
		}
		fields := strings.Fields(line)

		var width, height int
		fmt.Sscanf(fields[0], "%dx%d:", &width, &height)

		// Divide the width and height by three since every present fits in 3x3.
		area := (width / 3) * (height / 3)

		// Now we just need to find out whether we can fit every present as a 1x1 box inside the area.
		presents := 0
		for i := 1; i < len(fields); i++ {
			n, _ := strconv.Atoi(fields[i])
			presents += n
		}

		if area >= presents {
			count++
		}
	}
	return count
}

//go:embed puzzle.txt
var puzzle string
