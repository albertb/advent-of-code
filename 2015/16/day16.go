package main

import (
	"fmt"
	"strings"
)

func solve(input string) (int, int) {
	sues := []map[string]int{}
	for line := range strings.Lines(input) {
		if len(line) < 1 {
			continue
		}
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ":", "")

		sue := make(map[string]int)
		fields := strings.Fields(line)
		for _, i := range []int{2, 4, 6} {
			var n int
			fmt.Sscanf(fields[i+1], "%d", &n)
			sue[fields[i]] = n
		}
		sues = append(sues, sue)
	}

	gifter := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	var part1, part2 int
one:
	for i, sue := range sues {
		for k, want := range gifter {
			if got, ok := sue[k]; ok {
				if got != want {
					continue one
				}
			}
		}
		part1 = i + 1
	}

two:
	for i, sue := range sues {
		for k, want := range gifter {
			if got, ok := sue[k]; ok {
				switch k {
				case "cats", "trees":
					if got <= want {
						continue two
					}
				case "pomeranians", "goldfish":
					if got >= want {
						continue two
					}
				default:
					if got != want {
						continue two
					}
				}
			}
		}
		part2 = i + 1
	}
	return part1, part2
}
