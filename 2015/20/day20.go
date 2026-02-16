package main

func part1(input int) int {
	for i := 1; ; i++ {
		var total int
		if i == 1 {
			total = 10
		} else {
			for j := 1; j*j <= i; j++ {
				if i%j == 0 {
					total += j * 10
					total += (i / j) * 10
				}
			}
		}

		if total >= input {
			return i
		}
	}
}

func part2(input int) int {
	for i := 1; ; i++ {
		var total int

		if i == 1 {
			total = 11
		} else {
			for j := 1; j*j <= i; j++ {
				if i%j != 0 {
					continue
				}
				if j*50 >= i {
					total += j * 11
				}
				if (i/j)*50 >= i {
					total += (i / j) * 11
				}
			}
		}

		if total >= input {
			return i
		}
	}
}
