package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Vec struct {
	x, y int
}

func (v *Vec) add(other Vec) {
	v.x += other.x
	v.y += other.y
}

func (v *Vec) wrap(other Vec) {
	v.x = (v.x%other.x + other.x) % other.x
	v.y = (v.y%other.y + other.y) % other.y
}

type Robot struct {
	position Vec
	movement Vec
}

func parse(input string) []Robot {
	var re = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	mustParse := func(s string) int {
		num, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			log.Fatalln("failed to parse", s, ":", err)
		}
		return int(num)
	}

	var robots []Robot
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			continue
		}
		match := re.FindStringSubmatch(line)
		if len(match) != 5 {
			log.Fatalln("failed to match:", line)
		}

		robot := Robot{
			position: Vec{
				mustParse(match[1]),
				mustParse(match[2]),
			},
			movement: Vec{
				mustParse(match[3]),
				mustParse(match[4]),
			},
		}
		robots = append(robots, robot)
	}
	return robots
}

func part1(input string, space Vec, iters int) int {
	robots := parse(input)

	for range iters {
		for i := range robots {
			robot := &robots[i]
			robot.position.add(robot.movement)
			robot.position.wrap(space)
		}
	}

	halfX := space.x / 2
	halfY := space.y / 2

	topLeft, topRight, botLeft, botRight := 0, 0, 0, 0

	for _, robot := range robots {
		pos := robot.position
		if pos.x < halfX {
			if pos.y < halfY {
				topLeft++
			}
			if pos.y > halfY {
				botLeft++
			}
		}
		if pos.x > halfX {
			if pos.y < halfY {
				topRight++
			}
			if pos.y > halfY {
				botRight++
			}
		}
	}

	return topLeft * topRight * botLeft * botRight
}

func length(pos Vec, grid [][]bool, space Vec, visited map[Vec]struct{}) int {
	if pos.x < 0 || pos.y < 0 || pos.x >= space.x || pos.y >= space.y {
		return 0
	}
	if !grid[pos.x][pos.y] {
		return 0
	}
	if _, ok := visited[pos]; ok {
		return 0
	}

	downRight := length(Vec{pos.x + 1, pos.y + 1}, grid, space, visited)
	down := length(Vec{pos.x, pos.y + 1}, grid, space, visited)
	downLeft := length(Vec{pos.x - 1, pos.y + 1}, grid, space, visited)

	lengths := []int{downRight, down, downLeft}
	return 1 + slices.Max(lengths)
}

func part2(input string, space Vec) {
	robots := parse(input)

	grid := make([][]bool, space.x)
	for i := range space.x {
		grid[i] = make([]bool, space.y)
	}

	i := 0
	for {
		for x := range space.x {
			for y := range space.y {
				grid[x][y] = false
			}
		}

		for i := range robots {
			robot := &robots[i]

			robot.position.add(robot.movement)
			robot.position.wrap(space)

			grid[robot.position.x][robot.position.y] = true
		}
		i++

		// Look for lines?
		minLength := 20
		print := false
		visited := make(map[Vec]struct{})
		for x := range space.x {
			for y := range space.y {
				length := length(Vec{x, y}, grid, space, visited)
				if length > minLength {
					print = true
					break
				}
			}
			if print {
				break
			}
		}

		if print {
			fmt.Println("ITERATION", i)
			for x := range space.x {
				for y := range space.y {
					if grid[x][y] {
						fmt.Print("*")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println()
			}
			log.Fatalln("end")
		}
	}
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle, Vec{101, 103}, 100))
	part2(puzzle, Vec{101, 103})
}
