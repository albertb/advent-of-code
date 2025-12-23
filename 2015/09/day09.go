package main

import (
	"fmt"
	"math"
	"strings"
)

func parse(input string) map[string]map[string]int {
	cities := map[string]map[string]int{}
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		fields := strings.Fields(line)
		a := fields[0]
		b := fields[2]
		var distance int
		fmt.Sscanf(fields[4], "%d", &distance)

		if _, ok := cities[a]; !ok {
			cities[a] = make(map[string]int)
		}
		cities[a][b] = distance
		if _, ok := cities[b]; !ok {
			cities[b] = make(map[string]int)
		}
		cities[b][a] = distance
	}
	return cities
}

func visit(cities map[string]map[string]int, visited map[string]struct{}, city string, minimize bool) int {
	if len(visited) == len(cities) {
		return 0 // Base case: we visited all the cities.
	}

	best := math.MaxInt
	fn := func(a, b int) int {
		if minimize {
			return min(a, b)
		}
		return max(a, b)
	}
	if !minimize {
		best = 0
	}

	if len(city) < 1 {
		for city = range cities {
			visited[city] = struct{}{}
			best = fn(best, visit(cities, visited, city, minimize))
			delete(visited, city)
		}
	} else {
		for neighbour, dist := range cities[city] {
			if _, ok := visited[neighbour]; ok {
				continue
			}
			visited[neighbour] = struct{}{}
			best = fn(best, dist+visit(cities, visited, neighbour, minimize))
			delete(visited, neighbour)
		}
	}
	return best
}

func part1(input string) int {
	cities := parse(input)
	return visit(cities, make(map[string]struct{}), "", true)
}

func part2(input string) int {
	cities := parse(input)
	return visit(cities, make(map[string]struct{}), "", false)
}
