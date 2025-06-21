package main

import (
	"fmt"
	"math/big"
)

func main() {
	total := 0
	for n := 1; n <= 100; n++ {
		total += digitSumSqrt(n, 100)
	}
	fmt.Println("The answer is:", total)
}

func digitSumSqrt(n int, digits int) int {
	p := sq(n)
	if p*p == n {
		return 0
	}
	r := big.NewInt(int64(n))
	r.Sub(r, big.NewInt(int64(p*p)))
	sum := 0
	res := big.NewInt(int64(p))
	ten := big.NewInt(10)
	hundred := big.NewInt(100)
	for i := 0; i < digits; i++ {
		r.Mul(r, hundred)
		x := new(big.Int).Mul(res, big.NewInt(20))
		d := byte(9)
		for ; d > 0; d-- {
			t := new(big.Int).Add(x, big.NewInt(int64(d)))
			t.Mul(t, big.NewInt(int64(d)))
			if t.Cmp(r) <= 0 {
				break
			}
		}
		t := new(big.Int).Add(x, big.NewInt(int64(d)))
		t.Mul(t, big.NewInt(int64(d)))
		r.Sub(r, t)
		res.Mul(res, ten)
		res.Add(res, big.NewInt(int64(d)))
		sum += int(d)
	}
	return sum
}

func sq(n int) int {
	lo, hi := 0, n
	for lo <= hi {
		m := (lo + hi) / 2
		if m*m <= n {
			lo = m + 1
		} else {
			hi = m - 1
		}
	}
	return hi
}
