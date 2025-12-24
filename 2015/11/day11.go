package main

import "strings"

func solve(input string) string {
	s := []rune(input)

generate:
	for {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == 'z' {
				s[i] = 'a'
			} else {
				s[i]++
				break
			}
		}
		password := string(s)

		// Password may not contain the letters i, o or l.
		if strings.ContainsAny(password, "iol") {
			continue
		}

		// Password must contain at least two different, non-overlapping pairs of letters.
		var pairs int
		for i := 1; i < len(s); i++ {
			if s[i-1] == s[i] {
				pairs++
				i++
			}
		}
		if pairs < 2 {
			continue
		}

		// Password must include one increasing straight of at least three letters.
		for i := 2; i < len(s); i++ {
			if s[i-2]+1 == s[i-1] && s[i-1]+1 == s[i] {
				break generate
			}
		}
		continue
	}
	return string(s)
}
