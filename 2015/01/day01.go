package main

func part1(input string) int {
	floor := 0
	for _, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor
}

func part2(input string) int {
	floor := 0
	for i, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return -1
}
