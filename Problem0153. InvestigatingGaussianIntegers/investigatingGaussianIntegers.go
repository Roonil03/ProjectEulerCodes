package main

import (
	"fmt"
)

const n = 100_000_000
const lim = 2_000_000

var f [lim + 1]int64

func h1(x int) int64 {
	if x <= lim {
		return f[x]
	}
	var res int64
	for l := 1; l <= x; {
		v := x / l
		r := x / v
		sum := int64(l+r) * int64(r-l+1) / 2
		res += int64(v) * sum
		l = r + 1
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	for i := 1; i <= lim; i++ {
		for j := i; j <= lim; j += i {
			f[j] += int64(i)
		}
		f[i] += f[i-1]
	}
	sum := h1(n)
	for a := 1; a*a < n; a++ {
		for b := 1; a*a+b*b <= n; b++ {
			if gcd(a, b) == 1 {
				sum += 2 * int64(a) * h1(n/(a*a+b*b))
			}
		}
	}
	fmt.Println(sum)
}
