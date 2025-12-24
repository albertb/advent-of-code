package main

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
)

var number = regexp.MustCompile(`(-)?\d+`)

func part1(input string) int {
	var sum int
	matches := number.FindAllString(input, -1)
	for _, match := range matches {
		n, _ := strconv.Atoi(match)
		sum += n
	}
	return sum
}

func visit(data any) int {
	if m, ok := data.(map[string]any); ok {
		var sum int
		for _, v := range m {
			if v, ok := v.(string); ok {
				if v == "red" {
					// This objet contains a "red", skip it.
					return 0
				}
			}
			if v, ok := v.(float64); ok {
				sum += int(v)
				continue
			}
			sum += visit(v)
		}
		return sum
	}
	if a, ok := data.([]any); ok {
		var sum int
		for _, v := range a {
			if v, ok := v.(float64); ok {
				sum += int(v)
				continue
			}

			sum += visit(v)
		}
		return sum
	}
	if v, ok := data.(float64); ok {
		return int(v)
	}
	// Neither a map, nor an array, nor a number, skip it.
	return 0
}

func part2(input string) int {
	var data any
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		log.Fatalln(err)
	}
	return visit(data)
}
