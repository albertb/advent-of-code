package main

import "testing"

func Test_FirstDirectionPad(t *testing.T) {
	keypad := newKeypad()
	keypadMoves := enter(keypad, []rune("0"))

	directionPad := newDirectionPad()
	directionPadMoves := enter(directionPad, keypadMoves)

	if got, want := string(directionPadMoves), "v<<A>>^A"; got != want {
		t.Errorf("moves = %v, want %v", got, want)
	}
}

func Test_PreferRight(t *testing.T) {
	directionPad := newDirectionPad()
	directionPad1Moves := enter(directionPad, []rune("^^>"))
	directionPad2Moves := enter(directionPad, directionPad1Moves)

	if got, want := string(directionPad2Moves), "v<<A>>^AAvA<A"; got != want {
		t.Errorf("moves = %v, want %v", got, want)
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
