package main

import (
	"fmt"
	"math/bits"
)

const limit = 150000000

var good = [...]uint64{1, 3, 7, 9, 13, 27}
var bad = [...]uint64{5, 11, 15, 17, 19, 21, 23, 25}
var wheelPrimes = [...]int{3, 7, 11, 13, 17, 19}

func sievePrimes(n int) []int {
	isComposite := make([]bool, n+1)
	primes := make([]int, 0, n/10)
	for i := 2; i <= n; i++ {
		if !isComposite[i] {
			primes = append(primes, i)
			if i <= n/i {
				for j := i * i; j <= n; j += i {
					isComposite[j] = true
				}
			}
		}
	}
	return primes
}

func buildResidues() (int, []int) {
	mod := 10
	for _, p := range wheelPrimes {
		mod *= p
	}
	dues := make([]int, 0)
	for r := 0; r < mod; r += 10 {
		ok := true
		for _, p := range wheelPrimes {
			x := r % p
			sq := (x * x) % p
			for _, d := range good {
				if (sq+int(d))%p == 0 {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			dues = append(dues, r)
		}
	}
	return mod, dues
}

func mulMod(a, b, mod uint64) uint64 {
	h, l := bits.Mul64(a, b)
	_, rem := bits.Div64(h, l, mod)
	return rem
}

func powMod(a, e, mod uint64) uint64 {
	res := uint64(1)
	for e > 0 {
		if e&1 == 1 {
			res = mulMod(res, a, mod)
		}
		a = mulMod(a, a, mod)
		e >>= 1
	}
	return res
}

func isPrime(n uint64) bool {
	if n < 2 {
		return false
	}
	small := [...]uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	for _, p := range small {
		if n%p == 0 {
			return n == p
		}
	}
	d := n - 1
	s := 0
	for d&1 == 0 {
		d >>= 1
		s++
	}
	bases := [...]uint64{2, 325, 9375, 28178, 450775, 9780504, 1795265022} // https://cp-algorithms.com/algebra/primality_tests.html
	for _, a := range bases {
		if a%n == 0 {
			continue
		}
		x := powMod(a%n, d, n)
		if x == 1 || x == n-1 {
			continue
		}
		fg := true
		for r := 1; r < s; r++ {
			x = mulMod(x, x, n)
			if x == n-1 {
				fg = false
				break
			}
		}
		if fg {
			return false
		}
	}
	return true
}

func passesExtraFilter(n uint64, extra []int) bool {
	for _, p := range extra {
		r := n % uint64(p)
		sq := (r * r) % uint64(p)
		for _, d := range good {
			if (sq+d)%uint64(p) == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	mod, dues := buildResidues()
	t := sievePrimes(100)
	wheels := map[int]bool{
		2: true, 3: true, 5: true, 7: true,
		11: true, 13: true, 17: true, 19: true,
	}
	extra := make([]int, 0)
	for _, p := range t {
		if !wheels[p] {
			extra = append(extra, p)
		}
	}
	var sum uint64
	for base := 0; base < limit; base += mod {
		for _, r := range dues {
			n := base + r
			if n >= limit {
				break
			}
			nu := uint64(n)
			if !passesExtraFilter(nu, extra) {
				continue
			}
			n2 := nu * nu
			ok := true
			for _, d := range good {
				if !isPrime(n2 + d) {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}
			for _, d := range bad {
				if isPrime(n2 + d) {
					ok = false
					break
				}
			}
			if ok {
				sum += nu
			}
		}
	}
	fmt.Println(sum)
}
