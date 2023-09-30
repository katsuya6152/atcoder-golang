package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s\n", &s)

	result := strings.Index(s, "ABC")
	if result == -1 {
		fmt.Println(result)
	} else {
		fmt.Println(result + 1)
	}
}
