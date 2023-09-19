package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	s := "1"

	for i := 1; i <= n; i++ {
		non := true
		for j := 1; j <= 9; j++ {
			if n%j == 0 {
				if i%(n/j) == 0 {
					s += strconv.Itoa(j)
					non = false
					break
				}
			}
		}
		if non {
			s += "-"
		}
	}
	fmt.Println(s)
}
