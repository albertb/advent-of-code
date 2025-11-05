package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

type Equation struct {
	value   int64
	numbers []int64
}

func evaluate(numbers []int64, concat bool) []int64 {
	if len(numbers) == 1 {
		return []int64{numbers[0]}
	}

	lhs := numbers[0]
	rhss := evaluate(numbers[1:], concat)

	op := func(a, b int64) int64 {
		s := fmt.Sprintf("%d%d", a, b)
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		return n
	}

	var result []int64
	for _, rhs := range rhss {
		result = append(result, rhs+lhs)
		result = append(result, rhs*lhs)
		if concat {
			result = append(result, op(rhs, lhs))
		}
	}

	return result
}

func (e Equation) Valid(concat bool) bool {
	numbers := slices.Clone(e.numbers)
	slices.Reverse(numbers)

	for _, value := range evaluate(numbers, concat) {
		if value == e.value {
			return true
		}
	}
	return false
}

func parse(input string) []Equation {
	var result []Equation
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		var eq Equation

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Fatalln("invalid line:", line)
		}
		value, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		eq.value = value

		parts = strings.Split(parts[1], " ")
		if len(parts) < 1 {
			log.Fatalln("invalid numbers:", err)
		}

		for _, part := range parts {
			if len(part) == 0 {
				continue
			}
			number, err := strconv.ParseInt(part, 10, 64)
			if err != nil {
				log.Fatalln(err)
			}
			eq.numbers = append(eq.numbers, number)
		}
		result = append(result, eq)
	}
	return result
}

func part1(input string) int64 {
	equations := parse(input)

	sum := int64(0)
	for _, e := range equations {
		if e.Valid(false) {
			sum += e.value
		}
	}
	return sum
}

func part2(input string) int64 {
	equations := parse(input)

	sum := int64(0)
	for _, e := range equations {
		if e.Valid(true) {
			sum += e.value
		}
	}
	return sum
}
