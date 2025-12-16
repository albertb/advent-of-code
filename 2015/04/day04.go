package main

import (
	"crypto/md5"
	"strconv"
)

func part1(input string) int {
	for n := 0; ; n++ {
		data := input + strconv.Itoa(n)
		sum := md5.Sum([]byte(data))
		// Five leading zeroes when converted to hex.
		if sum[0] == 0 && sum[1] == 0 && sum[2]>>4 == 0 {
			return n
		}
	}
}

func part2(input string) int {
	for n := 0; ; n++ {
		data := input + strconv.Itoa(n)
		sum := md5.Sum([]byte(data))
		// Six leading zeroes when converted to hex.
		if sum[0] == 0 && sum[1] == 0 && sum[2] == 0 {
			return n
		}
	}
}
