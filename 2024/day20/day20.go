package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

type Maze struct {
	start  mathy.Vec
	end    mathy.Vec
	walls  map[mathy.Vec]struct{}
	bounds mathy.Vec
}

func parse(input string) Maze {
	var m Maze
	m.walls = make(map[mathy.Vec]struct{})

	y := 0
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		for x, rune := range line {
			switch rune {
			case '#':
				m.walls[mathy.Vec{X: x, Y: y}] = struct{}{}
			case 'S':
				m.start = mathy.Vec{X: x, Y: y}
			case 'E':
				m.end = mathy.Vec{X: x, Y: y}
			}
			if x > m.bounds.X {
				m.bounds.X = x
			}
		}
		y++
	}
	m.bounds.Y = y - 1
	return m
}

func dijkstra(m Maze, from mathy.Vec) map[mathy.Vec]int {
	costs := map[mathy.Vec]int{}
	costs[from] = 0

	item := mathy.PriorityItem[mathy.Vec]{
		Value:    from,
		Priority: 0,
	}

	pq := &mathy.PriorityQueue[mathy.Vec]{}
	heap.Init(pq)

	heap.Push(pq, &item)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*mathy.PriorityItem[mathy.Vec])

		where := current.Value
		nextCost := 1 + current.Priority

		for _, dir := range mathy.Cardinals() {
			next := where.Plus(dir)

			if next.X < 0 || next.Y < 0 || next.X > m.bounds.X || next.Y > m.bounds.Y {
				// Out of bounds.
				continue
			}

			if _, ok := m.walls[next]; ok {
				// Cannot go through walls.
				continue
			}

			if cost, ok := costs[next]; !ok || nextCost < cost {
				costs[next] = nextCost

				item := mathy.PriorityItem[mathy.Vec]{
					Value:    next,
					Priority: nextCost,
				}
				heap.Push(pq, &item)
			}
		}
	}

	return costs
}

func part1(input string, min int) int {
	m := parse(input)

	fromStart := dijkstra(m, m.start)
	fromEnd := dijkstra(m, m.end)

	cost := fromStart[m.end]

	cheats := 0
	for where := range fromStart {
		costStart := fromStart[where]

		for _, dir := range mathy.Cardinals() {
			// Jump two cells away.
			next := where.Plus(dir).Plus(dir)

			// If there's no end cost on that cell, either it's a wall or cannot reach the end.
			nextCostEnd, ok := fromEnd[next]
			if !ok {
				continue
			}

			cheatCost := costStart + 2 + nextCostEnd
			if cheatCost <= cost-min {
				cheats++
			}
		}
	}

	return cheats
}

func shortcuts(from mathy.Vec, max int) map[mathy.Vec]int {
	results := map[mathy.Vec]int{}
	for x := -max; x <= max; x++ {
		maxy := max - mathy.Abs(x)
		for y := -maxy; y <= maxy; y++ {
			results[mathy.Vec{X: from.X + x, Y: from.Y + y}] = mathy.Abs(x) + mathy.Abs(y)
		}
	}
	return results
}

func part2(input string, min int) int {
	m := parse(input)

	fromStart := dijkstra(m, m.start)
	fromEnd := dijkstra(m, m.end)
	normalCost := fromStart[m.end]

	cheats := 0
	for where := range fromStart {
		costStart := fromStart[where]

		for next, len := range shortcuts(where, 20) {
			// If there's no end cost on that cell, either it's a wall or cannot reach the end.
			nextCostEnd, ok := fromEnd[next]
			if !ok {
				continue
			}

			cheatCost := costStart + len + nextCostEnd
			if cheatCost <= normalCost-min {
				cheats++
			}
		}
	}

	return cheats
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle, 100))
	fmt.Println("2:", part2(puzzle, 100))
}
