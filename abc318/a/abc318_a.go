package main

import (
	"fmt"
)

func main() {
	var n, m, p int
	fmt.Scan(&n, &m, &p)
	result := 0

	for i := 0; i <= (n-m)/p; i++ {
		if i*p+m <= n {
			result++
		}
	}
	fmt.Println(result)
}
