package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/albertb/advent-of-code/mathy"
)

type Opcode int

const (
	Adv Opcode = iota // Division: A = int(A / 2^combo)
	Bxl               // Bitwise XOR: B = B XOR literal
	Bst               // Modulo: B = combo % 8
	Jnz               // Jump: if A != 0 then jump to literal
	Bxc               // Bitwise XOR: B = B XOR C
	Out               // Modulo: print(combo % 8)
	Bdv               // Division: B = int(A / 2^combo)
	Cdv               // Division: C = int(A / 2^combo)
)

func (o Opcode) String() string {
	switch o {
	case Adv:
		return "Adv"
	case Bxl:
		return "Bxl"
	case Bst:
		return "Bst"
	case Jnz:
		return "Jnz"
	case Bxc:
		return "Bxc"
	case Out:
		return "Out"
	case Bdv:
		return "Bdv"
	case Cdv:
		return "Cdv"
	default:
		return "UNRECOGNIZED"
	}
}

type Computer struct {
	A, B, C int64 // Registers
	verbose bool
}

func (c Computer) Combo(combo int64) int64 {
	switch combo {
	case 0, 1, 2, 3:
		return combo
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	default:
		log.Fatalln("invalid combo operand:", combo)
	}
	return -1
}

func parse(input string) []int {
	var program []int
	for token := range strings.SplitSeq(input, ",") {
		number, err := strconv.ParseInt(token, 10, 16)
		if err != nil {
			log.Fatal("failed to parse program:", err)
		}
		program = append(program, int(number))
	}
	return program
}

func (c *Computer) Execute(program []int64) []int64 {
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("A:", c.A, "recovered from", r)
		}
	}()

	if c.verbose {
		fmt.Printf("A:%26b\nB:%26b\nC:%26b\n..\n", c.A, c.B, c.C)
	}

	var out []int64
	i := 0
	for {
		if i >= len(program) {
			break
		}

		opcode := Opcode(program[i])
		operand := int64(program[i+1])

		switch opcode {
		case Adv:
			if c.verbose {
				fmt.Println("A = A >> B")
			}
			c.A = c.A / mathy.Pow64(2, c.Combo(operand))
		case Bxl:
			if c.verbose {
				fmt.Printf("B = B XOR %b\n", operand)
			}
			c.B = c.B ^ operand
		case Bst:
			if c.verbose {
				fmt.Println("B = A & 111")
			}
			c.B = c.Combo(operand) % 8
		case Jnz:
			if c.A != 0 {
				if c.verbose {
					fmt.Printf("JUMP\n\n")
				}
				i = int(operand)
				continue
			}
		case Bxc:
			if c.verbose {
				fmt.Println("B = B ^ C")
			}
			c.B = c.B ^ c.C
		case Out:
			if c.verbose {
				fmt.Println("O = B & 111")
			}
			out = append(out, c.Combo(operand)%8)
		case Bdv:
			if c.verbose {
				fmt.Println("B = A >> combo")
			}
			c.B = c.A / mathy.Pow64(2, c.Combo(operand))
		case Cdv:
			if c.verbose {
				fmt.Println("C = A >> B")
			}
			c.C = c.A / mathy.Pow64(2, c.Combo(operand))
		}
		i += 2

		if c.verbose {
			opname := fmt.Sprintf("%d", operand)
			switch opname {
			case "4":
				opname = "A"
			case "5":
				opname = "B"
			case "6":
				opname = "C"
			}

			//fmt.Printf("%s %s\nA:%26b\nB:%26b\nC:%26b\n..\n", opcode.String(), opname, c.A, c.B, c.C)
			fmt.Printf("A:%26b\nB:%26b\nC:%26b\n..\n", c.A, c.B, c.C)
		}
	}

	return out //return strings.Join(out, ",")
}

const Program = "2,4,1,2,7,5,0,3,1,7,4,1,5,5,3,0"

func main() {

	want := []int64{2, 4, 1, 2, 7, 5, 0, 3, 1, 7, 4, 1, 5, 5, 3, 0}

	if len(os.Args) == 2 {
		a, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		c := Computer{A: a}
		out := c.Execute(want)
		fmt.Printf("out(%d) = %v\n", len(out), out)
		return
	}

	if len(os.Args) == 3 {
		//a := int64(190617786268687)
		a := int64(190617744915471)
		besta := a
		for i := range math.MaxInt {
			c := Computer{A: a}
			out := c.Execute(want)

			matches := 0
			for i := len(want) - 1; i >= 0; i-- {
				if out[i] == want[i] {
					matches++
				} else {
					break
				}
			}

			if matches == 16 {
				fmt.Println("FOUND A=", a)
				besta = a
			}

			if i%(2<<16) == 0 {
				fmt.Printf("A: %32b; matches: %2d; best A: %d\n", a, matches, besta)
			}

			a -= 16
		}
	}

	//a := 190617786419811

	//a := 190617786268687 //100000000000000
	//x := -1000000000

	//a := 190615638934425
	//x := -10000000
	//n := 24

	//a := int64(132184372088832)
	a := int64(35184372088832)
	x := int64(2 << 26)

	besta := int64(0)
	bestmatches := 0
	for i := range math.MaxInt {
		c := Computer{A: a}
		out := c.Execute(want)

		if i%1000000 == 0 {
			fmt.Printf("A: %22d / %16x x: %8x len: %d besta: %d bestmatches %d\n", a, a, x, len(out), besta, bestmatches)
		}

		if len(out) < 16 {
			if x < 0 {
				x = -x >> 1
			}
			a += x
			continue
		}

		if len(out) > 16 {
			if x > 0 {
				x = -x >> 1
			}
			a += x
			continue
		}

		matches := 0
		for i := len(want) - 1; i >= 0; i-- {
			if out[i] == want[i] {
				matches++
			} else {
				break
			}
		}

		if matches == 16 {
			fmt.Println("FOUND A =", a)
			break
		}

		if matches > bestmatches {
			fmt.Printf("A: %22d O: %v x: %d\n", a, out[len(out)-matches:], x)
			x = -x >> 1
			bestmatches = matches
			besta = a
		}
		a += x
	}

	/*best := math.MaxInt
	for i := range math.MaxInt {
		s := run(int64(a))
		if s == Program {
			fmt.Println("FOUND A=", a)
			if a < best {
				best = a
			}
		}
		if i%100000 == 0 {
			fmt.Println("a", a, "best", best, "s", s)
		}
		a += x
	}*/
}
