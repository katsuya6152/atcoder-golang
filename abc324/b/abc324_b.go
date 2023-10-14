package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	c3 := 0
	c2 := 0
	cur := n

	for cur%3 == 0 && cur != 1 {
		c3++
		cur = cur / 3
	}
	for cur%2 == 0 && cur != 1 {
		c2++
		cur = cur / 2
	}
	if cur == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
