package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	w := make([]int, n)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&w[i])
		fmt.Scan(&x[i])
	}

	max := 0
	for i := 0; i < 24; i++ {
		x_ := getLocalTime(x, i)
		xi := getIndex(x_)
		total := getTotalW(xi, w)
		if total > max {
			max = total
		}
	}

	fmt.Println(max)
}

func getLocalTime(x []int, i int) []int {
	rx := make([]int, len(x))
	copy(rx, x)
	for ri := range rx {
		rx[ri] += i
		if rx[ri] >= 24 {
			rx[ri] = rx[ri] - 24
		}
	}
	return rx
}

func getIndex(x []int) []int {
	var r []int
	for i, c := range x {
		if c >= 9 && c < 18 {
			r = append(r, i)
		}
	}
	return r
}

func getTotalW(i []int, w []int) int {
	t := 0
	for _, ii := range i {
		t += w[ii]
	}
	return t
}
