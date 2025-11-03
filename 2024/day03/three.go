package main

import (
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("three.txt")
	if err != nil {
		log.Fatalln("failed to open file:", err)
	}

	contents, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	re := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)")
	matches := re.FindAllStringSubmatch(string(contents), -1)

	do := true
	sum := 0
	for _, match := range matches {
		log.Println(match[0], match[1], match[2])

		if match[0] == "do()" {
			do = true
			continue
		}
		if match[0] == "don't()" {
			do = false
			continue
		}
		if !do {
			continue
		}

		lhs, err := strconv.ParseInt(match[1], 10, 16)
		if err != nil {
			log.Fatalln("failed to parse int:", err)
		}
		rhs, err := strconv.ParseInt(match[2], 10, 16)
		if err != nil {
			log.Fatalln("failed to parse int:", err)
		}
		sum += int(lhs) * int(rhs)
	}
	log.Println("Sum:", sum)
}
