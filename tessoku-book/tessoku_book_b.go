package main

import (
	"fmt"
)

func main() {
	var n, x int
	var a int
	result := "No"

	fmt.Scan(&n, &x)
	for i:=0; i<n; i++ {
		fmt.Scan(&a)
		if a == x {
			result = "Yes"
		}
	}

	fmt.Print(result)
}
