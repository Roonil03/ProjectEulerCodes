package main

import "fmt"

func main() {
	const limit = 10000000
	const searchRange = 5000
	primes := sieve(searchRange)
	var cans []int
	for _, p := range primes {
		if p >= 1000 && p <= searchRange {
			cans = append(cans, p)
		}
	}
	mti := float64(limit)
	res := 0
	for i := 0; i < len(cans); i++ {
		for j := i + 1; j < len(cans); j++ {
			p := cans[i]
			q := cans[j]
			n := p * q
			if n >= limit {
				break
			}
			phi := (p - 1) * (q - 1)
			if isPermutation(n, phi) {
				rt := float64(n) / float64(phi)
				if rt < mti {
					mti = rt
					res = n
				}
			}
		}
	}

	fmt.Printf("The answer is: %d\n", res)
}

func sieve(limit int) []int {
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
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPermutation(a, b int) bool {
	var freqA, freqB [10]int
	tempA := a
	for tempA > 0 {
		freqA[tempA%10]++
		tempA /= 10
	}
	tempB := b
	for tempB > 0 {
		freqB[tempB%10]++
		tempB /= 10
	}
	return freqA == freqB
}
