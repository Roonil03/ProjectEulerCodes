package main

import (
	"fmt"
	"math"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func isSq(n int64) bool {
	r := int64(math.Sqrt(float64(n)))
	return r*r == n
}

func main() {
	const lim = 1_000_000_000_000
	set := map[int64]struct{}{}
	for b := int64(2); ; b++ {
		b3 := b * b * b
		if b3 >= lim {
			break
		}
		for a := int64(1); a < b; a++ {
			if gcd(a, b) != 1 {
				continue
			}
			for s := int64(1); ; s++ {
				n := a * s * (b3*s + a)
				if n >= lim {
					break
				}
				if isSq(n) {
					set[n] = struct{}{}
				}
			}
		}
	}
	sum := int64(0)
	for v := range set {
		if v > 1 {
			sum += v
		}
	}
	fmt.Println("The answer is:", sum)
}
