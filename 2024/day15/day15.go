package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
)

type Vec struct {
	x, y int
}

func (v Vec) plus(other Vec) Vec {
	return Vec{
		v.x + other.x,
		v.y + other.y,
	}
}

func (v *Vec) add(other Vec) {
	tmp := v.plus(other)
	v.x, v.y = tmp.x, tmp.y
}

func (v Vec) equal(other Vec) bool {
	return v.x == other.x && v.y == other.y
}

type World struct {
	robot  Vec
	boxes  []Vec
	walls  []Vec
	bounds Vec
}

type Move int

const (
	Up = iota
	Right
	Down
	Left
)

func parse(input string) (World, []Move) {
	var world World
	var moves []Move

	y := 0
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			world.bounds.x = len(line)

			// We're looking at the world.
			for x, rune := range line {
				switch rune {
				case '#':
					world.walls = append(world.walls, Vec{x, y})
				case '.':
					continue // This is empty space.
				case 'O':
					// This is a box.
					world.boxes = append(world.boxes, Vec{x, y})
				case '@':
					if world.robot.x != 0 {
						log.Fatalln("unexpected robot at", x, y)
					}
					world.robot = Vec{x, y}
				default:
					log.Fatalln("unexpected world tile:", rune)
				}

			}
			y++
		} else {
			// We're looking at moves.
			for _, direction := range line {
				switch direction {
				case '^':
					moves = append(moves, Up)
				case '>':
					moves = append(moves, Right)
				case 'v':
					moves = append(moves, Down)
				case '<':
					moves = append(moves, Left)
				default:
					log.Fatalf("unexpected move: `%c` in `%s`\n", direction, line)
				}

			}
		}
	}
	world.bounds.y = y

	return world, moves
}

func (m Move) toVec() Vec {
	switch m {
	case Up:
		return Vec{0, -1}
	case Right:
		return Vec{1, 0}
	case Down:
		return Vec{0, 1}
	case Left:
		return Vec{-1, 0}
	}
	log.Fatalln("unexpected move:", m)
	return Vec{}
}

func (w World) isWall(pos Vec) bool {
	for _, wall := range w.walls {
		if pos.equal(wall) {
			return true
		}
	}
	return false
}

func (w *World) getBox(pos Vec) *Vec {
	for i := range w.boxes {
		box := &w.boxes[i]
		if pos.equal(*box) {
			return box
		}
	}
	return nil
}

func (w *World) apply(m Move) {
	vec := m.toVec()
	next := w.robot.plus(vec)

	if w.isWall(next) {
		// We can't move.
		return
	}

	box := w.getBox(next)
	if box != nil {
		// There's a box in the way, see if we can move it.
		blocked := true
		boxesToMove := []*Vec{box}
		nextBox := next
		for {
			nextBox = nextBox.plus(vec)
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
			boxesToMove[i].add(vec)
		}
	}

	// We're free to move to the next tile.
	w.robot.add(vec)
}

func part1(input string) int {
	world, moves := parse(input)
	for _, move := range moves {
		world.apply(move)
	}

	sum := 0
	for _, box := range world.boxes {
		sum += box.x + box.y*100
	}
	return sum
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle))
}
