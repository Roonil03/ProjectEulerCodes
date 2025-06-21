package main

import "fmt"

func main() {
	const tar = 5000
	const res = 1000
	primes := sieve(res)
	ways := make([]int, res+1)
	ways[0] = 1
	for _, p := range primes {
		for i := p; i <= res; i++ {
			ways[i] += ways[i-p]
		}
	}
	for i := 2; i <= res; i++ {
		if ways[i] > tar {
			fmt.Printf("The answer is answer: %d\n", i)
			return
		}
	}
	fmt.Println("Increase maxNum")
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
