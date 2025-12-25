package main

import (
	"fmt"
	"slices"
	"strings"
)

func sit(happiness map[string]map[string]int, seated []string) int {
	if len(happiness) == len(seated) {
		// Everyone is seated, add up the happiness from the last and first neighbours.
		first := seated[0]
		last := seated[len(seated)-1]
		return happiness[first][last] + happiness[last][first]
	}

	var h int
	seats := len(seated)
	for name := range happiness {
		if slices.Contains(seated, name) {
			// This person is already seated.
			continue
		}
		var diff int
		if len(seated) > 0 {
			// Add the happiness diff from the previous neighbour.
			last := seated[len(seated)-1]
			diff += happiness[last][name] + happiness[name][last]
		}
		seated = append(seated, name)
		h = max(h, diff+sit(happiness, seated))
		seated = seated[:seats]
	}
	return h
}

func parse(input string) map[string]map[string]int {
	happiness := map[string]map[string]int{}
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		fields := strings.Fields(line)
		a, b := fields[0], strings.TrimSuffix(fields[10], ".")

		var h int
		fmt.Sscanf(fields[3], "%d", &h)
		if fields[2] == "lose" {
			h *= -1
		}

		if _, ok := happiness[a]; !ok {
			happiness[a] = make(map[string]int)
		}
		happiness[a][b] = h
	}
	return happiness
}

func part1(input string) int {
	happiness := parse(input)
	return sit(happiness, []string{})
}

func part2(input string) int {
	happiness := parse(input)
	happiness["You"] = make(map[string]int)
	for a := range happiness {
		happiness[a]["You"] = 0
		happiness["You"][a] = 0
	}
	return sit(happiness, []string{})
}
