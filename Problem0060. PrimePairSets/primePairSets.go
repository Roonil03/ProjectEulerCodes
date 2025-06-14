package main

import (
	"fmt"
	"math/bits"
)

func sieve(n int) []int {
	isP := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isP[i] = true
	}
	for p := 2; p*p <= n; p++ {
		if isP[p] {
			for m := p * p; m <= n; m += p {
				isP[m] = false
			}
		}
	}
	var primes []int
	for i := 3; i <= n; i += 2 {
		if isP[i] && i != 5 {
			primes = append(primes, i)
		}
	}
	return primes
}

func mulMod(a, b, m uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	_, r := bits.Div64(hi, lo, m)
	return r
}

func powMod(a, e, m uint64) uint64 {
	r := uint64(1)
	for e > 0 {
		if e&1 == 1 {
			r = mulMod(r, a, m)
		}
		a = mulMod(a, a, m)
		e >>= 1
	}
	return r
}

func isPrime(n uint64) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}
	d := n - 1
	s := bits.TrailingZeros64(d)
	d >>= s
	for _, a := range [...]uint64{2, 3, 5, 7, 11, 13, 17} {
		if a >= n {
			continue
		}
		x := powMod(a, d, n)
		if x == 1 || x == n-1 {
			continue
		}
		comp := true
		for r := 1; r < s; r++ {
			x = mulMod(x, x, n)
			if x == n-1 {
				comp = false
				break
			}
		}
		if comp {
			return false
		}
	}
	return true
}

func concat(a, b int) uint64 {
	x := uint64(b)
	p := uint64(1)
	for x > 0 {
		p *= 10
		x /= 10
	}
	return uint64(a)*p + uint64(b)
}

func main() {
	primes := sieve(10000)
	n := len(primes)
	ok := make([][]int, n)
	cache := make(map[uint64]bool)
	testPair := func(a, b int) bool {
		key := concat(a, b)
		if v, e := cache[key]; e {
			return v
		}
		res := isPrime(key)
		cache[key] = res
		return res
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if testPair(primes[i], primes[j]) && testPair(primes[j], primes[i]) {
				ok[i] = append(ok[i], j)
			}
		}
	}
	best := 1<<31 - 1
	stack := make([]int, 0, 5)
	var dfs func(start, depth, sum int)
	dfs = func(start, depth, sum int) {
		if depth == 5 {
			if sum < best {
				best = sum
			}
			return
		}
		for _, idx := range ok[start] {
			p := primes[idx]
			nsum := sum + p
			if nsum >= best {
				continue
			}
			okAll := true
			for _, s := range stack {
				if !contains(ok[s], idx) {
					okAll = false
					break
				}
			}
			if !okAll {
				continue
			}
			stack = append(stack, idx)
			dfs(idx, depth+1, nsum)
			stack = stack[:len(stack)-1]
		}
	}
	for i, p := range primes {
		if p*5 >= best {
			break
		}
		stack = append(stack, i)
		dfs(i, 1, p)
		stack = stack[:0]
	}
	fmt.Println("The answer is: ", best)
}

func contains(arr []int, x int) bool {
	i, j := 0, len(arr)
	for i < j {
		m := (i + j) >> 1
		if arr[m] < x {
			i = m + 1
		} else {
			j = m
		}
	}
	return i < len(arr) && arr[i] == x
}
