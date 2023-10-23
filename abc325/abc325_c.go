package main

import (
	"fmt"
	"math"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	s := make([][]string, h)
	for i := 0; i < h; i++ {
		s[i] = make([]string, w)
	}

	px, py := 0, 0
	var rs [][][]int
	rs = append(rs, [][]int{})
	ci := 0
	f := false
	result := 0
	for i := 0; i < h; i++ {
		var t string
		fmt.Scan(&t)
		r := []rune(t)
		for j := 0; j < w; j++ {
			s[i][j] = string(r[j])
			// fmt.Println("for")
			if string(r[j]) == "#" {
				// fmt.Println("if1")
				if isLinked(px, py, i, j) {
					// fmt.Println("if2")
					c := []int{i, j}
					fmt.Println(rs)
					fmt.Println(c)
					rs[ci] = append(rs[ci], c)
					fmt.Println(rs)
					px, py = i, j
				} else {
					// fmt.Println("else2")
					if f {
						result++
					}
				}
			}
		}
	}

	fmt.Println(f, result)
}

func isLinked(x1 int, y1 int, x2 int, y2 int) bool {
	b := false
	r := math.Max(math.Abs(float64(x1)-float64(x2)), math.Abs(float64(y1)-float64(y2)))
	if int(r) == 1 {
		b = true
	}
	return b
}
