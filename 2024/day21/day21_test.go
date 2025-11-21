package main

import (
	"fmt"
	"testing"
)

func Test_PadPaths(t *testing.T) {

	numPad := newKeypad()
	numKeys := []string{"A", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	fmt.Print("   ")
	for _, from := range numKeys {
		fmt.Printf("%8s", from)
	}
	fmt.Println()

	for _, from := range numKeys {
		fmt.Printf("%s: ", from)
		for _, to := range numKeys {
			fmt.Printf("%8s", numPad.paths[Path{from, to}])
		}
		fmt.Println()
	}

	dirPad := newDirectionPad()
	dirKeys := []string{"A", "^", "<", "v", ">"}

	fmt.Print("   ")
	for _, from := range dirKeys {
		fmt.Printf("%6s", from)
	}
	fmt.Println()

	for _, from := range dirKeys {
		fmt.Printf("%s: ", from)
		for _, to := range dirKeys {
			fmt.Printf("%6s", dirPad.paths[Path{from, to}])
		}
		fmt.Println()
	}
}

func Test_FirstDirectionPad(t *testing.T) {
	keypad := newKeypad()
	numSeq := sequences("0", keypad)

	dirpad := newDirectionPad()
	dirSeq := sequences(numSeq, dirpad)

	if got, want := string(dirSeq), "v<<A>>^A"; got != want {
		t.Errorf("moves = %v, want %v", got, want)
	}
}

func Test_379(t *testing.T) {
	keypad := newKeypad()
	dirpad := newDirectionPad()

	code := "379A"
	numSeq := sequences(code, keypad)
	dir1Seq := sequences(numSeq, dirpad)
	dir2Seq := sequences(dir1Seq, dirpad)

	fmt.Printf("code: %s\nnum: %s\ndir1: %s\ndir2: %s\n", code, numSeq, dir1Seq, dir2Seq)

	if got, want := len(dir2Seq), 64; got != want {
		t.Errorf("len = %v, want %v", got, want)
	}
}

func Test_Example(t *testing.T) {
	var input = `
029A
980A
179A
456A
379A`

	if got, want := part1(input), 126384; got != want {
		t.Errorf("sum = %v, want %v", got, want)
	}
}
