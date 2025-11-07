package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

type Block struct {
	file   int // -1 is free space
	length int // how long was this file/free space in the original order
}

func parse(input string) []Block {
	var blocks []Block
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}

		var file int64
		free := false
		for number := range strings.SplitSeq(line, "") {
			digit, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				log.Fatalln(err)
			}

			if free {
				for range digit {
					blocks = append(blocks, Block{-1, int(digit)})
				}
			} else {
				for range digit {
					blocks = append(blocks, Block{int(file), int(digit)})
				}
				file++
			}
			free = !free
		}
	}
	return blocks
}

func nextFree(from int, blocks []Block) int {
	for i := from; i < len(blocks); i++ {
		if blocks[i].file == -1 {
			return i
		}
	}
	return -1
}

func defrag(blocks []Block) int64 {
	// Find the first free block.
	free := nextFree(0, blocks)

	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i].file != -1 {
			// Swap this file block with the next free block.
			blocks[free].file = blocks[i].file
			blocks[i].file = -1

			// Find the next free block for the next file.
			free = nextFree(free, blocks)
		}
		if i <= free {
			break
		}
	}

	var sum int64
	for i, block := range blocks {
		if block.file == -1 {
			continue
		}
		sum += int64(i * block.file)
	}
	return sum
}

type Spans struct {
	start  int
	length int
}

func debug(prefix string, blocks []Block) {
	fmt.Printf("%-20s", prefix)
	for _, block := range blocks {
		if block.file == -1 {
			fmt.Printf(".")
		} else {
			fmt.Printf("%d", block.file)
		}
	}
	fmt.Println()
}

func defragContiguous(blocks []Block) int64 {
	//fmt.Println(">>> input:", blocks)

	//fmt.Printf("%-20s%-10s%-10s%-10s%-10s%-10s\n", "", "", "1", "2", "3", "4")
	//fmt.Printf("%-20s%s%s%s%s%s\n", "", "0123456789", "0123456789", "0123456789", "0123456789", "0123456789")
	debug("INPUT", blocks)

	free := []*Spans{}
	for i := 0; i < len(blocks); i++ {
		if blocks[i].file == -1 {
			// We're looknig at a block of free space.
			free = append(free, &Spans{i, blocks[i].length})
			i += int(blocks[i].length)
		}
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		//debug("", blocks)

		if blocks[i].file != -1 {
			// We're looking at a block of a file.
			file := blocks[i]
			//debug(fmt.Sprintf("look f%d @ %d (%d)", file.file, i, file.length), blocks)

			for _, span := range free {
				if span.start > (i - file.length) {
					// Don't move files backwards.
					//fmt.Println("out of free spans; not moving this file")
					//debug(fmt.Sprintf("cant mov f%d @ %d (%d)", file.file, i, file.length), blocks)
					break
				}

				if span.length >= file.length {
					//debug(fmt.Sprintf("mov f%d to i%d (%d)", file.file, span.start, span.length), blocks)

					// We found a big enough free span, move the file.
					for k := range file.length {
						blocks[span.start+k].file = file.file
						blocks[span.start+k].length = file.length
						blocks[i-k].file = -1
						blocks[i-k].length = file.length
						//debug(fmt.Sprintf("m %d to %d; c %d", i-k, span.start+k, i-k), blocks)
					}

					// Adjust the span start and length to account for the
					// file we just moved. (Length may go to zero)
					span.start += file.length
					span.length -= file.length

					//debug(fmt.Sprintf("span left %d", span.length), blocks)
					break
				}
			}
			i -= (file.length - 1)
		}
	}
	debug("OUTPUT", blocks)

	var sum int64
	for i, block := range blocks {
		if block.file == -1 {
			continue
		}
		sum += int64(i * block.file)
	}
	return sum
}

func part1(input string) int64 {
	blocks := parse(input)
	sum := defrag(blocks)
	return sum
}

func part2(input string) int64 {
	blocks := parse(input)
	sum := defragContiguous(blocks)
	return sum
}
