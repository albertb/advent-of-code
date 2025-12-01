package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

func parse(input string) []int {
	var numbers []int
	for line := range strings.SplitSeq(input, "\n") {
		sign := line[0]
		number, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			log.Fatalln("failed to parse", line, ":", err)
		}
		if sign == 'L' {
			number *= -1
		} else if sign != 'R' {
			log.Fatalln("unexpected sign:", sign)
		}
		numbers = append(numbers, int(number))
	}
	return numbers
}

func part1(input string) int {
	numbers := parse(input)

	count := 0
	position := 50
	for _, number := range numbers {
		position = (position + number + 1000) % 100
		if position == 0 {
			count++
		}
	}
	return count
}

func part2(input string) int {
	numbers := parse(input)

	count := 0
	position := 50
	for _, number := range numbers {
		right := true
		if number < 0 {
			right = false
			number *= -1
		}

		if number >= 100 {
			count += number / 100
		}
		number %= 100

		start := position
		if right {
			position += number
		} else {
			position -= number
		}

		if position > 100 || (start != 0 && position < 0) {
			count++
		}
		position %= 100

		if position == 0 {
			count++
		}
		if position < 0 {
			position += 100
		}
	}
	return count
}

//go:embed puzzle.txt
var puzzle string
