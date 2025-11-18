package main

import "fmt"

func main() {
	q := []string{
		"allo",
		"les",
		"amis",
	}

	for len(q) > 0 {
		lenq := len(q)
		fmt.Println(q[lenq-1])
		q = q[:lenq-1]
	}
}
