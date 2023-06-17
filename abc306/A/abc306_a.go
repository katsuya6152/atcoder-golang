package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	for _, c := range s {
		fmt.Printf("%c", c)
		fmt.Printf("%c", c)
	}
}
