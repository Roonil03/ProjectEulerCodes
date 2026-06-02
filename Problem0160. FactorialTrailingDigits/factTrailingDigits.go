package main

import (
	"fmt"
)

func main() {
	var N uint64 = 1000000000000
	fmt.Println(solve(N))
}

func solve(n uint64) int {
	mod := uint64(100000)
	g5 := h1(n, 5)
	res3125 := hmod(n)
	inv2 := uint64(1563)
	rem3125 := (res3125 * powMod(inv2, g5, 3125)) % 3125
	res := (rem3125 * 9376) % mod
	return int(res)
}

func h1(n, p uint64) uint64 {
	var count uint64
	for n > 0 {
		count += n / p
		n /= p
	}
	return count
}

func hmod(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return (h2(n) * hmod(n/5)) % 3125
}

func h2(n uint64) uint64 {
	blocks := n / 3125
	rem := n % 3125
	res := uint64(1)
	for i := uint64(1); i <= rem; i++ {
		if i%5 != 0 {
			res = (res * i) % 3125
		}
	}
	if blocks%2 != 0 {
		return (3124 * res) % 3125
	}
	return res
}

func powMod(base, exp, mod uint64) uint64 {
	res := uint64(1)
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return res
}
