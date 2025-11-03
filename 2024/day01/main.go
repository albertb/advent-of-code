package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("one.txt")
	if err != nil {
		log.Fatalln("failed to open file:", err)
	}
	defer file.Close()

	left := []int{}
	right := []int{}

	scanner := bufio.NewScanner(file)
	seperator := "   "

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, seperator, 2)
		if len(parts) != 2 {
			log.Fatalln("invalid line format:", line)
		}

		parseFn := func(s string) (int, error) {
			num, err := strconv.Atoi(s)
			if err != nil {
				return 0, err
			}
			return num, nil
		}

		leftNum, err := parseFn(parts[0])
		if err != nil {
			log.Fatalln("failed to convert left part to integer:", err)
		}
		left = append(left, leftNum)

		rightRight, err := parseFn(parts[1])
		if err != nil {
			log.Fatalln("failed to convert right part to integer:", err)
		}
		right = append(right, rightRight)
	}

	sort.IntSlice(left).Sort()
	sort.IntSlice(right).Sort()

	if len(left) != len(right) {
		log.Fatalln("left and right slices have different lengths")
	}

	rightCount := make(map[int]int64)

	var distance int64 = 0
	for i := range left {
		//log.Println("left is", left[i], "right is", right[i])
		distance += int64(math.Abs(float64(left[i] - right[i])))
		//log.Println("distance so far:", distance)
		rightCount[right[i]]++
	}
	log.Println("Total distance:", distance)

	var similarity int64 = 0
	for i := range left {
		//log.Println("left is", left[i], "right count is", rightCount[left[i]])
		similarity += int64(left[i]) * rightCount[left[i]]
		//log.Println("similarity so far:", similarity)
	}
	log.Println("Total similarity:", similarity)
}
