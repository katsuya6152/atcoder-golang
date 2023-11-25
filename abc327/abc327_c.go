package main

import (
	"fmt"
)

func main() {
	a := make([][]int, 9)
	for i := 0; i < 9; i++ {
		a[i] = make([]int, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	first := false
	for _, slice := range a {
		if contains(slice) {
			first = true
		}
	}

	second := false
	for col := 0; col < 9; col++ {
		var colSlice []int
		for row := 0; row < len(a); row++ {
			colSlice = append(colSlice, a[row][col])
		}
		if contains(colSlice) {
			second = true
		}
	}

	third := false
	for i := 0; i <= 6; i = i + 3 {
		for j := 0; j < 3; j++ {
			var slice []int
			for k := 0; k <= 6; k = k + 3 {
				tmp := a[j][k : k+3]
				for _, e := range tmp {
					slice = append(slice, e)
				}
			}
			fmt.Print(slice)
			fmt.Println(contains(slice))
		}
	}

	if !first || !second || !third {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

func contains(slice []int) bool {
	occurrences := make(map[int]bool)

	for _, num := range slice {
		if occurrences[num] {
			return false
		}
		occurrences[num] = true
	}

	for i := 1; i <= 9; i++ {
		if !occurrences[i] {
			return false
		}
	}

	return true
}
