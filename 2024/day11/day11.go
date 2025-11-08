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
	fmt.Println("X:", part2(puzzle, 46))
	//fmt.Println("2:", part2(puzzle, 75))
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

		lhs := stone / splitPower
		rhs := stone % splitPower

		return lhs, rhs
	}

	return stone * 2024, -1
}

type CacheEntry struct {
	left, right int64
}

func countStones(stone int64, depth int, cache *map[int64]CacheEntry) int {
	var left, right int64
	if cached, ok := (*cache)[stone]; ok {
		left, right = cached.left, cached.right
	} else {
		left, right = singleBlink(stone)
	}

	if depth == 1 {
		if right >= 0 {
			return 2
		}
		return 1
	}

	count := countStones(left, depth-1, cache)
	if right >= 0 {
		count += countStones(right, depth-1, cache)
	}
	return count
}

func part1(input string, blinks int) int {
	stones := parse(input)

	for i := range blinks {
		fmt.Println(i + 1)
		stones = blink(stones)
	}
	return len(stones)
}

func part2(input string, blinks int) int {
	stones := parse(input)

	sum := 0
	cache := make(map[int64]CacheEntry)
	for i, stone := range stones {
		fmt.Println(i)
		sum += countStones(stone, blinks, &cache)
	}
	return sum
}
