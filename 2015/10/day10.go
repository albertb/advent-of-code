package main

import (
	"strconv"
	"strings"
)

func solve(s string, n int) int {
	for range n {
		var sb strings.Builder

		prev := s[0]
		count := 1
		for i := 1; i < len(s); i++ {
			this := s[i]
			if prev == this {
				count++
			} else {
				sb.WriteString(strconv.Itoa(count))
				sb.WriteByte(prev)
				count = 1
			}
			prev = this
		}
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(prev)

		s = sb.String()
	}
	return len(s)
}
