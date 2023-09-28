package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	strN := strconv.Itoa(n)
	prev := 10000
	isYes := true

	if len(strN) != 1 {
		for i := 0; i < len(strN); i++ {
			cur := int(strN[i] - '0')

			if cur >= prev {
				isYes = false
				break
			}
			prev = cur
		}
	}

	if isYes {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
