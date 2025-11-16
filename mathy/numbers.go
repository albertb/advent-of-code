package mathy

import "math"

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Pow(x, y int64) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}
