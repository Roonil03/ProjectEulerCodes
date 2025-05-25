package main

import "fmt"

func main() {
	primes := doPrimes(1000)
	count, prod := 0, 0
	for _, b := range primes {
		s, step := -999, 1
		if b == 2 {
			s, step = -998, 2
		} else {
			s, step = -999, 2
		}
		for a := s; a < 1000; a += step {
			c := helper(a, b)
			if c > count {
				count = c
				prod = a * b
			}
		}
	}
	fmt.Println("The product of the coefficients is: ", prod)
}

func doPrimes(n int) []int {
	prime := make([]bool, n+1)
	for i := 2; i < n; i++ {
		prime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if prime[i] {
			for j := i * i; j <= n; j += i {
				prime[j] = false
			}
		}
	}
	var res []int
	for i := 2; i <= n; i++ {
		if prime[i] {
			res = append(res, i)
		}
	}
	return res
}

func helper(a, b int) int {
	n := 0
	for {
		val := n*n + a*n + b
		if val < 2 || !isPrime(val) {
			return n
		}
		n++
	}
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
