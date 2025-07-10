package main

import (
	"fmt"
)

func makeSieve(n int) []bool {
	s := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		s[i] = true
	}
	for p := 2; p*p <= n; p++ {
		if s[p] {
			for q := p * p; q <= n; q += p {
				s[q] = false
			}
		}
	}
	return s
}

func main() {
	const target = 2000
	limit := 100000
	sieve := makeSieve(limit*12 + 10)
	count := 2
	var result int
	for k := 2; count < target; k++ {
		if sieve[6*k-1] && sieve[6*k+1] && sieve[12*k+5] {
			result = 3*k*(k-1) + 2
			count++
		}
		if count < target && sieve[6*k-1] && sieve[6*k+5] && sieve[12*k-7] {
			result = 3*k*(k+1) + 1
			count++
		}
	}
	fmt.Println("The answer is:", result)
}
