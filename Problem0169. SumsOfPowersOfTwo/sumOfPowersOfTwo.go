package main

import (
	"fmt"
	"math/big"
)

var (
	one = big.NewInt(1)
	mem = map[string]*big.Int{"0": one}
)

func f(n *big.Int) *big.Int {
	s := n.String()
	if v, ok := mem[s]; ok {
		return v
	}
	k := new(big.Int).Rsh(n, 1)
	if n.Bit(0) == 1 {
		mem[s] = f(k)
	} else {
		mem[s] = new(big.Int).Add(f(k), f(new(big.Int).Sub(k, one)))
	}
	return mem[s]
}

func main() {
	n, _ := new(big.Int).SetString("10000000000000000000000000", 10)
	fmt.Println(f(n))
}
