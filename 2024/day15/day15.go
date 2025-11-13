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

func (w World) debug(desc string) {
	runes := make([][]rune, w.bounds.Y)
	for y := range w.bounds.Y {
		runes[y] = make([]rune, w.bounds.X)
		for x := range w.bounds.X {
			runes[y][x] = '.'
		}
	}

	for _, wall := range w.walls {
		runes[wall.Y][wall.X] = '#'
		if wall.Width > 0 {
			runes[wall.Y][wall.X+1] = '#'
		}
	}

	for _, box := range w.boxes {
		if box.Width > 0 {
			runes[box.Y][box.X] = '['
			runes[box.Y][box.X+1] = ']'
		} else {
			runes[box.Y][box.X] = 'O'
		}
	}

	runes[w.robot.Y][w.robot.X] = '@'

	fmt.Println(desc)
	for _, line := range runes {
		for _, r := range line {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
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
			// We're looking at the world.
			x := 0
			for _, rune := range line {
				maxX := x + 1 + width
				if maxX > world.bounds.X {
					world.bounds.X = maxX
				}

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

func (w World) intersectsWall(pos mathy.Rect) bool {
	for _, wall := range w.walls {
		if wall.Intersects(pos) {
			return true
		}
	}
	return false
}

func (w *World) intersectingBoxes(pos mathy.Rect) []*mathy.Rect {
	var boxes []*mathy.Rect
	for i := range w.boxes {
		box := &w.boxes[i]
		if box.Intersects(pos) {
			boxes = append(boxes, box)
		}
	}
	return boxes
}

func (w *World) apply(vec mathy.Vec) {
	next := mathy.Rect{Vec: w.robot.Plus(vec), Width: 0, Height: 0}

	if w.intersectsWall(next) {
		// We can't move.
		return
	}

	if boxes := w.intersectingBoxes(next); boxes != nil {
		// There's boxes in the way, see if we can move them.
		blocked := true

		boxesToCheck := []*mathy.Rect{}
		boxesToCheck = append(boxesToCheck, boxes...)

		boxesToMove := map[*mathy.Rect]struct{}{}

		i := 0
		for {
			if i >= len(boxesToCheck) {
				// We found some free space!
				blocked = false
				break
			}

			box := boxesToCheck[i]
			i++

			boxesToMove[box] = struct{}{}

			tile := box.Translate(vec)

			if w.intersectsWall(tile) {
				// We can't move the boxes, give up.
				break
			}

			if boxes = w.intersectingBoxes(tile); boxes != nil {
				for _, b := range boxes {
					// We don't care if we intersect with ourselves.
					if b == box {
						continue
					}
					for _, seen := range boxesToCheck {
						if b == seen {
							continue
						}
					}
					// We found another box, queue it up.
					boxesToCheck = append(boxesToCheck, b)
				}
			}
		}

		if blocked {
			return
		}

		for box := range boxesToMove {
			box.Add(vec)
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

func part2(input string) int {
	world, moves := parse(input, true)

	for _, move := range moves {
		world.apply(move)
	}

	sum := 0
	for _, box := range world.boxes {
		sum += box.X + box.Y*100
	}
	return sum
}

func manual(input string) {
	world, _ := parse(input, true)

	world.debug("INITIAL STATE")

	for {
		move := mathy.Vec{}

		var key string
		fmt.Print("Direction?")
		fmt.Scanf("%s", &key)
		switch key {
		case ".":
			move.Y = -1
		case "e":
			move.Y = 1
		case "o":
			move.X = -1
		case "u":
			move.X = 1
		}

		world.apply(move)
		world.debug("NEW STATE")
	}
}

//go:embed puzzle.txt
var puzzle string

func main() {
	fmt.Println("1:", part1(puzzle))
	fmt.Println("2:", part2(puzzle))
	//manual(puzzle)
}
