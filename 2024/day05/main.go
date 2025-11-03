package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

type Rule struct {
	lhs, rhs int
}

func ParseRule(line string) (Rule, error) {
	var result Rule
	parts := strings.Split(line, "|")
	if len(parts) != 2 {
		return result, fmt.Errorf("invalid rule format: %s", line)
	}
	lhs, err := strconv.ParseInt(parts[0], 10, 16)
	if err != nil {
		return result, fmt.Errorf("invalid rule lhs for %s: %w", line, err)
	}
	rhs, err := strconv.ParseInt(parts[1], 10, 16)
	if err != nil {
		return result, fmt.Errorf("invalid rule rhs for %s: %w", line, err)
	}
	if lhs == rhs {
		return result, fmt.Errorf("invalid rule: %s", line)
	}
	return Rule{
		lhs: int(lhs),
		rhs: int(rhs),
	}, nil
}

func (r Rule) Violated(update []int) bool {
	seen := false
	for _, page := range update {
		if page == r.lhs {
			if seen {
				return true
			}
		}
		if page == r.rhs {
			seen = true
		}
	}
	return false
}

func (r Rule) Apply(update []int) []int {
	violation := -1
	for i := range update {
		if update[i] == r.rhs {
			violation = i
		}
		if update[i] == r.lhs {
			if violation >= 0 {
				swapped := update[violation]
				update[violation] = update[i]
				update[i] = swapped
			}
		}
	}
	return update
}

func ParseUpdate(line string) ([]int, error) {
	var result []int
	parts := strings.Split(line, ",")
	for _, part := range parts {
		page, err := strconv.ParseInt(part, 10, 16)
		if err != nil {
			return result, fmt.Errorf("invalid update for %s: %w", line, err)
		}
		result = append(result, int(page))
	}
	return result, nil
}

func parse(input string) ([]Rule, [][]int) {
	rules := []Rule{}
	updates := [][]int{}

	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "|") {
			rule, err := ParseRule(line)
			if err != nil {
				log.Fatalln(err)
			}
			rules = append(rules, rule)
		} else if strings.Contains(line, ",") {
			update, err := ParseUpdate(line)
			if err != nil {
				log.Fatalln(err)
			}
			updates = append(updates, update)
		}
	}

	return rules, updates
}

func part1(input string) int {
	rules, updates := parse(input)

	count := 0
	for _, update := range updates {
		violated := false
		for _, rule := range rules {
			if rule.Violated(update) {
				//fmt.Println("Rule", rule, "violated in", update)
				violated = true
				break
			}
		}
		if !violated {
			middle := len(update) / 2
			count += update[middle]
		}
	}
	return count
}

func part2(input string) int {
	rules, updates := parse(input)

	count := 0
	for _, update := range updates {
		reordered := false
		for i := 0; i < len(rules); i++ {
			if rules[i].Violated(update) {
				//fmt.Println("VIOLATION", rules[i], "in", update)
				update = rules[i].Apply(update)
				i = 0
				reordered = true
				//fmt.Println("REORDERED", update)
			}
		}
		if reordered {
			middle := len(update) / 2
			count += update[middle]
		}
	}
	return count
}
