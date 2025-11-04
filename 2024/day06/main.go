package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

type State struct {
	grid   [][]rune
	x, y   int
	dir    rune
	visits map[string]bool
}

func parse(input string) State {
	var result State
	j := 0
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}

		row := []rune{}
		for i, rune := range line {
			if rune == '#' {
				row = append(row, '#')
				continue
			}

			row = append(row, '.')
			if rune == '.' {
				continue
			}

			result.x = i
			result.y = j
			result.dir = rune
		}
		result.grid = append(result.grid, row)
		j++
	}
	result.visits = make(map[string]bool)
	return result
}

// Returns 1. whether the guard can still move, 2. whether the guard looped
func (s *State) Next() (bool, bool) {
	dx, dy := 0, 0
	switch s.dir {
	case '^':
		dx, dy = 0, -1
	case '>':
		dx, dy = 1, 0
	case 'v':
		dx, dy = 0, 1
	case '<':
		dx, dy = -1, 0
	}

	x, y := s.x+dx, s.y+dy
	if x < 0 || y < 0 || y >= len(s.grid) || x >= len(s.grid[y]) {
		// Out of bounds.
		return false, false
	}

	if s.grid[y][x] == '#' {
		// Obstactle in the way. Turn right!
		switch s.dir {
		case '^':
			s.dir = '>'
		case '>':
			s.dir = 'v'
		case 'v':
			s.dir = '<'
		case '<':
			s.dir = '^'
		}
		s.grid[s.y][s.x] = '+'
		return true, false
	}

	visit := '+'
	if dx == 0 {
		visit = '|'
	}
	if dy == 0 {
		visit = '-'
	}

	key := fmt.Sprintf("%d:%d,%d:%d", dx, x, dy, y)
	if s.visits[key] {
		return true, true
	}
	s.visits[key] = true

	if s.grid[y][x] == '.' {
		s.grid[y][x] = visit
	}
	if s.grid[y][x] != visit {
		s.grid[y][x] = '+'
	}

	s.x, s.y = x, y
	return true, false
}

func (s State) Visited() int {
	visits := 0
	for _, row := range s.grid {
		for _, loc := range row {
			if loc == '-' || loc == '|' || loc == '+' {
				visits++
			}
		}
	}
	return visits
}

func (s State) String() string {
	var sb strings.Builder
	sb.WriteString("State:\n")
	for j, row := range s.grid {
		for i, col := range row {
			if s.x == i && s.y == j {
				sb.WriteRune(s.dir)
			} else {
				sb.WriteRune(col)
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func part1(input string) int {
	state := parse(input)
	//fmt.Println(state.String())
	for {
		next, _ := state.Next()
		//fmt.Println(state.String())
		if !next {
			break
		}
	}
	return state.Visited()
}

func (s State) Loops() bool {
	//fmt.Println("MODIFIED", s.String())
	for {
		next, loop := s.Next()
		//fmt.Println("LOOP", s.String())
		if !next {
			//fmt.Println("OOB")
			return false
		}
		if loop {
			//fmt.Println("LOOP")
			return true
		}
	}
}

func (s State) Copy() State {
	grid := make([][]rune, len(s.grid))
	for i, row := range s.grid {
		grid[i] = make([]rune, len(row))
		copy(grid[i], row)
	}
	return State{
		grid:   grid,
		x:      s.x,
		y:      s.y,
		dir:    s.dir,
		visits: make(map[string]bool),
	}
}

func part2(input string) int {
	start := parse(input)

	state := start.Copy()
	for {
		next, _ := state.Next()
		if !next {
			break
		}
	}

	tries := float64(len(state.grid)*len(state.grid[0]) - 1)
	try := 0.0

	count := 0
	for j, row := range state.grid {
		for i, loc := range row {
			if start.x == i && start.y == j {
				continue
			}
			if loc == '-' || loc == '|' || loc == '+' {
				modified := start.Copy()
				modified.grid[j][i] = '#'
				if modified.Loops() {
					modified.grid[j][i] = 'O'
					count++
				}
			}
			fmt.Printf("%f\n", try/tries)
			try++
		}
	}

	return count
}
