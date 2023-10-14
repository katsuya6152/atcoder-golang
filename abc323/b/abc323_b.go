package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	s := make([]string, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
		for _, c := range s[i] {
			if string(c) == "o" {
				r[i]++
			}
		}
	}

	result := argsortDesc(r)
	for _, c := range result {
		fmt.Print(c+1, " ")
	}
}

func argsortDesc(arr []int) []int {
	type Pair struct {
		Value int
		Index int
	}
	pairs := make([]Pair, len(arr))
	for i, v := range arr {
		pairs[i] = Pair{v, i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Value == pairs[j].Value {
			return pairs[i].Index < pairs[j].Index
		}
		return pairs[i].Value > pairs[j].Value
	})

	sortedIndices := make([]int, len(arr))
	for i, pair := range pairs {
		sortedIndices[i] = pair.Index
	}
	return sortedIndices
}
