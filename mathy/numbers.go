package mathy

import (
	"fmt"
	"math"
	"strconv"
)

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Pow64(x, y int64) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func MustParseInt(s string) int {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprint("failed to parse", s, ":", err))
	}
	return int(num)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
