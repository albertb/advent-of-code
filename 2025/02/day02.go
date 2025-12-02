package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Range struct {
	first, last int
}

func parse(input string) []Range {
	var ranges []Range
	for idRange := range strings.SplitSeq(input, ",") {
		ids := strings.Split(idRange, "-")
		if len(ids) != 2 {
			log.Fatalln("unexpected range:", ids)
		}
		first, err := strconv.ParseInt(ids[0], 10, 64)
		if err != nil {
			log.Fatalln("failed to parse ID:", ids[0], err)
		}
		last, err := strconv.ParseInt(ids[1], 10, 64)
		if err != nil {
			log.Fatalln("failed to parse ID:", ids[1], err)
		}
		ranges = append(ranges, Range{int(first), int(last)})
	}
	return ranges
}

func part1(input string) int {
	ranges := parse(input)

	sum := 0
	for _, r := range ranges {
		for i := r.first; i <= r.last; i++ {
			number := fmt.Sprint(i)
			if number[:len(number)/2] == number[len(number)/2:] {
				sum += i
			}
		}
	}
	return sum
}

func hasRepeatingSegments(number string, length int) bool {
	last := number[:length]
	for segment := 1; segment < len(number)/length; segment++ {
		next := number[segment*length : (segment+1)*length]
		if next != last {
			return false
		}
	}
	return true
}

func part2(input string) int {
	ranges := parse(input)

	sum := 0
	for _, r := range ranges {
		for i := r.first; i <= r.last; i++ {
			number := fmt.Sprint(i)
			numLength := len(number)
			for segLength := numLength / 2; segLength > 0; segLength-- {
				if numLength%segLength != 0 {
					// Skip this segment length if it doesn't evenly split the number.
					continue
				}

				if hasRepeatingSegments(number, segLength) {
					sum += i
					break
				}
			}
		}
	}
	return sum
}

//go:embed puzzle.txt
var puzzle string
