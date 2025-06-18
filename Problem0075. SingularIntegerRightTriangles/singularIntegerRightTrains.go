package main

import (
	"fmt"
	"math"
)

func main() {
	const lim = 1500000
	c := make([]int, lim+1)
	a := int(math.Sqrt(float64(lim / 2)))
	for m := 2; m <= a; m++ {
		s := 1
		if m&1 == 1 {
			s = 2
		}
		for n := s; n < m; n += 2 {
			if gcd(m, n) != 1 {
				continue
			}
			p := 2 * m * (m + n)
			if p > lim {
				break
			}
			for k := p; k <= lim; k += p {
				c[k]++
			}
		}
	}
	res := 0
	for _, c := range c {
		if c == 1 {
			res++
		}
	}
	fmt.Println("The answer is:", res)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
