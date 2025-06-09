package main

import (
	"fmt"
	"math"
)

func main() {
	const limit = 10000
	isPrime := sieve(limit)
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	for n := 9; n <= limit; n += 2 {
		if isPrime[n] {
			continue
		}
		r := false
		for _, p := range primes {
			if p >= n {
				break
			}
			rem := n - p
			if rem%2 != 0 {
				continue
			}
			s := int(math.Sqrt(float64(rem / 2)))
			if s*s*2 == rem {
				r = true
				break
			}
		}
		if !r {
			fmt.Println("The answer is: ", n)
			return
		}
	}
}

func sieve(limit int) []bool {
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}
	return isPrime
}
