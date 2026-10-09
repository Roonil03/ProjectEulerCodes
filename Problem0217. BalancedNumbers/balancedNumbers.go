package main

import "fmt"

const mod int64 = 14348907

func step(c, s []int64, l int) ([]int64, []int64) {
	a := make([]int64, len(c)+9)
	b := make([]int64, len(c)+9)
	for i := range c {
		for d := l; d < 10; d++ {
			a[i+d] = (a[i+d] + c[i]) % mod
			b[i+d] = (b[i+d] + 10*s[i] + int64(d)*c[i]) % mod
		}
	}
	return a, b
}

func main() {
	rc, rs := []int64{1}, []int64{0}
	lc, ls := []int64{1}, []int64{0}
	p, res := int64(1), int64(0)
	for m := 0; m <= 23; m++ {
		q := p * 10 % mod
		for i := range rc {
			res = (res + 10*q%mod*ls[i]%mod*rc[i] + 45*p%mod*lc[i]%mod*rc[i] + 10*lc[i]%mod*rs[i]) % mod
			if m > 0 {
				res = (res + p*ls[i]%mod*rc[i] + lc[i]*rs[i]) % mod
			}
		}
		if m < 23 {
			rc, rs = step(rc, rs, 0)
			if m == 0 {
				lc, ls = step(lc, ls, 1)
			} else {
				lc, ls = step(lc, ls, 0)
			}
			p = p * 10 % mod
		}
	}
	fmt.Println(res)
}
