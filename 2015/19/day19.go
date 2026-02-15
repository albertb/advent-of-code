package main

import (
	"strings"
)

func parse(input string) (map[string][]string, string) {
	replacements := make(map[string][]string)
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		if strings.Contains(line, "=>") {
			fields := strings.Fields(line)
			replacements[fields[0]] = append(replacements[fields[0]], fields[2])
		} else {
			return replacements, line
		}
	}
	return replacements, ""
}

func part1(input string) int {
	replacements, molecule := parse(input)

	generated := make(map[string]struct{})
	for from, tos := range replacements {
		for i := range molecule {
			if i+len(from) > len(molecule) {
				break
			}
			if molecule[i:i+len(from)] == from {
				for _, to := range tos {
					result := molecule[:i] + to + molecule[i+len(from):]
					generated[result] = struct{}{}
				}
			}
		}
	}

	return len(generated)
}

func part2(input string) int {
	reps, target := parse(input)

	curr := target
	steps := 0

	// BFS is too slow for part2. Instead, we do a greedy search, starting from the target and shortening
	// the string until we reach just 'e'.
outer:
	for curr != "e" {
		changed := false
		for from, tos := range reps {
			for _, to := range tos {
				if strings.Contains(curr, to) {
					curr = strings.Replace(curr, to, from, 1)
					changed = true
					steps++
					continue outer
				}
			}
		}
		if !changed {
			// If we get stuck, start over, map iteration order is randomized, so we'll
			// eventually get unstuck.
			curr = target
			steps = 0
		}
	}
	return steps
}
