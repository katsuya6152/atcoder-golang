package main

import (
	"fmt"
)

func main() {

	var d int
	var n int
	fmt.Scan(&d, &n)

	b := make([]int, d+2)
	result := 0

	for i := 1; i <= n; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		b[l]++
		b[r+1]--
	}

	for i := 1; i <= d; i++ {
		result += b[i]
		fmt.Println(result)
	}
}
