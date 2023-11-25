package main

import (
	"fmt"
	"math/big"
)

func main() {
	var b int64
	fmt.Scan(&b)
	bBig := big.NewInt(b)
	r := false
	for i := int64(1); i <= 1000; i++ {
		iBig := big.NewInt(i)
		pow := new(big.Int).Exp(iBig, iBig, nil)
		if pow.Cmp(bBig) == 0 {
			r = true
			fmt.Println(i)
			break
		}
	}
	if !r {
		fmt.Println(-1)
	}
}
