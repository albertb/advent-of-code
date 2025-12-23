package main

import (
	"fmt"
	"math"
	"strings"
)

func visit2(cities map[string]map[string]int, visited map[string]struct{}, city string) int {
	if len(visited) == len(cities) {
		return 0 // Base case: we visited all the cities.
	}

	shortest := math.MaxInt
	if len(city) < 1 {
		for city = range cities {
			visited[city] = struct{}{}
			shortest = min(shortest, visit2(cities, visited, city))
			delete(visited, city)
		}
	} else {
		for neighbour, distance := range cities[city] {
			if _, ok := visited[neighbour]; ok {
				continue
			}
			visited[neighbour] = struct{}{}
			shortest = min(shortest, distance+visit2(cities, visited, neighbour))
			delete(visited, neighbour)
		}
	}
	return shortest
}

func part1(input string) int {
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

	return visit2(cities, make(map[string]struct{}), "")
}
