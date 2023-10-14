package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	a := make([]int, n)

	r := true
	p := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if i == 0 {
			p = a[i]
		} else {
			if p != a[i] {
				r = false
				break
			}
		}
	}

	if r {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
