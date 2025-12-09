package main

import "testing"

var example = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input       string
		connections int
		want        int
	}{
		"example": {
			input:       example,
			connections: 10,
			want:        40,
		},
		"puzzle": {
			input:       puzzle,
			connections: 1000,
			want:        135169,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.input, tt.connections), tt.want; got != want {
				t.Errorf("part1() = %v, want = %v", got, want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: example,
			want:  25272,
		},
		"puzzle": {
			input: puzzle,
			want:  302133440,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part2(tt.input), tt.want; got != want {
				t.Errorf("part2() = %v, want = %v", got, want)
			}
		})
	}
}
