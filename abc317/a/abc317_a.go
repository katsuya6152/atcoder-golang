package main

import (
	"fmt"
)

func main() {
	var n, h, x int
	fmt.Scan(&n, &h, &x)

	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&p[i])
	}

	for i := 1; i <= n; i++ {
		if p[i]+h >= x {
			fmt.Println(i)
			break
		}
	}
}
