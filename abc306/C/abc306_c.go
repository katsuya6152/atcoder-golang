package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n*3)

	for i := 0; i < n*3; i++ {
		fmt.Scan(&a[i])
	}

	count := 0
	f := make([]int, n+1)
	for i:=0; i<n; i++ {
		for j:=0; j<n*3; j++ {
			if a[j] == i {
				count++
				if count == 2 {
					f[i] = j
					break
				}
			}
		}
	}


	sort_f := f
	sort.Ints(sort_f)
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			if sort_f[i] == f[j] {
				print(j)
			}
		}
	}
}
