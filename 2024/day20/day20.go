package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"math"
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

func heuristic(from mathy.Vec, to mathy.Vec) int {
	return mathy.Abs(to.X-from.X) + mathy.Abs(to.Y-from.Y)
}

type AStarNode struct {
	pos    mathy.Vec
	cost   int
	rank   int
	parent *AStarNode
}

type AstarHeap []*AStarNode

func (h AstarHeap) Len() int           { return len(h) }
func (h AstarHeap) Less(i, j int) bool { return h[i].rank < h[j].rank }
func (h AstarHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *AstarHeap) Push(item any) {
	*h = append(*h, item.(*AStarNode))
}

func (h *AstarHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func astar(m Maze) int {
	nm := map[mathy.Vec]*AStarNode{}
	pq := &AstarHeap{}
	heap.Init(pq)

	node := AStarNode{
		pos:  m.start,
		cost: 0,
		rank: heuristic(m.start, m.end),
	}
	heap.Push(pq, &node)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*AStarNode)

		if current.pos.Equals(m.end) {
			return current.cost
		}

		for _, dir := range mathy.RightLeftUpDown() {
			next := current.pos.Plus(dir)

			if next.X < 0 || next.Y < 0 || next.X > m.bounds.X || next.Y > m.bounds.Y {
				// Out of bounds.
				continue
			}

			cost := 1 + current.cost

			if _, ok := m.walls[next]; ok {
				continue
			}

			if node, ok := nm[next]; !ok {
				node := AStarNode{
					pos:    next,
					cost:   cost,
					rank:   cost + heuristic(next, m.end),
					parent: current,
				}
				nm[next] = &node
				heap.Push(pq, &node)
			} else if cost <= node.cost {
				node.cost = cost
				node.rank = cost + heuristic(next, m.end)
				node.parent = current
				heap.Push(pq, node)
			}
		}
	}

	return math.MaxInt
}

func dfs(m Maze, where, cheat mathy.Vec, cost, max int, visited map[mathy.Vec]struct{}) map[mathy.Vec]int {

	// XXX
	/*grid := make([][]rune, m.bounds.Y+1)
	for y := range m.bounds.Y + 1 {
		grid[y] = make([]rune, m.bounds.X+1)
	}
	grid[m.start.Y][m.start.X] = 'S'
	grid[m.end.Y][m.end.X] = 'E'
	for w := range m.walls {
		grid[w.Y][w.X] = '#'
	}
	for v := range visited {
		grid[v.Y][v.X] = 'O'
	}
	if !cheat.Equals(mathy.Vec{}) {
		grid[cheat.Y][cheat.X] = '$'
	}
	grid[where.Y][where.X] = 'X'

	fmt.Println("cost", cost)
	for _, row := range grid {
		for _, col := range row {
			if col == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(col))
			}
		}
		fmt.Println()
	}
	fmt.Scanln()*/
	// XXX

	if cost > max {
		return map[mathy.Vec]int{}
	}

	if where.Equals(m.end) {
		//fmt.Println("cheat", cheat, "cost", cost)
		return map[mathy.Vec]int{cheat: cost}
	}
	visited[where] = struct{}{}

	costs := map[mathy.Vec]int{}
	for _, dir := range mathy.RightLeftUpDown() {
		nextWhere := where.Plus(dir)
		nextCost := cost + 1
		nextCheat := cheat

		if _, ok := visited[nextWhere]; ok {
			// We already visited this node, give up on this branch.
			continue
		}

		if nextWhere.X < 0 || nextWhere.Y < 0 || nextWhere.X > m.bounds.X || nextWhere.Y > m.bounds.Y {
			// Out of bounds.
			continue
		}

		if _, ok := m.walls[nextWhere]; ok {
			if !nextCheat.Equals(mathy.Vec{}) {
				// We used our cheat already, give up on this branch.
				continue
			}
			// Try to use the cheat to get across this wall.
			nextCheat = nextWhere
			//nextCost++
			next := nextWhere.Plus(dir)
			if _, ok := m.walls[next]; ok {
				// We hit a second wall, give up.
				continue
			}
		}

		nextCosts := dfs(m, nextWhere, nextCheat, nextCost, max, visited)
		for k, v := range nextCosts {
			if oldV, ok := costs[k]; !ok || oldV > v {
				costs[k] = v
			}
		}
	}
	delete(visited, where)

	return costs
}

func part1(input string, minSaving int) int {
	m := parse(input)

	cheapest := astar(m)
	cheatCosts := dfs(m, m.start, mathy.Vec{}, 0, cheapest-minSaving, map[mathy.Vec]struct{}{})

	//XXX
	/*costs := map[int]int{}
	for _, path := range cheatPaths {
		costs[path]++
	}
	for k, v := range costs {
		fmt.Println("there are", v, "cheat(s) that saves", (shortestPath - k))
	}*/
	//XXX

	sum := 0
	for _, cost := range cheatCosts {
		if cost <= cheapest-minSaving {
			sum++
		}
	}
	return sum
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle, 100))
}
