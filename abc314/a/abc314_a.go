package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := "1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"
	fmt.Print("3.")
	for i := 0; i < n; i++ {
		fmt.Print(string(s[i]))
	}
}
