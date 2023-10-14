package main

import (
	"fmt"
)

func main() {
	var n int
	var t string
	fmt.Scan(&n, &t)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	var r []int
	for i, c := range s {
		if c == t {
			r = append(r, i+1)
			continue
		} else if term2(t, c) {
			r = append(r, i+1)
			continue
		} else if term3(t, c) {
			r = append(r, i+1)
			continue
		} else if term4(t, c) {
			r = append(r, i+1)
			continue
		}
	}

	fmt.Println(len(r))
	for _, c := range r {
		fmt.Print(c, " ")
	}
}

func term2(t string, s string) bool {
	r := false
	for i := 0; i < len(t); i++ {
		if RemoveNthCharacter(t, i+1) == s {
			r = true
			break
		}
	}
	if r {
		return true
	}
	return false
}

func term3(t string, s string) bool {
	r := false
	for i := 0; i < len(s); i++ {
		if RemoveNthCharacter(s, i+1) == t {
			r = true
			break
		}
	}
	if r {
		return true
	}
	return false
}

func term4(t string, s string) bool {
	r := 0
	l := 0
	if len(t) > len(s) {
		l = len(s)
	} else {
		l = len(t)
	}

	for i := 0; i < l; i++ {
		if string(t[i]) != string(s[i]) {
			r++
		}
		if r > 1 {
			break
		}
	}

	if r <= 1 {
		return true
	} else {
		return false
	}
}

func RemoveNthCharacter(s string, n int) string {
	runes := []rune(s)
	return string(runes[:n-1]) + string(runes[n:])
}
