package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:", part(input, false))
	fmt.Println("Part 2:", part(input, true))
}

type coord struct {
	x, y int
}

func (c coord) oob(bounds coord) bool {
	return c.x < 0 || c.y < 0 || c.x >= bounds.x || c.y >= bounds.y
}

func antinodes(c1, c2 coord, bounds coord, resonant bool) []coord {
	var result []coord

	lastAnti := c1
	for {
		anti := coord{
			lastAnti.x - (c2.x - c1.x),
			lastAnti.y - (c2.y - c1.y),
		}
		if anti.oob(bounds) {
			break
		}
		result = append(result, anti)
		lastAnti = anti
		if !resonant {
			break
		}
	}

	lastAnti = c2
	for {
		anti := coord{
			lastAnti.x - (c1.x - c2.x),
			lastAnti.y - (c1.y - c2.y),
		}
		if anti.oob(bounds) {
			break
		}
		result = append(result, anti)
		lastAnti = anti
		if !resonant {
			break
		}
	}

	if resonant {
		result = append(result, c1)
		result = append(result, c2)
	}

	return result
}

type grid struct {
	bounds   coord
	antennas map[rune][]coord
}

func parse(input string) grid {
	var result grid
	result.antennas = make(map[rune][]coord)

	y := 0
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		for x, rune := range line {
			if rune == '.' {
				continue
			}
			result.antennas[rune] = append(result.antennas[rune], coord{x, y})
		}
		y++

		if len(line) > result.bounds.x {
			result.bounds.x = len(line)
		}

	}
	result.bounds.y = y
	return result
}

func part(input string, resonant bool) int {
	m := parse(input)

	nodess := map[coord]bool{}
	for _, antenna := range m.antennas {
		for i, first := range antenna {
			for j := i + 1; j < len(antenna); j++ {
				second := antenna[j]
				nodes := antinodes(first, second, m.bounds, resonant)

				for _, node := range nodes {
					nodess[node] = true
				}
			}
		}
	}
	return len(nodess)
}
