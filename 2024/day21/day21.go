package main

import (
	"container/heap"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

type Pad struct {
	buttons map[rune]mathy.Vec
	gaps    map[mathy.Vec]struct{}
	bounds  mathy.Vec
}

func newKeypad() Pad {
	return Pad{
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
}

func newDirectionPad() Pad {
	return Pad{
		buttons: map[rune]mathy.Vec{
			/*      GAP     */ '^': {X: 1, Y: 0}, 'A': {X: 2, Y: 0},
			'<': {X: 0, Y: 1}, 'v': {X: 1, Y: 1}, '>': {X: 2, Y: 1},
		},
		gaps: map[mathy.Vec]struct{}{
			{X: 0, Y: 0}: {},
		},
		bounds: mathy.Vec{X: 2, Y: 1},
	}
}

func shortestPath(pad Pad, from, to mathy.Vec) []rune {
	type AstarNode struct {
		where  mathy.Vec
		cost   int
		parent *AstarNode
	}

	nm := map[mathy.Vec]*mathy.PriorityItem[AstarNode]{}
	pq := &mathy.PriorityQueue[AstarNode]{}
	heap.Init(pq)

	// Heuristic is the Manhattan distance.
	h := func(from, to mathy.Vec) int {
		return from.Distance(to)
	}

	item := mathy.PriorityItem[AstarNode]{
		Value: AstarNode{
			where:  from,
			cost:   0,
			parent: nil,
		},
		Priority: h(from, to),
	}
	nm[from] = &item
	heap.Push(pq, &item)

	var path *AstarNode
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*mathy.PriorityItem[AstarNode])

		if current.Value.where.Equals(to) {
			path = &current.Value
			break
		}

		var ahead *mathy.Vec
		if parent := current.Value.parent; parent != nil {
			// Discount moves that don't require turns.
			temp := current.Value.where.Minus(parent.where)
			ahead = &temp
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

			if ahead != nil && !diff.Equals(*ahead) {
				// Penalize turns.
				cost += 10000
			}

			if node, ok := nm[next]; !ok {
				node := mathy.PriorityItem[AstarNode]{
					Value: AstarNode{
						where:  next,
						cost:   cost,
						parent: &current.Value,
					},
					Priority: cost + h(next, to),
				}
				nm[next] = &node
				heap.Push(pq, &node)
			} else if cost <= node.Value.cost {
				node.Value.cost = cost
				node.Value.parent = &current.Value
				node.Priority = cost + h(next, to)
				heap.Push(pq, node)
			}
		}
	}

	if path == nil {
		log.Fatalln("failed to find a path from", from, "to", to)
	}

	runes := []rune{}

	current := path
	previous := path.parent
	for previous != nil {
		diff := current.where.Minus(previous.where)
		if diff.X > 0 {
			runes = append(runes, '>')
		} else if diff.X < 0 {
			runes = append(runes, '<')
		} else if diff.Y > 0 {
			runes = append(runes, 'v')
		} else if diff.Y < 0 {
			runes = append(runes, '^')
		}
		current, previous = previous, current.parent
	}
	slices.Reverse(runes)

	return runes
}

func enter(pad Pad, keys []rune) []rune {
	var presses []rune

	from := pad.buttons['A']
	for _, key := range keys {
		to := pad.buttons[key]
		moves := shortestPath(pad, from, to)
		presses = append(presses, moves...)
		presses = append(presses, 'A')
		from = to
	}

	return presses
}

func sequence(keys string) string {
	//fmt.Println("code:", keys)
	keypad := newKeypad()
	keypadMoves := enter(keypad, []rune(keys))
	//fmt.Println("keys:", string(keypadMoves))
	directionPad := newDirectionPad()
	directionPad1Moves := enter(directionPad, keypadMoves)
	//fmt.Println("pad1:", string(directionPad1Moves))
	directionPad2Moves := enter(directionPad, directionPad1Moves)
	//fmt.Println("pad2:", string(directionPad2Moves))
	//fmt.Println("---")

	return string(directionPad2Moves)
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

func part1(input string) int {
	codes := parse(input)

	sum := 0
	for _, code := range codes {
		moves := sequence(code)
		number, err := strconv.ParseInt(code[0:len(code)-1], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		sum += len(moves) * int(number)
	}
	return sum
}

/* TODO
   - Generate all the shortest paths p0 from each pair of key on keypad
   - For each p0, generate all the shortest paths p1 from each pair of key on the direction pad
   - For each p1, generate all the shortest paths p2 from each pair of key on the direction pad
   - Select p2
*/

func main() {
	var puzzle = `
803A
528A
586A
341A
319A`

	fmt.Println("1:", part1(puzzle))
}
