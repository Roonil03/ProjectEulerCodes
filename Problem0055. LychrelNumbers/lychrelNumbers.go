package main

import (
	"fmt"
	"math/big"
)

var (
	ten  = big.NewInt(10)
	zero = big.NewInt(0)
)

func main() {
	res := 0
	for i := 1; i < 10000; i++ {
		if check(i) {
			res++
		}
	}
	fmt.Print("The answer is: ", res)
}

func check(x int) bool {
	n := big.NewInt(int64(x))
	for i := 0; i < 50; i++ {
		n.Add(n, reverse(n))
		if palin(n) {
			return false
		}
	}
	return true
}

func palin(n *big.Int) bool {
	return n.Cmp(reverse(n)) == 0
}

func reverse(n *big.Int) *big.Int {
	r := new(big.Int)
	t := new(big.Int).Set(n)
	mod := new(big.Int)
	for t.Cmp(zero) > 0 {
		t.DivMod(t, ten, mod)
		r.Mul(r, ten)
		r.Add(r, mod)
	}
	return r
}
