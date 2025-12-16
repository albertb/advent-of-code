package main

import "image"

func part1(input string) int {
	where := image.Point{}
	visited := map[image.Point]struct{}{where: {}}
	for _, sym := range input {
		var diff image.Point
		switch sym {
		case '^':
			diff = image.Point{0, -1}
		case 'v':
			diff = image.Point{0, 1}
		case '>':
			diff = image.Point{1, 0}
		case '<':
			diff = image.Point{-1, 0}
		}
		where = where.Add(diff)
		visited[where] = struct{}{}
	}
	return len(visited)
}

func part2(input string) int {
	santa := image.Point{}
	robot := image.Point{}
	visited := map[image.Point]struct{}{santa: {}, robot: {}}
	for i, sym := range input {
		var diff image.Point
		switch sym {
		case '^':
			diff = image.Point{0, -1}
		case 'v':
			diff = image.Point{0, 1}
		case '>':
			diff = image.Point{1, 0}
		case '<':
			diff = image.Point{-1, 0}
		}
		if i%2 == 0 {
			santa = santa.Add(diff)
			visited[santa] = struct{}{}
		} else {
			robot = robot.Add(diff)
			visited[robot] = struct{}{}
		}
	}
	return len(visited)
}
