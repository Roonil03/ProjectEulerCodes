package main

import (
	"fmt"
	"math/big"
)

func main() {
	mod := big.NewInt(1e10)
	exp := big.NewInt(7830457)
	two := big.NewInt(2)
	res := new(big.Int).Exp(two, exp, mod)
	res.Mul(res, big.NewInt(28433))
	res.Add(res, big.NewInt(1))
	res.Mod(res, mod)
	fmt.Println("The answer is:", res)
}
