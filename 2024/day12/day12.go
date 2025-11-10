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

func getRegions(plots map[Plot]Crop) []Region {
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
	return regions
}

func part1(input string) int {
	plots := parse(input)
	regions := getRegions(plots)

	cost := 0
	for _, region := range regions {
		cost += len(region.plots) * region.perimeter
	}
	return cost
}

type Direction int

const (
	Up Direction = iota
	UpRight
	Right
	RightDown
	Down
	LeftDown
	Left
	LeftUp
)

func isOob(plot Plot, dir Direction, plots map[Plot]struct{}) bool {
	var ok bool
	switch dir {
	case Up:
		_, ok = plots[Plot{plot.x, plot.y - 1}]
	case UpRight:
		_, ok = plots[Plot{plot.x + 1, plot.y - 1}]
	case Right:
		_, ok = plots[Plot{plot.x + 1, plot.y}]
	case RightDown:
		_, ok = plots[Plot{plot.x + 1, plot.y + 1}]
	case Down:
		_, ok = plots[Plot{plot.x, plot.y + 1}]
	case LeftDown:
		_, ok = plots[Plot{plot.x - 1, plot.y + 1}]
	case Left:
		_, ok = plots[Plot{plot.x - 1, plot.y}]
	case LeftUp:
		_, ok = plots[Plot{plot.x - 1, plot.y - 1}]
	}
	// If the map was empty for the given plot, we're out of bounds.
	return !ok
}

func countCorners(plots map[Plot]struct{}) int {
	corners := 0
	for plot := range plots {
		oobs := map[Direction]bool{}
		for _, dir := range []Direction{Up, UpRight, Right, RightDown, Down, LeftDown, Left, LeftUp, Up} {
			oobs[dir] = isOob(plot, dir, plots)
		}

		// Outer corners.
		if oobs[Up] && oobs[Right] {
			corners++
		}
		if oobs[Right] && oobs[Down] {
			corners++
		}
		if oobs[Down] && oobs[Left] {
			corners++
		}
		if oobs[Left] && oobs[Up] {
			corners++
		}

		// Inner corners.
		if !oobs[Up] && !oobs[Right] && oobs[UpRight] {
			corners++
		}
		if !oobs[Right] && !oobs[Down] && oobs[RightDown] {
			corners++
		}
		if !oobs[Down] && !oobs[Left] && oobs[LeftDown] {
			corners++
		}
		if !oobs[Left] && !oobs[Up] && oobs[LeftUp] {
			corners++
		}
	}
	return corners
}

func part2(input string) int {
	plots := parse(input)
	regions := getRegions(plots)

	sum := 0
	for _, region := range regions {
		plots := map[Plot]struct{}{}
		for _, plot := range region.plots {
			plots[plot] = struct{}{}
		}
		corners := countCorners(plots)
		area := len(region.plots)
		sum += corners * area
	}
	return sum
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle))
	fmt.Println("2:", part2(puzzle))
}
