package main

import "fmt"

func sieve(n int) []bool {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for m := p * p; m <= n; m += p {
				isPrime[m] = false
			}
		}
	}
	return isPrime
}

func main() {
	const lim = 1000000
	isPrime := sieve(lim)
	var primes []int
	for i := 2; i <= lim; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	psum := make([]int, len(primes)+1)
	for i, p := range primes {
		psum[i+1] = psum[i] + p
	}
	ml, res := 0, 0
	for i := 0; i < len(primes); i++ {
		for j := i + ml + 1; j < len(psum) && psum[j]-psum[i] < lim; j++ {
			sum := psum[j] - psum[i]
			if isPrime[sum] {
				ml = j - i
				res = sum
			}
		}
	}
	fmt.Println("The answer is: ", res)
}
