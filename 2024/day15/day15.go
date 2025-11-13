package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

type World struct {
	robot  mathy.Rect
	boxes  []mathy.Rect
	walls  []mathy.Rect
	bounds mathy.Vec
}

func parse(input string, wide bool) (World, []mathy.Vec) {
	var world World
	var moves []mathy.Vec

	width := 0
	if wide {
		width = 1
	}

	y := 0
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			world.bounds.X = len(line)

			// We're looking at the world.
			x := 0
			for _, rune := range line {
				switch rune {
				case '#':
					world.walls = append(world.walls, mathy.Rect{Vec: mathy.Vec{X: x, Y: y}, Width: width, Height: 0})
				case '.':
					// This is empty space. Ignore.
				case 'O':
					// This is a box.
					world.boxes = append(world.boxes, mathy.Rect{Vec: mathy.Vec{X: x, Y: y}, Width: width, Height: 0})
				case '@':
					if world.robot.X != 0 {
						log.Fatalln("unexpected robot at", x, y)
					}
					world.robot = mathy.Rect{Vec: mathy.Vec{X: x, Y: y}, Width: 0, Height: 0}
				default:
					log.Fatalln("unexpected world tile:", rune)
				}
				x += 1 + width
			}
			y++
		} else {
			// We're looking at moves.
			for _, direction := range line {
				switch direction {
				case '^':
					moves = append(moves, mathy.Vec{X: 0, Y: -1})
				case '>':
					moves = append(moves, mathy.Vec{X: 1, Y: 0})
				case 'v':
					moves = append(moves, mathy.Vec{X: 0, Y: 1})
				case '<':
					moves = append(moves, mathy.Vec{X: -1, Y: 0})
				default:
					log.Fatalf("unexpected move: `%c` in `%s`\n", direction, line)
				}

			}
		}
	}
	world.bounds.Y = y

	return world, moves
}

func (w World) isWall(pos mathy.Rect) bool {
	for _, wall := range w.walls {
		if wall.Intersects(pos) {
			return true
		}
	}
	return false
}

func (w *World) getBox(pos mathy.Rect) *mathy.Rect {
	for i := range w.boxes {
		box := &w.boxes[i]
		if box.Intersects(pos) {
			return box
		}
	}
	return nil
}

func (w *World) apply(vec mathy.Vec) {
	next := mathy.Rect{Vec: w.robot.Plus(vec), Width: 0, Height: 0}

	if w.isWall(next) {
		// We can't move.
		return
	}

	box := w.getBox(next)
	if box != nil {
		// There's a box in the way, see if we can move it.
		blocked := true
		boxesToMove := []*mathy.Rect{box}
		nextBox := next
		for {
			nextBox = mathy.Rect{Vec: nextBox.Plus(vec), Width: 0, Height: 0}
			if w.isWall(nextBox) {
				// We can't move the boxes, give up.
				break
			}
			if box = w.getBox(nextBox); box != nil {
				// We found another box, keep going.
				boxesToMove = append(boxesToMove, box)
				continue
			}
			// We found a free space for the box!
			blocked = false
			break
		}

		if blocked {
			return
		}

		for i := range boxesToMove {
			boxesToMove[i].Add(vec)
		}
	}

	// We're free to move to the next tile.
	w.robot.Add(vec)
}

func part1(input string) int {
	world, moves := parse(input, false)
	for _, move := range moves {
		world.apply(move)
	}

	sum := 0
	for _, box := range world.boxes {
		sum += box.X + box.Y*100
	}
	return sum
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle))
}
