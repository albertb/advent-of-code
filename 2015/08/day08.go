package main

import (
	"log"
	"strconv"
	"strings"
)

func part1(input string) int {
	var count int
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		s, err := strconv.Unquote(line)
		if err != nil {
			log.Fatalln(line, ":", err)
		}
		count += (len(line) - len(s))
	}
	return count
}

func part2(input string) int {
	var count int
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		s := strconv.Quote(line)
		count += (len(s) - len(line))
	}
	return count
}
