package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	a := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	l := 0
	r := n

	for l <= r {
		m := (l + r) / 2
		if a[m] < x {
			l = m + 1
		}
		if a[m] > x {
			r = m - 1
		}
		if a[m] == x {
			fmt.Print(m)
			break
		}
	}
}
