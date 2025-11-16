package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

type Memory struct {
	corrupted map[mathy.Vec]struct{}
	end       mathy.Vec
}

func parse(input string, n int) map[mathy.Vec]struct{} {
	mustParse := func(s string) int {
		n, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			log.Fatalln(err)
		}
		return int(n)
	}

	i := 1
	corrupted := make(map[mathy.Vec]struct{})
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		if i > n {
			break
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			log.Fatalln("unexpected input:", line)
		}
		x, y := mustParse(parts[0]), mustParse(parts[1])
		corrupted[mathy.Vec{X: x, Y: y}] = struct{}{}
		i++
	}
	return corrupted
}

func heuristic(pos mathy.Vec, end mathy.Vec) int {
	return mathy.Abs(end.X-pos.X) + mathy.Abs(end.Y-pos.Y)
}

type Node struct {
	pos    mathy.Vec
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

func part1(input string, n int, end mathy.Vec) int {
	m := Memory{
		corrupted: parse(input, n),
		end:       end,
	}

	nm := map[mathy.Vec]*Node{}
	pq := &AstarHeap{}
	heap.Init(pq)

	start := mathy.Vec{}
	node := Node{
		pos:  start,
		cost: 0,
		rank: heuristic(start, m.end),
	}
	heap.Push(pq, &node)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)

		if current.pos.Equals(m.end) {

			// XXX
			/*grid := make([][]rune, m.end.Y+1)
			for y := range m.end.Y + 1 {
				grid[y] = make([]rune, m.end.X+1)
				for x := range m.end.X + 1 {
					grid[y][x] = '.'
				}
			}

			for c := range m.corrupted {
				grid[c.Y][c.X] = '#'
			}

			node := current
			for node != nil {
				grid[node.pos.Y][node.pos.X] = 'O'
				node = node.parent
			}

			for _, line := range grid {
				for _, col := range line {
					fmt.Print(string(col))
				}
				fmt.Println()
			}
			fmt.Println("EXIT AT", current.cost)*/
			// XXX

			return current.cost
		}

		for _, dir := range []mathy.Vec{
			mathy.NewVec(1, 0),
			mathy.NewVec(-1, 0),
			mathy.NewVec(0, 1),
			mathy.NewVec(0, -1),
		} {
			next := current.pos.Plus(dir)
			if next.X < 0 || next.Y < 0 || next.X > m.end.X || next.Y > m.end.Y {
				// Out of bounds.
				continue
			}

			if _, ok := m.corrupted[next]; ok {
				// The cell is corrupted.
				continue
			}

			cost := 1 + current.cost
			node, ok := nm[next]
			if !ok {
				node := Node{
					pos:    next,
					cost:   cost,
					rank:   cost + heuristic(next, m.end),
					parent: current,
				}
				nm[next] = &node
				heap.Push(pq, &node)
			} else if cost < node.cost {
				node.cost = cost
				node.rank = cost + heuristic(next, m.end)
				node.parent = current
				heap.Push(pq, node)
			}
		}
	}

	return -1
}

func part2(input string, end mathy.Vec) string {
	n := 1025
	for {
		if n > 3450 {
			return "error"
		}
		if part1(input, n, end) == -1 {
			break
		}
		n++
	}

	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		if n == 1 {
			return line
		}
		n--
	}
	return "error"
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle, 1024, mathy.Vec{X: 70, Y: 70}))
	fmt.Println("2:", part2(puzzle, mathy.Vec{X: 70, Y: 70}))
}
