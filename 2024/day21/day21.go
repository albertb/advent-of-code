package main

import (
	"container/heap"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

type Path struct {
	from, to string
}

type Pad struct {
	buttons map[rune]mathy.Vec
	gaps    map[mathy.Vec]struct{}
	bounds  mathy.Vec
	paths   map[Path]string
}

func newKeypad() Pad {
	pad := Pad{
		buttons: map[rune]mathy.Vec{
			'7': {X: 0, Y: 0}, '8': {X: 1, Y: 0}, '9': {X: 2, Y: 0},
			'4': {X: 0, Y: 1}, '5': {X: 1, Y: 1}, '6': {X: 2, Y: 1},
			'1': {X: 0, Y: 2}, '2': {X: 1, Y: 2}, '3': {X: 2, Y: 2},
			/*      GAP     */ '0': {X: 1, Y: 3}, 'A': {X: 2, Y: 3},
		},
		gaps: map[mathy.Vec]struct{}{
			{X: 0, Y: 3}: {},
		},
		bounds: mathy.Vec{X: 2, Y: 3},
	}
	pad.paths = mapBestPaths(pad)

	return pad
}

func newDirectionPad() Pad {
	pad := Pad{
		buttons: map[rune]mathy.Vec{
			/*      GAP     */ '^': {X: 1, Y: 0}, 'A': {X: 2, Y: 0},
			'<': {X: 0, Y: 1}, 'v': {X: 1, Y: 1}, '>': {X: 2, Y: 1},
		},
		gaps: map[mathy.Vec]struct{}{
			{X: 0, Y: 0}: {},
		},
		bounds: mathy.Vec{X: 2, Y: 1},
	}
	pad.paths = mapBestPaths(pad)
	return pad
}

func mapBestPaths(pad Pad) map[Path]string {
	result := make(map[Path]string)
	for fromLabel, fromWhere := range pad.buttons {
		for toLabel, toWhere := range pad.buttons {
			if fromWhere.Equals(toWhere) {
				continue
			}
			if fromLabel == 'A' && toLabel == '<' {
				fmt.Println("aaa")
			}
			path := bestPath(pad, fromWhere, toWhere) + "A"
			result[Path{string(fromLabel), string(toLabel)}] = path
		}
	}
	return result
}

func bestPath(pad Pad, from, to mathy.Vec) string {
	type AstarNode struct {
		where  mathy.Vec
		cost   int
		parent *AstarNode
	}

	// Keep track of the min-cost at each location.
	nm := map[mathy.Vec]int{from: 0}

	pq := &mathy.PriorityQueue[AstarNode]{}
	heap.Init(pq)

	// Heuristic is the Manhattan distance.
	h := func(from, to mathy.Vec) int {
		return from.Distance(to)
	}

	item := mathy.PriorityItem[AstarNode]{
		Value: AstarNode{
			where: from,
			cost:  0,
		},
		Parent:   nil,
		Priority: h(from, to),
	}
	heap.Push(pq, &item)

	var shortest *mathy.PriorityItem[AstarNode]
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*mathy.PriorityItem[AstarNode])

		var discounted *mathy.Vec
		if current.Parent != nil {
			diff := current.Value.where.Minus(current.Parent.Value.where)
			discounted = &diff
		}

		if current.Value.where.Equals(to) {
			shortest = current
			break
		}

		for _, diff := range mathy.Cardinals() {
			next := current.Value.where.Plus(diff)
			if next.X < 0 || next.Y < 0 || next.X > pad.bounds.X || next.Y > pad.bounds.Y {
				// Out of bounds.
				continue
			}
			if _, ok := pad.gaps[next]; ok {
				// We can't go over the gap.
				continue
			}

			cost := 1 + current.Value.cost

			if discounted != nil && !discounted.Equals(diff) {
				cost += 1000
			}

			if minCost, ok := nm[next]; !ok || minCost >= cost {
				nm[next] = cost

				node := mathy.PriorityItem[AstarNode]{
					Value: AstarNode{
						where: next,
						cost:  cost,
					},
					Parent:   current,
					Priority: cost + h(next, to),
				}
				heap.Push(pq, &node)
			}
		}
	}

	if shortest == nil {
		log.Fatalln("failed to find a path from", from, "to", to)
	}

	sequence := []rune{}
	current := shortest
	previous := current.Parent
	for previous != nil {
		diff := current.Value.where.Minus(previous.Value.where)
		if diff.X > 0 {
			sequence = append(sequence, '>')
		} else if diff.X < 0 {
			sequence = append(sequence, '<')
		} else if diff.Y > 0 {
			sequence = append(sequence, 'v')
		} else if diff.Y < 0 {
			sequence = append(sequence, '^')
		}
		current, previous = previous, current.Parent
	}
	slices.Reverse(sequence)

	return string(sequence)
}

func sequences(sequence string, pad Pad) string {
	var keys string
	previous := "A"
	for _, key := range sequence {
		next := string(key)
		keys += pad.paths[Path{previous, next}]
		previous = next
	}
	return keys
}

func part1(input string) int {
	code := "029A"
	fmt.Println("code:", code)

	keyPad := newKeypad()
	dirPad := newDirectionPad()

	fmt.Println("[A to <]:", dirPad.paths[Path{"A", "<"}])

	keySeq := sequences(code, keyPad)
	fmt.Println("keypad:", keySeq)

	dir1Seq := sequences(keySeq, dirPad)
	fmt.Println("dirpad:", dir1Seq)

	dir2Seq := sequences(dir1Seq, dirPad)
	fmt.Println("dirpad:", dir2Seq)

	return 0
}

func parse(input string) []string {
	var result []string
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		result = append(result, line)
	}
	return result
}

var puzzle = `
803A
528A
586A
341A
319A`

func main() {
	fmt.Println("Part 1:", part1(puzzle))
}
