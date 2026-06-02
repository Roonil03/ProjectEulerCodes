package main

import "fmt"

func h1(a, b int64) (int64, int64, int64) {
	if a == 0 {
		return b, 0, 1
	}
	g, x1, y1 := h1(b%a, a)
	x := y1 - (b/a)*x1
	y := x1
	return g, x, y
}

func main() {
	var res int64 = 0
	mod := int64(100000)
	for c := int64(1); c <= 9; c++ {
		for d := c; d <= 9; d++ {
			dd := 10*c - 1
			a, b := dd, d
			for b != 0 {
				a, b = b, a%b
			}
			q := dd / a
			_, x, _ := h1(dd, mod)
			invdd := (x%mod + mod) % mod
			p10q := int64(100) % q
			p10m := int64(100) % mod
			for k := int64(2); k <= 100; k++ {
				if p10q == 1%q {
					num := (d * (p10m - 1 + mod)) % mod
					term := (num * invdd) % mod
					res = (res + term) % mod
				}
				p10q = (p10q * 10) % q
				p10m = (p10m * 10) % mod
			}
		}
	}
	fmt.Printf("%05d\n", res)
}
