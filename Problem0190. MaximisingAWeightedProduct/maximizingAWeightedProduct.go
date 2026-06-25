package main

import (
	"fmt"
	"math/big"
)

func main() {
	res := big.NewInt(0)
	for i := int64(2); i <= 15; i++ {
		t := i * (i + 1) / 2
		n := big.NewInt(2)
		n.Exp(n, big.NewInt(t), nil)
		for j := int64(1); j <= i; j++ {
			p := big.NewInt(j)
			p.Exp(p, p, nil)
			n.Mul(n, p)
		}
		d := big.NewInt(i + 1)
		d.Exp(d, big.NewInt(t), nil)
		n.Div(n, d)
		res.Add(res, n)
	}
	fmt.Println(res)
}
