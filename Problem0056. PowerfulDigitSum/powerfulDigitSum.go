package main

import (
	"fmt"
	"math/big"
)

func main() {
	res := 0
	b := new(big.Int)
	p := new(big.Int)
	for a := 1; a < 100; a++ {
		b.SetInt64(int64(a))
		p.SetInt64(1)
		for c := 1; c < 100; c++ {
			p.Mul(p, b)
			if ds := dig(p); ds > res {
				res = ds
			}
		}
	}
	fmt.Print("The answer is: ", res)
}

func dig(n *big.Int) int {
	s := 0
	for _, c := range n.String() {
		s += int(c - '0')
	}
	return s
}
