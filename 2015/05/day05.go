package main

import (
	"strings"
)

func part1(input string) int {
	var good int
nextword:
	for line := range strings.SplitSeq(input, "\n") {
		var vowels, repeated int
		var last byte
		for i := range line {
			// Look for vowels.
			if strings.ContainsAny(line[i:i+1], "aeiou") {
				vowels++
			}

			// Look for repeated letters.
			curr := line[i]
			if curr == last {
				repeated++
			}

			if i > 0 {
				s := line[i-1 : i+1]
				// If any of the bad strings appear, give up, go to the next word.
				if s == "ab" || s == "cd" || s == "pq" || s == "xy" {
					continue nextword
				}
			}
			last = curr
		}
		// At least 3 vowels, and 1 repeated letter.
		if vowels >= 3 && repeated > 0 {
			good++
		}
	}
	return good
}

func part2(input string) int {
	var good int
nextword:
	for line := range strings.SplitSeq(input, "\n") {
		var repeatedPairs, repeatedLetters bool
		pairs := map[string]int{}
		for i := range line {
			if i > 0 {
				s := line[i-1 : i+1]
				if j, ok := pairs[s]; ok {
					if j < i-1 {
						repeatedPairs = true
					}
				} else {
					pairs[s] = i
				}
			}
			if i > 1 {
				if line[i-2] == line[i] {
					repeatedLetters = true
				}
			}
			if repeatedPairs && repeatedLetters {
				good++
				continue nextword
			}
		}
	}
	return good
}
