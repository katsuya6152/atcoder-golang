package main

import (
	"fmt"
)

func main() {
	var h, w, n int
	fmt.Scan(&h, &w, &n)
	algo := make([][]int, h+2)

	for i := 0; i <= h; i++ {
		algo[i] = make([]int, w+2)
	}

	for i := 0; i < n; i++ {
		var a, b, c, d int
		fmt.Scan(&a, &b, &c, &d)
		algo[a][b]++
		algo[a][d+1]--
		algo[c+1][b]--
		algo[c+1][d+1]++
	}
	fmt.Println(algo)

	for i := 0; i <= w; i++ {
		for j := 1; j <= h; j++ {
			algo[j][i] += algo[j-1][i]
		}
	}

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			algo[i][j] += algo[i][j-1]
			fmt.Print(algo[i][j])
			if j != w {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
