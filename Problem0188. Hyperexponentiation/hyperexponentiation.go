package main

import "fmt"

func phi(n int) int {
	r := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			r -= r / i
		}
	}
	if n > 1 {
		r -= r / n
	}
	return r
}

func pow(a, b, m int) int {
	r := 1
	a %= m
	for b > 0 {
		if b%2 != 0 {
			r = (r * a) % m
		}
		a = (a * a) % m
		b /= 2
	}
	return r
}

func solve(a, b, m int) int {
	if m == 1 {
		return 0
	}
	if b == 0 {
		return 1
	}
	return pow(a, solve(a, b-1, phi(m)), m)
}

func main() {
	fmt.Printf("%08d\n", solve(1777, 1855, 100000000))
}
