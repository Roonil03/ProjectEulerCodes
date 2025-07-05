package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := binom(109, 9)
	b := binom(110, 10)
	sum := new(big.Int).Add(a, b)
	sub := big.NewInt(10*100 + 2)
	ans := sum.Sub(sum, sub)
	fmt.Println("The answer is:", ans)
}

func binom(n, k int64) *big.Int {
	if k < 0 || k > n {
		return big.NewInt(0)
	}
	if k > n-k {
		k = n - k
	}
	res := big.NewInt(1)
	for i := int64(1); i <= k; i++ {
		res.Mul(res, big.NewInt(n-k+i))
		res.Div(res, big.NewInt(i))
	}
	return res
}
