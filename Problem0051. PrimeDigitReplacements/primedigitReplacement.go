package main

import (
	"fmt"
	"strconv"
)

func main() {
	lim := 1000000
	isPrime := sieve(lim)
	var primes []int
	for i := 2; i < lim; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	grps := make(map[string][]int)
	for _, p := range primes {
		s := strconv.Itoa(p)
		d := len(s)
		for mask := 1; mask < (1<<d)-1; mask++ {
			if mask&1 == 1 {
				continue
			}
			var firstDig byte
			for i := 0; i < d; i++ {
				if mask&(1<<(d-1-i)) != 0 {
					firstDig = s[i]
					break
				}
			}
			ok := true
			for i := 0; i < d; i++ {
				if mask&(1<<(d-1-i)) != 0 && s[i] != firstDig {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}
			pat := []rune(s)
			for i := 0; i < d; i++ {
				if mask&(1<<(d-1-i)) != 0 {
					pat[i] = '*'
				}
			}
			grps[string(pat)] = append(grps[string(pat)], p)
		}
	}
	res := lim
	for _, list := range grps {
		if len(list) >= 8 && list[0] < res {
			res = list[0]
		}
	}
	fmt.Println("The answer is: ", res)
}

func sieve(limit int) []bool {
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}
	for p := 2; p*p <= limit; p++ {
		if isPrime[p] {
			for m := p * p; m <= limit; m += p {
				isPrime[m] = false
			}
		}
	}
	return isPrime
}
