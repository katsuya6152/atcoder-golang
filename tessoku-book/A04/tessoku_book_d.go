package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 9; i >= 0; i-- {
		v := n / (1 << i)
		fmt.Print(v % 2)
	}
}
