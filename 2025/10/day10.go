package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/draffensperger/golp"
)

type Machine struct {
	diagram []bool
	buttons [][]int
	joltage []int
}

func (m Machine) Initial() []bool {
	return make([]bool, len(m.diagram))
}

func (m Machine) Ready(lights []bool) bool {
	for i := range lights {
		if lights[i] != m.diagram[i] {
			return false
		}
	}
	return true
}

func (m Machine) Press(lights []bool, button int) []bool {
	for _, toggle := range m.buttons[button] {
		lights[toggle] = !lights[toggle]
	}
	return lights
}

func parse(input string) []Machine {
	var machines []Machine
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		var machine Machine
		parts := strings.Fields(line)

		// Parse the indicator lights diagram, e.g., [..##.]
		machine.diagram = make([]bool, len(parts[0])-2)
		for i := 1; i < len(parts[0])-1; i++ {
			machine.diagram[i-1] = parts[0][i] == '#'
		}

		// Parse the buttons, e.g., (0) (1,2,3) (0,2)
		for i := 1; i < len(parts)-1; i++ {
			part := parts[i]
			s := part[1 : len(part)-1]
			var button []int
			for _, id := range strings.Split(s, ",") {
				var toggle int
				fmt.Sscanf(id, "%d", &toggle)
				button = append(button, toggle)
			}
			machine.buttons = append(machine.buttons, button)
		}

		// Parse the joltage, eg., {1,2,3}
		part := parts[len(parts)-1]
		s := part[1 : len(part)-1]
		for _, val := range strings.Split(s, ",") {
			var joltage int
			fmt.Sscanf(val, "%d", &joltage)
			machine.joltage = append(machine.joltage, joltage)
		}
		machines = append(machines, machine)
	}
	return machines
}

func part1(input string) int {
	machines := parse(input)

	total := 0
	for _, m := range machines {
		fewest := math.MaxInt

		// For each machine, for each button, try every combination of press and skip.
		n := len(m.buttons)
		combinations := 1 << n // 2^n possible combinations
		for i := range combinations {
			lights := m.Initial()

			press := 0
			for j := 0; j < n; j++ {
				if (i & (1 << j)) != 0 {
					lights = m.Press(lights, j)
					press++
				}
			}
			if m.Ready(lights) {
				fewest = min(fewest, press)
			}
		}
		total += fewest
	}

	return total
}

// TODO: Get rid of LPSolve, and rewrite this by hand.
func part2(input string) int {
	machines := parse(input)

	sum := 0
	for _, m := range machines {
		//fmt.Printf("Machine %d of %d\n", i, len(machines))

		lp := golp.NewLP(len(m.joltage), len(m.buttons))

		// Add names to the button variables for debugging.
		for i := range len(m.buttons) {
			lp.SetColName(i, fmt.Sprintf("b_%d", i))
		}

		// Minimize the sum of all the variables, i.e., button presses.
		obj := make([]float64, len(m.buttons))
		for i := range len(m.buttons) {
			obj[i] = 1.0
		}
		lp.SetObjFn(obj)

		// Each row represents the buttons that can affect that joltage.
		//
		// E.g., (0) (1) (0,1) {1,2} would become:
		// coeffs: [1.0, 0.0, 1.0]; constraint: =1
		// coeffs: [0.0, 1.0, 1.0]; constraint: =2
		for i, jolt := range m.joltage {
			coeffs := make([]float64, len(m.buttons))
			for j, btn := range m.buttons {
				if slices.Contains(btn, i) {
					coeffs[j] = 1.0
				}
			}
			lp.AddConstraint(coeffs, golp.EQ, float64(jolt))
		}

		// We want integer solutions only.
		for i := range m.buttons {
			lp.SetInt(i, true)
		}

		status := lp.Solve()
		if status != golp.OPTIMAL {
			fmt.Println("failed to solve machine", status)
			return 0
		}

		// Add up the number of button presses from the solution.
		for _, v := range lp.Variables() {
			sum += int(v)
		}
	}

	return sum
}

//go:embed puzzle.txt
var puzzle string
