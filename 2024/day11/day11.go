package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle, 25))
	fmt.Println("2:", part2(puzzle, 75))
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

var powerOf10 = []int64{
	1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000,
	10000000000, 100000000000, 1000000000000, 10000000000000, 100000000000000,
	1000000000000000, 10000000000000000, 100000000000000000, 1000000000000000000,
}

func blink(before []int64) []int64 {
	after := make([]int64, len(before)*2)

	index := 0
	for _, stone := range before {
		if stone == 0 {
			after[index] = 1
			index++
			continue
		}

		digits := int(math.Log10(float64(stone))) + 1

		if digits > 1 && digits%2 == 0 {
			splitPower := powerOf10[digits/2]

			lhs := stone / splitPower
			after[index] = lhs

			rhs := stone % splitPower
			after[index+1] = rhs

			index += 2
			continue
		}

		after[index] = stone * 2024
		index++
	}
	return after[:index]
}

func singleBlink(stone int64) (int64, int64) {
	if stone == 0 {
		return 1, -1
	}

	digits := int(math.Log10(float64(stone))) + 1

	if digits > 1 && digits%2 == 0 {
		splitPower := powerOf10[digits/2]

		left := stone / splitPower
		right := stone % splitPower

		return left, right
	}

	next := stone * 2024
	return next, -1
}

type CacheKey struct {
	stone int64
	depth int
}

func countStones(stone int64, depth int, countCache *map[CacheKey]int) int {
	key := CacheKey{stone, depth}

	if cached, ok := (*countCache)[key]; ok {
		return cached
	}

	left, right := singleBlink(stone)

	if depth == 1 {
		if right >= 0 {
			(*countCache)[key] = 2
			return 2
		}
		(*countCache)[key] = 1
		return 1
	}

	count := countStones(left, depth-1, countCache)
	if right >= 0 {
		count += countStones(right, depth-1, countCache)
	}
	(*countCache)[key] = count
	return count
}

func part1(input string, blinks int) int {
	stones := parse(input)

	for range blinks {
		stones = blink(stones)
	}
	return len(stones)
}

func part2(input string, blinks int) int {
	stones := parse(input)

	sum := 0

	cache := make(map[CacheKey]int)

	for _, stone := range stones {
		sum += countStones(stone, blinks, &cache)
	}
	return sum
}
