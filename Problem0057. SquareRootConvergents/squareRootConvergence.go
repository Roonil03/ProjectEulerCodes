package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(3)
	b := big.NewInt(2)
	res := 0
	for i := 2; i <= 1000; i++ {
		na := new(big.Int).Add(a, new(big.Int).Mul(b, big.NewInt(2)))
		nb := new(big.Int).Add(a, b)
		a, b = na, nb
		if len(a.String()) > len(b.String()) {
			res++
		}
	}
	fmt.Print("The answer is: ", res)
}
