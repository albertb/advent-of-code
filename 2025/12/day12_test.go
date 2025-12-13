package main

import "testing"

var example = `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		/*"example": {
			input: example,
			want:  2,
		},*/
		"puzzle": {
			input: puzzle,
			want:  425,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := part1(tt.input), tt.want; got != want {
				t.Errorf("got part1() = %v, want = %v", got, want)
			}
		})
	}
}
