package main

import (
	"fmt"
)

func main() {
	var n, m int
	var s string
	var t string
	fmt.Scan(&n, &m)
	fmt.Scan(&s)
	fmt.Scan(&t)

	result := 3

	if s == t[:n] && s == t[len(t)-n:] {
		result = 0
	} else if s == t[:n] {
		result = 1
	} else if s == t[len(t)-n:] {
		result = 2
	}
	fmt.Println(result)
}
