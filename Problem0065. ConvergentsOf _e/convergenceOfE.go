package main

import (
	"fmt"
	"math/big"
)

func main() {
	p_prev := big.NewInt(1)
	p_curr := big.NewInt(2)
	for n := 1; n < 100; n++ {
		var a_n int
		if (n+1)%3 == 0 {
			a_n = 2 * ((n + 1) / 3)
		} else {
			a_n = 1
		}
		temp := new(big.Int).Set(p_curr)
		p_curr.Mul(p_curr, big.NewInt(int64(a_n)))
		p_curr.Add(p_curr, p_prev)
		p_prev.Set(temp)
	}
	res := 0
	numStr := p_curr.String()
	for _, digit := range numStr {
		res += int(digit - '0')
	}
	fmt.Println("The answer is:", res)
}
