package main

import (
	"testing"

	"github.com/albertb/advent-of-code/mathy"
)

func Test_heurisitc(t *testing.T) {
	t.Skip()
	tests := []struct {
		name string
		pos  mathy.Vec
		dir  mathy.Vec
		end  mathy.Vec
		want int
	}{
		{
			name: "ahead",
			pos:  mathy.Vec{X: 0, Y: 0},
			dir:  mathy.Vec{X: 1, Y: 0},
			end:  mathy.Vec{X: 1, Y: 0},
			want: 1,
		},
		{
			name: "left",
			pos:  mathy.Vec{X: 0, Y: 0},
			dir:  mathy.Vec{X: 1, Y: 0},
			end:  mathy.Vec{X: 0, Y: -1},
			want: 1001,
		},
		{
			name: "right",
			pos:  mathy.Vec{X: 0, Y: 0},
			dir:  mathy.Vec{X: 1, Y: 0},
			end:  mathy.Vec{X: 0, Y: 1},
			want: 1001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Reindeer{
				pos: tt.pos,
				dir: tt.dir,
			}
			m := Maze{
				end: tt.end,
			}
			if got := heuristic(r, m); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "first",
			input: `
###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`,
			want: 7036,
		},
		{
			name: "second",
			input: `
#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`,
			want: 11048,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
