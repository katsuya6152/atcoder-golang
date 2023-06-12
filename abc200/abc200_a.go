package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	if n%100 == 0 {
		fmt.Print(n / 100)
	} else {
		fmt.Print(n/100 + 1)
	}
}
