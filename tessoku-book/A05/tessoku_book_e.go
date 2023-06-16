package main

import (
	"fmt"
)

func main() {
	var n, k int
	var result int
	fmt.Scan(&n, &k)

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			l := k - i - j
			if l >= 1 && l <= n {
				result++
			}
		}
	}

	fmt.Print(result)
}
