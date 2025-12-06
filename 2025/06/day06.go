package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

func part1(input string) int {
	operands := [][]int{}
	operators := []bool{} // Whether the operator is *

	j := 0
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			continue
		}
		operands = append(operands, []int{})
		for symbol := range strings.FieldsSeq(line) {
			switch symbol {
			case "+":
				operators = append(operators, false)
			case "*":
				operators = append(operators, true)
			default:
				number, err := strconv.ParseInt(symbol, 10, 64)
				if err != nil {
					log.Fatalln("failed to parse number:", symbol, err)
				}
				operands[j] = append(operands[j], int(number))
			}
		}
		j++
	}

	total := 0
	for i := 0; i < len(operators); i++ {
		result := operands[0][i]
		for j := 1; j < len(operands[j]); j++ {
			if operators[i] {
				result *= operands[j][i]
			} else {
				result += operands[j][i]
			}
		}
		total += result
	}
	return total
}

func rotate(s string) string {
	lines := strings.Split(s, "\n")

	// Add an empty line before the operators so we get a space after the rotation.
	lines = append(lines, lines[len(lines)-1])
	lines[len(lines)-2] = ""

	rows := len(lines)
	cols := 0
	for _, line := range lines {
		cols = mathy.Max(cols, len(line))
	}

	rotated := make([]string, cols)

	for i := 0; i < rows; i++ {
		line := lines[i]
		for j := 0; j < cols; j++ {
			if j < len(line) {
				rotated[j] += string(line[j])
			} else {
				rotated[j] += " "
			}
		}
	}

	return strings.Join(rotated, "\n")
}

func part2(input string) int {
	input = rotate(input)

	total := 0
	result := 0
	mult := false

	for line := range strings.SplitSeq(input, "\n") {
		fields := strings.Fields(line)
		if len(fields) < 1 {
			// Blank line.
			continue
		}
		if len(fields) > 1 {
			total += result

			switch fields[1] {
			case "*":
				result = 1
				mult = true
			case "+":
				result = 0
				mult = false
			default:
				log.Fatalln("unexpected symbol:", fields[1])
			}
		}
		number, err := strconv.ParseInt(fields[0], 10, 64)
		if err != nil {
			log.Fatalln("failed to parse number:", fields[0], err)
		}
		if mult {
			result *= int(number)
		} else {
			result += int(number)
		}
	}
	total += result

	return total
}

//go:embed puzzle.txt
var puzzle string
