package main

import "fmt"

func main() {
	lim := 1000000
	primes := seive(lim)
	res := 0
	for p := range primes {
		if circ(p, primes) {
			res++
		}
	}
	fmt.Print("The answer is: ", res)
}

func seive(n int) map[int]bool {
	primes := make(map[int]bool)
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes[i] = true
		}
	}
	return primes
}

func circ(n int, primes map[int]bool) bool {
	if !valid(n) {
		return n == 2 || n == 5
	}
	og, cur := n, n
	for {
		cur = rot(cur)
		if !primes[cur] {
			return false
		}
		if cur == og {
			break
		}
	}
	return true
}

func valid(n int) bool {
	if n == 2 || n == 5 {
		return true
	}
	if n > 0 {
		d := n % 10
		if d != 1 && d != 3 && d != 7 && d != 9 {
			return false
		}
		n /= 10
	}
	return true
}

func rot(n int) int {
	if n < 10 {
		return n
	}
	d := n % 10
	r := n / 10
	s := 1
	temp := r
	for temp > 0 {
		s *= 10
		temp /= 10
	}
	return r + d*s
}
