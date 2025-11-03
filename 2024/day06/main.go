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
}

type State struct {
	walls  [][]bool
	x, y   int
	dir    rune
	visits [][]bool
}

func parse(input string) State {
	var result State
	j := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) < 1 {
			continue
		}

		visitRow := []bool{}

		row := []bool{}
		for i, rune := range line {
			visitRow = append(visitRow, false)

			if rune == '#' {
				row = append(row, true)
				continue
			}

			row = append(row, false)
			if rune == '.' {
				continue
			}

			result.x = i
			result.y = j
			result.dir = rune
		}

		result.visits = append(result.visits, visitRow)
		result.walls = append(result.walls, row)
		j++
	}
	result.visits[result.y][result.x] = true

	return result
}

func (s *State) Next() bool {
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
	if x < 0 || y < 0 || y >= len(s.walls) || x >= len(s.walls[y]) {
		// Out of bounds.
		return false
	}

	if s.walls[y][x] {
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
		return true
	}

	s.x, s.y = x, y
	s.visits[y][x] = true
	return true
}

func (s State) Visited() int {
	visits := 0
	for _, row := range s.visits {
		for _, visit := range row {
			if visit {
				visits++
			}
		}
	}
	return visits
}

func (s State) ToString() string {
	var sb strings.Builder
	sb.WriteString("State:")
	for j, row := range s.walls {
		for i, col := range row {
			if s.x == i && s.y == j {
				sb.WriteRune(s.dir)
			} else {
				if col {
					sb.WriteRune('#')
				} else {
					if s.visits[j][i] {
						sb.WriteRune('X')
					} else {
						sb.WriteRune('.')
					}
				}
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func part1(input string) int {
	state := parse(input)
	fmt.Println(state.ToString())
	for {
		next := state.Next()
		//fmt.Println(state.ToString())
		if !next {
			break
		}
	}
	return state.Visited()
}
