package main

import (
	_ "embed"
	"fmt"
	"strings"
)

func parse(input string) (patterns []string, desgins []string) {
	n := 0
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		if n == 0 {
			for pattern := range strings.SplitSeq(line, ",") {
				pattern = strings.TrimSpace(pattern)
				patterns = append(patterns, pattern)
			}
		} else {
			design := strings.TrimSpace(line)
			desgins = append(desgins, design)
		}
		n++
	}
	return patterns, desgins
}

func assemble(design string, patterns []string) bool {
	// Keeps track of substrings that are possible to assemble using the patterns.
	possible := make(map[string]struct{})
	for _, pattern := range patterns {
		possible[pattern] = struct{}{}
	}
	len := len(design)

	// Keeps track of the prefix lengths of design that can be assembled from the patterns.
	dp := make([]bool, len+1)
	dp[0] = true

	for end := 1; end <= len; end++ {
		for start := 0; start < end; start++ {
			// IF
			//   design[0:start] can be assembled
			// AND
			//   pats contains design[start:end]
			// THEN
			//   design[0:end] can be assembled
			if _, ok := possible[design[start:end]]; dp[start] && ok {
				dp[end] = true
				break
			}
		}
	}
	return dp[len]
}

func count(design string, patterns []string) int {
	possible := make(map[string]struct{})
	for _, pattern := range patterns {
		possible[pattern] = struct{}{}
	}
	len := len(design)

	// Keeps track of the prefix lengths of design that can be assembled from the patterns.
	dp := make([]int, len+1)
	dp[0] = 1

	for end := 1; end <= len; end++ {
		for start := 0; start < end; start++ {
			// IF
			//   design[0:start] can be assembled
			// AND
			//   pats contains design[start:end]
			// THEN
			//   design[0:end] can be assembled as many different ways as design[0:start]
			if _, ok := possible[design[start:end]]; dp[start] > 0 && ok {
				dp[end] += dp[start]
			}
		}
	}
	return dp[len]
}

func part1(input string) int {
	patterns, designs := parse(input)

	n := 0
	for _, design := range designs {
		if assemble(design, patterns) {
			n++
		}
	}
	return n
}

func part2(input string) int {
	patterns, designs := parse(input)

	n := 0
	for _, design := range designs {
		n += count(design, patterns)
	}
	return n
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle))
	fmt.Println("2:", part2(puzzle))
}
