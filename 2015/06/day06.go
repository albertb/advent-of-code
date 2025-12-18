package main

import (
	"fmt"
	"image"
	"strings"
)

func part1(input string) int {
	lights := [1000][1000]bool{}
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		fields := strings.Fields(line)
		var op int // -1 turn off, 0 toggle, 1 turn on
		var first, last string
		if len(fields) == 4 {
			first, last = fields[1], fields[3]
		} else {
			if fields[1] == "on" {
				op = 1
			} else {
				op = -1
			}
			first, last = fields[2], fields[4]
		}

		var r image.Rectangle
		fmt.Sscanf(first, "%d,%d", &r.Min.X, &r.Min.Y)
		fmt.Sscanf(last, "%d,%d", &r.Max.X, &r.Max.Y)
		r = r.Canon()

		for x := r.Min.X; x <= r.Max.X; x++ {
			for y := r.Min.Y; y <= r.Max.Y; y++ {
				switch op {
				case -1:
					lights[x][y] = false
				case 0:
					lights[x][y] = !lights[x][y]
				case 1:
					lights[x][y] = true
				}
			}
		}
	}

	var lit int
	for y := range len(lights) {
		for x := range len(lights[y]) {
			if lights[x][y] {
				lit++
			}
		}
	}
	return lit
}

func part2(input string) int {
	lights := [1000][1000]int{}

	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		fields := strings.Fields(line)
		var op int // -1 turn off, 0 toggle, 1 turn on
		var first, last string
		if len(fields) == 4 {
			first, last = fields[1], fields[3]
		} else {
			if fields[1] == "on" {
				op = 1
			} else {
				op = -1
			}
			first, last = fields[2], fields[4]
		}

		var r image.Rectangle
		fmt.Sscanf(first, "%d,%d", &r.Min.X, &r.Min.Y)
		fmt.Sscanf(last, "%d,%d", &r.Max.X, &r.Max.Y)
		r = r.Canon()

		for x := r.Min.X; x <= r.Max.X; x++ {
			for y := r.Min.Y; y <= r.Max.Y; y++ {
				switch op {
				case -1:
					lights[y][x] = max(0, lights[y][x]-1)
				case 0:
					lights[y][x] += 2
				case 1:
					lights[y][x]++
				}
			}
		}
	}

	var brightness int
	for y := range len(lights) {
		for x := range len(lights[y]) {
			brightness += lights[y][x]
		}
	}
	return brightness
}
