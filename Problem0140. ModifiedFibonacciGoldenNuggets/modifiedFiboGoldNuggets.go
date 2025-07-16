package main

import (
	"fmt"
	"math/big"
)

func fib(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	a := big.NewInt(0)
	b := big.NewInt(1)
	for i := 1; i < n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}

func main() {
	sum := big.NewInt(0)
	for i := 1; i <= 30; i++ {
		f1 := fib(2 * i)
		f2 := fib(2*i + 1)
		prod := new(big.Int).Mul(f1, f2)
		sum.Add(sum, prod)
	}
	fmt.Println("The answer is:", sum)
}
