package main

import (
	"fmt"
)

func main() {

	var h, w int
	fmt.Scan(&h, &w)

	x := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		x[i] = make([]int, w+1)
		if i != 0 {
			for j := 1; j <= w; j++ {
				fmt.Scan(&x[i][j])
			}
		}
	}

	s := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		s[i] = make([]int, w+1)
		if i != 0 {
			for j := 1; j <= w; j++ {
				s[i][j] = s[i][j-1] + x[i][j]
			}
		}
	}

	for i := 0; i <= w; i++ {
		for j := 1; j <= h; j++ {
			if i != 0 {
				s[j][i] += s[j-1][i]
			}
		}
	}

	var q int
	fmt.Scan(&q)

	for i := 1; i <= q; i++ {
		var a, b, c, d int
		fmt.Scan(&a, &b, &c, &d)
		fmt.Println(s[c][d] - s[c][b-1] - s[a-1][d] + s[a-1][b-1])
	}
}
