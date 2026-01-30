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
