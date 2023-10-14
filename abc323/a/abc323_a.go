package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanf("%s\n", &s)

	isOne := false

	for i, c := range s {
		if (i+1)%2 == 0 {
			if string(c) == "1" {
				isOne = true
			}
		}
	}
	if isOne {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}

}
