package main

import (
	"fmt"
	"math"
	"strings"
)

func parse(capacities string) []int {
	var containers []int
	for line := range strings.Lines(capacities) {
		if len(line) < 1 {
			continue
		}
		var capacity int
		fmt.Sscanf(line, "%d", &capacity)
		containers = append(containers, capacity)
	}
	return containers
}

func part1(capacities string, quantity int) int {
	containers := parse(capacities)

	var store func(containers []int, used []bool, n int) int
	store = func(containers []int, used []bool, n int) int {
		if n == len(containers) {
			remaining := quantity
			for i, use := range used {
				if use {
					remaining -= containers[i]
				}
			}
			if remaining == 0 {
				return 1
			}
			return 0
		}

		used[n] = true
		ways := store(containers, used, n+1)
		used[n] = false
		ways += store(containers, used, n+1)

		return ways
	}

	return store(containers, make([]bool, len(containers)), 0)
}

func part2(capacities string, quantity int) int {
	containers := parse(capacities)

	count, ways := math.MaxInt, 0
	var store func(containers []int, used []bool, n int)
	store = func(containers []int, used []bool, n int) {
		if n == len(containers) {
			remaining, n := quantity, 0
			for i, use := range used {
				if use {
					remaining -= containers[i]
					n++
				}
			}
			if remaining == 0 {
				if n < count {
					count = n
					ways = 1
				} else if n == count {
					ways++
				}
			}
			return
		}

		used[n] = true
		store(containers, used, n+1)
		used[n] = false
		store(containers, used, n+1)
	}

	store(containers, make([]bool, len(containers)), 0)
	return ways
}
