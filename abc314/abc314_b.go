package main

import (
	"fmt"
)

func main() {
	var n int
	var x int
	fmt.Scan(&n)

	c := make([]int, n+1)
	a := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&c[i])
		a[i] = make([]int, c[i]+1)
		for j := 1; j <= c[i]; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	fmt.Scan(&x)

	fmt.Print(n, c, a, x)
}
