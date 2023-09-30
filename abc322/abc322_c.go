package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, m+1)
	b := make([]int, m)
	c := 0
	for i := 1; i <= m; i++ {
		fmt.Scan(&a[i])
		b[i-1] = a[i] - c - 1
		c = a[i]
	}
	for i := 0; i < m; i++ {
		for j := b[i]; j >= 0; j-- {
			fmt.Println(j)
		}

	}
}
