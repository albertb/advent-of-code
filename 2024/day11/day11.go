package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle, 25))
	fmt.Println("2:", part1(puzzle, 75))
}

func parse(input string) []int64 {
	numbers := []int64{}
	for s := range strings.SplitSeq(input, " ") {
		number, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func blink(before []int64) []int64 {
	var after []int64
	for _, number := range before {
		if number == 0 {
			after = append(after, 1)
			continue
		}

		mustParseInt := func(s string) int64 {
			number, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatalln(err)
			}
			return number
		}

		s := strconv.FormatInt(number, 10)
		len := len(s)
		if len%2 == 0 {
			lhs := mustParseInt(s[:len/2])
			after = append(after, lhs)
			rhs := mustParseInt(s[len/2:])
			after = append(after, rhs)
			continue
		}

		after = append(after, number*2024)
	}
	return after
}

func part1(input string, blinks int) int {
	numbers := parse(input)

	for i := range blinks {
		fmt.Println("blink", i+1, "of", blinks)
		numbers = blink(numbers)
	}
	return len(numbers)
}
