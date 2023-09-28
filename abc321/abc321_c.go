package main

import (
	"fmt"
)

func findValueAtIndex(index int) int {
	return index
}

func main() {
	var k int
	fmt.Scanf("%d\n", &k)

	value := findValueAtIndex(k)
	fmt.Println(value)
}
