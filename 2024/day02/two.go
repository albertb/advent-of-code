package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("two.txt")
	if err != nil {
		log.Fatalln("failed to open file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	xxx := 0
	safe := 0
	for scanner.Scan() {
		if xxx > 1 {
			//break
		}
		xxx++

		line := scanner.Text()
		reports := strings.Split(line, " ")

		levels := []int{}
		for _, report := range reports {
			level, err := strconv.ParseInt(report, 10, 16)
			if err != nil {
				log.Fatalln("failed to parse level:", err)
			}
			levels = append(levels, int(level))
		}

		var isSafe bool
		for j := range len(levels) + 1 {
			isSafe = true
			var lastLevel *int
			var lastDiff *float64

			for i, level := range levels {
				if (j - 1) == i {
					log.Println("Skipping", level, "in", levels)
					continue
				}

				if lastLevel == nil {
					lastLevel = &level
					continue
				}
				diff := float64(*lastLevel - level)

				absDiff := math.Abs(diff)
				if absDiff < 1.0 || absDiff > 3.0 {
					log.Println("UNSAFE jump", level, "in", levels)
					isSafe = false
					break
				}

				if lastDiff == nil {
					lastLevel = &level
					lastDiff = &diff
					continue
				}

				if (diff >= 0.0) != (*lastDiff >= 0.0) {
					log.Println("UNSAFE sign", level, "in", levels)
					isSafe = false
					break
				}

				lastLevel = &level
				lastDiff = &diff
			}
			if isSafe {
				if j == 0 {
					log.Println("SAFE", levels)
				} else {
					log.Println("SAFE", levels, "skip", levels[j-1])
				}

				break
			}
		}

		if isSafe {
			safe++
		}
	}
	log.Println("Safe:", safe)
}
