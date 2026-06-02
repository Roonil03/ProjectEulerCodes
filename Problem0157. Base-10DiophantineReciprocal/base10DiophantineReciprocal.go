package main

import (
	"fmt"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func h1(n int64) int64 {
	count := int64(1)
	for d := int64(2); d*d <= n; d++ {
		if n%d == 0 {
			exp := int64(0)
			for n%d == 0 {
				exp++
				n /= d
			}
			count *= (exp + 1)
		}
	}
	if n > 1 {
		count *= 2
	}
	return count
}

func main() {
	pow2 := make([]int64, 19)
	pow5 := make([]int64, 19)
	p2, p5 := int64(1), int64(1)
	for i := 0; i <= 18; i++ {
		pow2[i] = p2
		pow5[i] = p5
		p2 *= 2
		p5 *= 5
	}
	var res int64
	for n := 1; n <= 9; n++ {
		var a int64 = 1
		for i := 0; i < n; i++ {
			a *= 10
		}
		aa := a * a
		var temp int64
		for i := 0; i <= 2*n; i++ {
			for y := 0; y <= 2*n; y++ {
				d := pow2[i] * pow5[y]
				if d <= a {
					dPrime := aa / d
					g := gcd(a+d, a+dPrime)
					temp += h1(g)
				}
			}
		}
		fmt.Printf("n = %d: %d solutions\n", n, temp)
		res += temp
	}
	fmt.Printf("\nFinal Answer: %d\n", res)
}
