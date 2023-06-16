package main

import (
	"fmt"
	"sort"
)

func main() {

	a := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&a[i])
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	if a[2]-a[1] == a[1]-a[0] {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
