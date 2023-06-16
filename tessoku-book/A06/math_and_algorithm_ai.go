package main

import (
	"fmt"
)

func main() {

	var n, q int
	fmt.Scan(&n, &q)
	a := make([]int, n)
	s := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if i == 0 {
			s[0] = a[0]
		} else {
			s[i] = s[i-1] + a[i]
		}
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		if l == 1 {
			fmt.Println(s[r-1])
		} else if l == 1 && r == 1 {
			fmt.Println(s[0])
		} else {
			fmt.Println(s[r-1] - s[l-2])
		}
	}
}
