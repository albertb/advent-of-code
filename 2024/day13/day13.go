package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Vec struct {
	x, y int
}

type Machine struct {
	a    Vec
	b    Vec
	goal Vec
}

func parse(input string) []Machine {
	var machines []Machine

	var re = regexp.MustCompile(
		`Button ([A-Z]): X\+(\d+), Y\+(\d+)|Prize: X=(\d+), Y=(\d+)`)

	var machine Machine
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			continue
		}

		match := re.FindStringSubmatch(line)
		if len(match[1]) > 0 {
			var button *Vec
			label := match[1]
			if label == "A" {
				button = &machine.a
			} else {
				button = &machine.b
			}

			x, err := strconv.ParseInt(match[2], 10, 16)
			if err != nil {
				log.Fatalln(err)
			}
			y, err := strconv.ParseInt(match[3], 10, 16)
			if err != nil {
				log.Fatalln(err)
			}

			button.x, button.y = int(x), int(y)
		}
		if len(match[4]) > 0 {
			x, err := strconv.ParseInt(match[4], 10, 16)
			if err != nil {
				log.Fatalln(err)
			}
			y, err := strconv.ParseInt(match[5], 10, 16)
			if err != nil {
				log.Fatalln(err)
			}

			machine.goal.x, machine.goal.y = int(x), int(y)

			// Prize is the last description line, append this machine and start a new one.
			machines = append(machines, machine)
			machine = Machine{}
		}
	}
	return machines
}

func minimize(machine Machine, max int) int {
	goal := machine.goal

	minCost := math.MaxInt
	for a := range max {
		for b := range max {
			vec := Vec{
				a*machine.a.x + b*machine.b.x,
				a*machine.a.y + b*machine.b.y,
			}
			if vec.x > goal.x || vec.y > goal.y {
				// We went too far, give up on this branch.
				break
			}
			if a*3+b > minCost {
				// Too expensive, give up on this branch.
				break
			}
			if vec.x == goal.x && vec.y == goal.y {
				cost := a*3 + b
				if cost < minCost {
					minCost = cost
				}
				continue
			}
		}
	}

	return minCost
}

func part1(input string) int {
	machines := parse(input)

	var sum int
	for _, machine := range machines {
		cost := minimize(machine, 100)
		if cost < math.MaxInt {
			sum += cost
		}
	}
	return sum
}

func equations(machine Machine) int {
	ax, ay := machine.a.x, machine.a.y
	bx, by := machine.b.x, machine.b.y
	px, py := machine.goal.x, machine.goal.y

	b := (py*ax - px*ay)
	div := (by*ax - bx*ay)
	if b%div != 0 {
		return math.MaxInt
	}
	b = b / div

	a := (px - b*bx)
	if a%ax != 0 {
		return math.MaxInt
	}
	a = a / ax

	if a < 0 || b < 0 {
		return math.MaxInt
	}

	fmt.Println("machine", machine, "a", a, "b", b)

	return a*3 + b
}

func part2(input string) int {
	machines := parse(input)

	for i := range machines {
		machines[i].goal.x += 10000000000000
		machines[i].goal.y += 10000000000000
	}

	var sum int
	for _, machine := range machines {
		cost := equations(machine)
		if cost < math.MaxInt {
			sum += cost
		}
	}
	return sum
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle))
	fmt.Println("2:", part2(puzzle))
}
