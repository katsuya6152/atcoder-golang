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
	if strings.Contains(s, "ab") || strings.Contains(s, "ba") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
