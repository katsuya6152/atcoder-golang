package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a int
	r := big.NewInt(0)
	bigOne := big.NewInt(1)
	end := big.NewInt(64)

	for i := big.NewInt(0); i.Cmp(end) == -1; i.Add(i, bigOne) {
		fmt.Scan(&a)
		if a != 0 {
			bigTwo := big.NewInt(2)
			r = r.Add(r, bigTwo.Exp(bigTwo, i, nil))
		}
	}

	fmt.Print(r)
}
