package main

import "fmt"

func main() {
	const limit = 1000000
	primes := sieve(50)
	res := 1
	for _, p := range primes {
		if res*p > limit {
			break
		}
		res *= p
	}
	fmt.Println("The answer is: ", res)
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
