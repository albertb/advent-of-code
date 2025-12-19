package main

import (
	"log"
	"strconv"
	"strings"
)

func eval(id string, exprs map[string][]string, signals map[string]int) int {
	if strings.ContainsAny(id, "0123456789") {
		// It's an integer, parse it, and return it's value.
		val, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln("failed to parse value", id, err)
		}
		return val
	}

	if signal, ok := signals[id]; ok {
		// We've already evaluated this signal, return its value.
		return signal
	}

	expr, ok := exprs[id]
	if !ok {
		log.Fatalln("cannot find expr for", id)
	}

	var result int
	switch len(expr) {
	case 1:
		// Either another signal, or an integer.
		result = eval(expr[0], exprs, signals)
	case 2:
		// NOT x
		if expr[0] != "NOT" {
			log.Fatalln("expected NOT expr", expr)
		}
		rhs := eval(expr[1], exprs, signals)
		result = ^rhs
	case 3:
		// x AND/OR/LSHIFT/RSHIFT y
		lhs := eval(expr[0], exprs, signals)
		rhs := eval(expr[2], exprs, signals)
		switch expr[1] {
		case "AND":
			result = lhs & rhs
		case "OR":
			result = lhs | rhs
		case "LSHIFT":
			result = lhs << rhs
		case "RSHIFT":
			result = lhs >> rhs
		default:
			log.Fatalln("unexpected operator", expr)
		}
	default:
		log.Fatalln("unexpected expr", expr)
	}
	signals[id] = result
	return result
}

func parse(input string) map[string][]string {
	exprs := map[string][]string{}
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		parts := strings.Split(line, "->")
		output := strings.TrimSpace(parts[1])
		exprs[output] = strings.Fields(parts[0])
	}
	return exprs
}

func part1(input string) int {
	exprs := parse(input)
	return eval("a", exprs, make(map[string]int))
}

func part2(input string) int {
	exprs := parse(input)
	a := eval("a", exprs, make(map[string]int))
	return eval("a", exprs, map[string]int{"b": a})
}
