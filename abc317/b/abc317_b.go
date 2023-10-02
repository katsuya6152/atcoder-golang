package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	for i := 0; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	next := 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			next = a[i] + 1
		} else {
			if a[i] != next {
				fmt.Println(next)
				break
			}
			next++
		}
	}
}
