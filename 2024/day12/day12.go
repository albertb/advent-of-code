package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type Crop string

type Plot struct {
	x, y int
}

func parse(input string) map[Plot]Crop {
	plots := map[Plot]Crop{}

	y := 0
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			continue
		}

		x := 0
		for crop := range strings.SplitSeq(line, "") {
			plots[Plot{x, y}] = Crop(crop)
			x++
		}
		y++
	}

	return plots
}

type Region struct {
	crop      Crop
	plots     []Plot
	perimeter int
}

func explore(plots map[Plot]Crop, plot Plot, want Crop, visited map[Plot]struct{}) Region {
	crop, ok := plots[plot]
	if !ok {
		// Out of bounds.
		return Region{perimeter: 1}
	}

	if crop != want {
		// Wrong label.
		return Region{perimeter: 1}
	}

	if _, ok := visited[plot]; ok {
		// Already visited.
		return Region{}
	}

	// Mark this plot as visited.
	visited[plot] = struct{}{}

	// Explore neighbouring plots.
	left := explore(plots, Plot{plot.x - 1, plot.y}, want, visited)
	up := explore(plots, Plot{plot.x, plot.y - 1}, want, visited)
	right := explore(plots, Plot{plot.x + 1, plot.y}, want, visited)
	down := explore(plots, Plot{plot.x, plot.y + 1}, want, visited)

	// Merge the regions together.
	region := Region{
		crop:      want,
		plots:     []Plot{plot},
		perimeter: 0,
	}

	for _, other := range []Region{left, up, right, down} {
		region.plots = append(region.plots, other.plots...)
		region.perimeter += other.perimeter
	}
	return region
}

func part1(input string) int {
	plots := parse(input)

	visited := map[Plot]struct{}{}
	regions := []Region{}

	x := 0
	for {
		if _, ok := plots[Plot{x, 0}]; !ok {
			// Out of bounds.
			break
		}

		y := 0
		for {
			plot := Plot{x, y}

			crop, ok := plots[plot]
			if !ok {
				// Out of bounds.
				break
			}

			if _, ok := visited[plot]; ok {
				// Already visited.
				y++
				continue
			}

			region := explore(plots, plot, crop, visited)
			regions = append(regions, region)

			y++
		}
		x++
	}

	cost := 0
	for _, region := range regions {
		cost += len(region.plots) * region.perimeter
	}
	return cost
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle))
}
