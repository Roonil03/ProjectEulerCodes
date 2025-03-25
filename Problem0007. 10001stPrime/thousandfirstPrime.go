package main

import (
	"fmt"
	"math"
)

// const BIN_SIZE = 1 << 16

// func main() {
// 	n := 10001 - 1
// 	fmt.Print("The prime number is: ", nthPrime(n))
// }

// func sieveOfEratosthenes(start, end int, baseprimes []int) []bool {
// 	size := end - start
// 	sieve := make([]bool, size)
// 	for _, p := range baseprimes {
// 		firstMultiple := start
// 		if start%p != 0 {
// 			firstMultiple = (start/p + 1) * p
// 		}
// 		for i := firstMultiple; i < end; i += p {
// 			sieve[i-start] = true
// 		}
// 	}
// 	return sieve
// }

// func generateBasePrimes(limit int) []int {
// 	sieve := make([]bool, limit+1)
// 	primes := []int{}
// 	for i := 2; i <= limit; i++ {
// 		if !sieve[i] {
// 			primes = append(primes, i)
// 			for j := i * i; j <= limit; j += i {
// 				sieve[j] = true
// 			}
// 		}
// 	}
// 	return primes
// }

// func nthPrime(n int) int {
// 	if n <= 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 2
// 	}
// 	count := 1
// 	num := 1
// 	maxPossiblePrime := int(float64(n) * math.Log(float64(n)) * 1.15)
// 	sqrtMax := int(math.Sqrt(float64(maxPossiblePrime)))
// 	baseprimes := generateBasePrimes(sqrtMax)
// 	for count < n {
// 		start := num
// 		end := start + BIN_SIZE
// 		if end > maxPossiblePrime {
// 			end = maxPossiblePrime
// 		}
// 		sieve := sieveOfEratosthenes(start, end, baseprimes)
// 		for i := 0; i < len(sieve); i++ {
// 			if !sieve[i] && (start+i) > 1 {
// 				count++
// 				if count == n {
// 					return start + i
// 				}
// 			}
// 		}
// 		num = end
// 		if num >= maxPossiblePrime {
// 			break
// 		}
// 	}
// 	return -1
// }

func main() {
	n := 10001
	result := nthPrime(n)
	fmt.Printf("The %dth prime number is: %d\n", n, result)
}

func nthPrime(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 2
	}
	maxNum := int(float64(n) * (math.Log(float64(n)) + math.Log(math.Log(float64(n)))))
	sieve := make([]bool, maxNum+1)
	for i := range sieve {
		sieve[i] = true
	}
	for i := 2; i*i <= maxNum; i++ {
		if sieve[i] {
			for j := i * i; j <= maxNum; j += i {
				sieve[j] = false
			}
		}
	}
	count := 0
	for i := 2; i <= maxNum; i++ {
		if sieve[i] {
			count++
			if count == n {
				return i
			}
		}
	}
	return -1
}
