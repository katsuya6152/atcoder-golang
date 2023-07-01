package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	l := 1
	r := 1000000000
	for l < r {
		m := (l + r) / 2
		sum := 0
		for i := 1; i <= n; i++ {
			sum += m / a[i]
		}
		if sum < k {
			l = m + 1
		}
		if sum >= k {
			r = m
		}
	}
	fmt.Print(l)
}
