package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

type Bank struct {
	batteries []int
}

func parse(input string) []Bank {
	var banks []Bank
	for line := range strings.SplitSeq(input, "\n") {
		var batteries []int
		for _, label := range strings.Split(line, "") {
			joltage, err := strconv.ParseInt(label, 10, 64)
			if err != nil {
				log.Fatalln(err)
			}
			batteries = append(batteries, int(joltage))
		}
		banks = append(banks, Bank{batteries})
	}
	return banks
}

func part1(input string) int {
	banks := parse(input)

	output := 0
	for _, bank := range banks {
		max := 0
		for i := 0; i < len(bank.batteries)-1; i++ {
			for j := i + 1; j < len(bank.batteries); j++ {
				joltage := bank.batteries[i]*10 + bank.batteries[j]
				if joltage > max {
					max = joltage
				}
			}
		}
		output += max
	}
	return output
}

func part2(input string) int {
	banks := parse(input)

	output := 0
	for _, bank := range banks {
		batts := bank.batteries
		next := 0

		joltage := 0
		for n := 1; n <= 12; n++ {
			max := -1
			for i := next; i < len(batts)-(12-n); i++ {
				if batts[i] > max {
					max = batts[i]
					next = i + 1
				}
			}
			joltage += max * int(mathy.Pow(10, 12-n))
		}
		output += joltage
	}
	return output
}

//go:embed puzzle.txt
var puzzle string
