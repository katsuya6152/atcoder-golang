package main

import (
	"fmt"
)

func check(s string) bool {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	rStr := string(r)

	return s == rStr
}

func main() {
	var s string
	fmt.Scan(&s)
	result := 1
	for i := len(s); i >= 0; i-- {
		for j := 0; j < i; j++ {
			k := s[j:i]
			if check(k) {
				if result < len(k) {
					result = len(k)
				}
			}
		}

	}
	fmt.Println(result)
}
