package main

import (
	"fmt"
	"math/bits"
)

func mulMod(a, b, m uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	_, rem := bits.Div64(hi, lo, m)
	return rem
}

func powMod(base, exp, mod uint64) uint64 {
	res := uint64(1)
	for exp > 0 {
		if exp&1 == 1 {
			res = mulMod(res, base, mod)
		}
		base = mulMod(base, base, mod)
		exp >>= 1
	}
	return res
}

func isPrime(n uint64) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}
	d := n - 1
	s := bits.TrailingZeros64(d)
	d >>= s
	for _, a := range [...]uint64{2, 3, 5, 7, 11, 13, 17} {
		if a >= n {
			continue
		}
		x := powMod(a, d, n)
		if x == 1 || x == n-1 {
			continue
		}
		composite := true
		for r := 1; r < s; r++ {
			x = mulMod(x, x, n)
			if x == n-1 {
				composite = false
				break
			}
		}
		if composite {
			return false
		}
	}
	return true
}

func main() {
	side := uint64(1)
	primes, diag := uint64(0), uint64(1)
	for {
		side += 2
		sq := side * side
		step := side - 1
		if isPrime(sq - step) {
			primes++
		}
		if isPrime(sq - 2*step) {
			primes++
		}
		if isPrime(sq - 3*step) {
			primes++
		}
		diag += 4
		if primes*10 < diag {
			fmt.Println("The answer is: ", side)
			return
		}
	}
}
