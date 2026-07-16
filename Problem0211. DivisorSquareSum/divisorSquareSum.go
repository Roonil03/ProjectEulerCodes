package main

import (
	"fmt"
	"math"
)

var primes []uint64
var res uint64

func h1(n uint64) bool {
	if (0x0202021202030213 & (1 << (n & 63))) == 0 {
		return false
	}
	rt := uint64(math.Sqrt(float64(n)))
	return rt*rt == n || (rt+1)*(rt+1) == n || (rt-1)*(rt-1) == n
}

func dfs(pid int, cur uint64, sig2 uint64) {
	if h1(sig2) {
		res += cur
	}
	for i := pid; i < len(primes); i++ {
		p := primes[i]
		if cur*p >= 64000000 {
			break
		}
		n := cur * p
		pp := p * p
		t := 1 + pp
		for n < 64000000 {
			dfs(i+1, n, sig2*t)
			if n*p >= 64000000 {
				break
			}
			n *= p
			pp *= p * p
			t += pp
		}
	}
}

func main() {
	sieve := make([]bool, 32000000)
	primes = append(primes, 2)
	for i := uint64(3); i*i < 64000000; i += 2 {
		if !sieve[i/2] {
			for j := i * i; j < 64000000; j += 2 * i {
				sieve[j/2] = true
			}
		}
	}
	for i := uint64(3); i < 64000000; i += 2 {
		if !sieve[i/2] {
			primes = append(primes, i)
		}
	}
	dfs(0, 1, 1)
	fmt.Print(res)
}
