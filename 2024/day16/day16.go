package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

type Tile int

const (
	Free Tile = iota
	Wall
)

type Maze struct {
	start mathy.Vec
	end   mathy.Vec
	grid  [][]Tile
}

func parse(input string) Maze {
	var maze Maze

	y := 0
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			continue
		}
		row := []Tile{}
		for x, rune := range line {
			switch rune {
			case '.':
				row = append(row, Free)
			case '#':
				row = append(row, Wall)
			case 'S':
				row = append(row, Free)
				maze.start = mathy.Vec{X: x, Y: y}
			case 'E':
				row = append(row, Free)
				maze.end = mathy.Vec{X: x, Y: y}
			default:
				log.Fatalln("Unexpected tile:", string(rune))
			}
		}
		maze.grid = append(maze.grid, row)
		y++
	}

	return maze
}

type Reindeer struct {
	pos mathy.Vec
	dir mathy.Vec
}

// Heuristic function for A*: Manhattan distance + 1000*number of rotations.
func heuristic(r Reindeer, m Maze) int {
	cost := mathy.Abs(m.end.X-r.pos.X) + mathy.Abs(m.end.Y-r.pos.Y)

	/*
		TODO: Add rotation to cost heuristic

		dot := r.dir.Dot(m.end)
		cross := r.dir.Cross(m.end)

		if dot >= 0 {
			if cross >= 0 {
				// No rotation necessary.
				return cost
			} else {
				// One rotation necessary.
				return cost + 1000
			}
		} else {
			if cross >= 0 {
				// One rotation necessary.
				return cost + 1000
			} else {
				// Two rotation necessary.
				return cost + 2000
			}
		}
	*/

	return cost
}

type Node struct {
	r      Reindeer
	cost   int
	rank   int
	parent *Node
}

type AstarHeap []*Node

func (h AstarHeap) Len() int           { return len(h) }
func (h AstarHeap) Less(i, j int) bool { return h[i].rank < h[j].rank }
func (h AstarHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *AstarHeap) Push(item any) {
	*h = append(*h, item.(*Node))
}

func (h *AstarHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func cost(r Reindeer, m Maze) int {
	nm := map[Reindeer]*Node{}
	pq := &AstarHeap{}
	heap.Init(pq)

	node := Node{
		r:      r,
		cost:   0,
		rank:   heuristic(r, m),
		parent: nil,
	}
	heap.Push(pq, &node)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)

		if current.r.pos.Equals(m.end) {
			node := current
			for {
				if node == nil {
					break
				}
				node = node.parent
			}

			// We reached the end!
			return current.cost
		}

		// Try to see if the reindeer can just walk forward, whatever way it's pointing.
		forward := Reindeer{
			pos: current.r.pos.Plus(current.r.dir),
			dir: current.r.dir,
		}

		if m.grid[forward.pos.Y][forward.pos.X] == Free {
			cost := 1 + current.cost
			forwardNode, ok := nm[forward]
			if !ok {
				node := Node{
					r:      forward,
					cost:   cost,
					rank:   cost + heuristic(forward, m),
					parent: current,
				}
				nm[forward] = &node
				heap.Push(pq, &node)
			} else if cost < forwardNode.cost {
				forwardNode.cost = cost
				forwardNode.rank = cost + heuristic(forward, m)
				forwardNode.parent = current
				heap.Push(pq, forwardNode)
			}
		}

		// Rotate the reindeer 90deg clockwise, and 90deg counter-clockwise.
		for _, n := range []int{-1, 1} {
			rotated := Reindeer{
				pos: current.r.pos,
				dir: current.r.dir.Rotate90(n),
			}

			cost := 1000 + current.cost
			rotatedNode, ok := nm[rotated]
			if !ok {
				node := Node{
					r:      rotated,
					cost:   cost,
					rank:   cost + heuristic(rotated, m),
					parent: current,
				}
				nm[rotated] = &node
				heap.Push(pq, &node)
			} else if cost < rotatedNode.cost {
				rotatedNode.cost = cost
				rotatedNode.rank = cost + heuristic(rotated, m)
				rotatedNode.parent = current
				heap.Push(pq, rotatedNode)
			}
		}
	}

	// No path found.
	return math.MaxInt
}

func part1(input string) int {
	maze := parse(input)

	r := Reindeer{
		pos: maze.start,
		dir: mathy.Vec{X: 1, Y: 0},
	}

	return cost(r, maze)
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle))
}
