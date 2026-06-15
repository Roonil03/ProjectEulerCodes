package main

import (
	"fmt"
	"math"
	"math/big"
)

type gg struct {
	n, d int
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func h1(n, d int) gg {
	x := gcd(n, d)
	return gg{n / x, d / x}
}

func h3(n int) (int, bool) {
	s := int(math.Sqrt(float64(n)) + 0.5)
	return s, s*s == n
}

func main() {
	var p []gg
	vis := make(map[gg]bool)
	for d := 2; d <= 35; d++ {
		for n := 1; n < d; n++ {
			if gcd(n, d) == 1 {
				f := gg{n, d}
				p = append(p, f)
				vis[f] = true
			}
		}
	}
	memo := make(map[gg]bool)
	for _, x := range p {
		for _, y := range p {
			nxy, dxy := x.n*y.d+y.n*x.d, x.d*y.d
			h2 := func(zn, zd int) {
				z := h1(zn, zd)
				if vis[z] {
					memo[h1(nxy*z.d+z.n*dxy, dxy*z.d)] = true
				}
			}
			h2(nxy, dxy)
			h2(x.n*y.n, nxy)
			w := x.n*x.n*y.d*y.d + y.n*y.n*x.d*x.d
			if k, ok := h3(w); ok {
				h2(k, dxy)
				h2(x.n*y.n, k)
			}
		}
	}
	t := big.NewRat(0, 1)
	for s := range memo {
		t.Add(t, big.NewRat(int64(s.n), int64(s.d)))
	}
	fmt.Print(new(big.Int).Add(t.Num(), t.Denom()))
}
