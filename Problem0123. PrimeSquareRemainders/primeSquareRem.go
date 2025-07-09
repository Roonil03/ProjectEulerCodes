package main

import "fmt"

func main() {
	const target = 1e10
	primes := sieve(500000)
	for i, p := range primes {
		n := i + 1
		if n%2 == 1 {
			if float64(2*n)*float64(p) > target {
				fmt.Println("The answer is:", n)
				return
			}
		}
	}
}

func sieve(max int) []int {
	isPrime := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= max; i++ {
		if isPrime[i] {
			for j := i * i; j <= max; j += i {
				isPrime[j] = false
			}
		}
	}
	var ps []int
	for i := 2; i <= max; i++ {
		if isPrime[i] {
			ps = append(ps, i)
		}
	}
	return ps
}
