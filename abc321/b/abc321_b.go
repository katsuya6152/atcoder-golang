package main

import (
	"fmt"
	"sort"
)

func comp(a []int, x int) bool {
	sort.Ints(a)
	a_ := a[1 : len(a)-1]
	sum := 0
	for _, num := range a_ {
		sum += num
	}
	if sum >= x {
		return true
	} else {
		return false
	}

}

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	a := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&a[i])
	}

	result := -1
	for i:=0; i<=100; i++ {
		append := append(a, i)
		if comp(append, x) {
			result = i
			break
		}
	}
	fmt.Println(result)
}
