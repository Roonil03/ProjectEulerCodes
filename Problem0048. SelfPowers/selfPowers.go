package main

import (
	"fmt"
	"math/big"
)

func main() {
	const mod = 10000000000
	total := big.NewInt(0)
	modulus := big.NewInt(mod)
	for i := 1; i <= 1000; i++ {
		b := big.NewInt(int64(i))
		term := new(big.Int).Exp(b, b, modulus)
		total.Add(total, term)
		total.Mod(total, modulus)
	}
	fmt.Printf("The answer is: %010s\n", total.String())
}
