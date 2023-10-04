package main

import (
	"fmt"
)

func main() {
	var n int
	var x int
	fmt.Scan(&n)

	c := make([]int, n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
		a[i] = make([]int, c[i])
		for j := 0; j < c[i]; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	fmt.Scan(&x)

	existI := make([]int, 0)
	existLen := make([]int, 0)
	for i := 0; i < n; i++ {
		for _, e := range a[i] {
			if e == x {
				existI = append(existI, i)
				existLen = append(existLen, len(a[i]))
			}
		}
	}

	if len(existLen) == 0 {
		fmt.Print(0)
	} else {
		minLenIs := FindMinIndices(existLen)
		resultN := len(minLenIs)
		fmt.Println(resultN)

		for _, e := range minLenIs {
			fmt.Print(existI[e]+1, " ")
		}
	}
}

func FindMin(nums []int) int {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return min
}

func FilterValue(nums []int, value int) []int {
	var result []int
	for _, num := range nums {
		if num == value {
			result = append(result, num)
		}
	}
	return result
}

func FindMinIndices(nums []int) []int {
	min := nums[0]
	indices := []int{0}

	for i, num := range nums[1:] {
		if num < min {
			min = num
			indices = []int{i + 1}
		} else if num == min {
			indices = append(indices, i+1)
		}
	}
	return indices
}
