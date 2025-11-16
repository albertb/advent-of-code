package main

import (
	"slices"
	"testing"
)

func Test_simple(t *testing.T) {
	c := Computer{C: 9}
	c.Execute([]int64{2, 6})
	if got, want := c.B, int64(1); got != want {
		t.Errorf("got B = %v, want %v", got, want)
	}

	c = Computer{A: 10}
	if got, want := c.Execute([]int64{5, 0, 5, 1, 5, 4}), []int64{0, 1, 2}; !slices.Equal(got, want) {
		t.Errorf("got out = %v, want %v", got, want)
	}

	c = Computer{A: 2024}
	if got, want := c.Execute([]int64{0, 1, 5, 4, 3, 0}), []int64{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}; !slices.Equal(got, want) {
		t.Errorf("got out = %v, want %v", got, want)
	}
	if got, want := c.A, int64(0); got != want {
		t.Errorf("got A = %v, want %v", got, want)
	}

	c = Computer{B: 29}
	c.Execute([]int64{1, 7})
	if got, want := c.B, int64(26); got != want {
		t.Errorf("got B = %v, want %v", got, want)
	}

	c = Computer{B: 2024, C: 43690}
	c.Execute([]int64{4, 0})
	if got, want := c.B, int64(44354); got != want {
		t.Errorf("got B = %v, want %v", got, want)
	}

	c = Computer{A: 729}
	if got, want := c.Execute([]int64{0, 1, 5, 4, 3, 0}), []int64{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}; !slices.Equal(got, want) {
		t.Errorf("got out = %v, want %v", got, want)
	}
}
