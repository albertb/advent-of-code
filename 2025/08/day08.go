package main

import (
	"cmp"
	_ "embed"
	"log"
	"slices"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

func parse(input string) []mathy.Vec3 {
	var box []mathy.Vec3
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		components := strings.Split(line, ",")
		if len(components) != 3 {
			log.Fatalln("unexpected vector:", line)
		}
		box = append(box, mathy.Vec3{
			mathy.MustParseInt(components[0]),
			mathy.MustParseInt(components[1]),
			mathy.MustParseInt(components[2]),
		})
	}
	return box
}

type PairOfBoxes struct {
	from, to mathy.Vec3
	distance float64
}

func part1(input string, connections int) int {
	boxes := parse(input)

	pairs := []PairOfBoxes{}

	// Compute the distaces between every pair of junction box.
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, PairOfBoxes{
				from:     boxes[i],
				to:       boxes[j],
				distance: boxes[i].Distance(boxes[j]),
			})
		}
	}

	// Sort the pairs by distance, shortest distance first.
	slices.SortFunc(pairs, func(a, b PairOfBoxes) int {
		return cmp.Compare(a.distance, b.distance)
	})

	// Add all the junction boxes a disjoint-set.
	dsu := mathy.NewDSU()
	for _, box := range boxes {
		dsu.Add(box)
	}

	// Connect the specified number of boxes, in ascending distance order.
	for i := 0; i < connections; i++ {
		pair := pairs[i]
		dsu.Union(pair.from, pair.to)
	}

	// Each sub-graph has one root box, count the number of times each root box appears
	// to figure out the size of each circuit.
	circuits := map[mathy.Vec3]int{}
	for _, box := range boxes {
		root := dsu.Find(box)
		if root == nil {
			log.Fatalln("null root for:", box)
		}

		rootBox := root.(mathy.Vec3)
		circuits[rootBox]++
	}

	// Make a list of all the circuit lengths, and sort them in descending order.
	lengths := []int{}
	for _, length := range circuits {
		lengths = append(lengths, length)
	}
	slices.Sort(lengths)
	slices.Reverse(lengths)

	// The result is the product of the 3 longest circuits.
	result := 1
	for _, length := range lengths[:3] {
		result *= length
	}
	return result
}

func part2(input string) int {
	boxes := parse(input)

	pairs := []PairOfBoxes{}

	// Compute the distaces between every pair of junction box.
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, PairOfBoxes{
				from:     boxes[i],
				to:       boxes[j],
				distance: boxes[i].Distance(boxes[j]),
			})
		}
	}

	// Sort the pairs by distance, shortest distance first.
	slices.SortFunc(pairs, func(a, b PairOfBoxes) int {
		return cmp.Compare(a.distance, b.distance)
	})

	// Add all the junction boxes a disjoint-set.
	dsu := mathy.NewDSU()
	for _, box := range boxes {
		dsu.Add(box)
	}

	// Connect the specified number of boxes, in ascending distance order.
	// Stop once all the boxes are part of a single circuit.
	for _, pair := range pairs {
		if dsu.Union(pair.from, pair.to) == len(boxes) {
			return pair.from.X * pair.to.X
		}
	}

	return -1
}

//go:embed puzzle.txt
var puzzle string
