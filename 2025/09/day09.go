package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"
)

func parse(input string) []image.Point {
	var points []image.Point
	for _, line := range strings.Fields(input) {
		var p image.Point
		fmt.Sscanf(line, "%d,%d", &p.X, &p.Y)
		points = append(points, p)
	}
	return points
}

func part1(input string) int {
	points := parse(input)

	largest := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			rect := image.Rectangle{points[i], points[j]}.Canon()
			// Stretch the rectangle to fill the row/col it ends on.
			rect.Max = rect.Max.Add(image.Point{1, 1})

			area := rect.Dx() * rect.Dy()
			largest = max(largest, area)
		}
	}
	return largest
}

func part2(input string) int {
	points := parse(input)

	var lines []image.Rectangle
	p := points[len(points)-1]
	for i := range points {
		q := points[i]
		line := image.Rectangle{p, q}.Canon()
		// Stretch the line to fill the row/col it ends on.
		line.Max = line.Max.Add(image.Point{1, 1})
		lines = append(lines, line)
		p = q
	}

	var rects []image.Rectangle
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			rect := image.Rectangle{points[i], points[j]}.Canon()
			// Stretch the rectangle to fill the row/col it ends on.
			rect.Max = rect.Max.Add(image.Point{1, 1})
			rects = append(rects, rect)
		}
	}

	largest := 0
	for _, r := range rects {
		contained := true
		for _, l := range lines {
			// Check whether the line intersects with the rectangle, but without its extension into the
			// row/col it ends on.
			if l.Overlaps(r.Inset(1)) {
				contained = false
				break
			}
		}
		if contained {
			area := r.Dx() * r.Dy()
			largest = max(largest, area)
		}
	}

	return largest
}

//go:embed puzzle.txt
var puzzle string
