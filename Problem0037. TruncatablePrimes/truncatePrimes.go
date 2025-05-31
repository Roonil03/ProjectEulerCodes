package main

import "fmt"

func main() {
	limit := 10000000
	primes := sieve(limit)
	var truncatablePrimes []int
	count := 0
	for i := 11; i <= limit && count < 11; i++ {
		if primes[i] && truc(i, primes) {
			truncatablePrimes = append(truncatablePrimes, i)
			count++
		}
	}
	sum := 0
	for _, prime := range truncatablePrimes {
		sum += prime
		fmt.Print(prime, " ")
	}
	fmt.Println("Found the numbers: ", sum)
}

func sieve(limit int) map[int]bool {
	primes := make(map[int]bool)
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
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes[i] = true
		}
	}
	return primes
}

func valid(n int) bool {
	str := fmt.Sprintf("%d", n)
	first := str[0]
	if first != '2' && first != '3' && first != '5' && first != '7' {
		return false
	}
	last := str[len(str)-1]
	if last != '3' && last != '7' {
		return false
	}
	for i := 1; i < len(str)-1; i++ {
		digit := str[i]
		if digit != '1' && digit != '3' && digit != '7' && digit != '9' {
			return false
		}
	}
	return true
}

func truc(n int, primes map[int]bool) bool {
	if n < 10 {
		return false
	}
	if !valid(n) {
		return false
	}
	temp := n
	for temp > 0 {
		if !primes[temp] {
			return false
		}
		temp /= 10
	}
	temp = n / 10
	for temp > 0 {
		if !primes[temp] {
			return false
		}
		temp /= 10
	}
	return true
}
