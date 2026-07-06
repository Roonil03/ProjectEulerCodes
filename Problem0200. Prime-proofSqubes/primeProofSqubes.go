package main

import (
	"fmt"
	"slices"
	"strconv"
)

func isPrime(x int64) bool {
	if x < 2 {
		return false
	}
	for i := int64(2); i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func contain200(x int64) bool {
	for x >= 200 {
		if x%1000 == 200 {
			return true
		}
		x /= 10
	}
	return false
}

func main() {
	const limit = 1000000
	np := make([]bool, limit+1)
	var primes []int64
	for i := 2; i <= limit; i++ {
		if !np[i] {
			primes = append(primes, int64(i))
			for j := 2 * i; j <= limit; j += i {
				np[j] = true
			}
		}
	}
	var squbes []int64
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		p2 := p * p
		if p2 > 1e12 {
			break
		}
		for j := 0; j < len(primes); j++ {
			if i == j {
				continue
			}
			q := primes[j]
			q3 := q * q * q
			if p2 > 1e12/q3 {
				break
			}
			v := p2 * q3
			if contain200(v) {
				squbes = append(squbes, v)
			}
		}
	}
	slices.Sort(squbes)
	count := 0
	for _, v := range squbes {
		s := strconv.FormatInt(v, 10)
		proof := true
		for id, c := range s {
			origDigit := int(c - '0')
			for d := 0; d <= 9; d++ {
				if d == origDigit || (id == 0 && d == 0) {
					continue
				}
				var temp int64
				for k := 0; k < len(s); k++ {
					if k == id {
						temp = temp*10 + int64(d)
					} else {
						temp = temp*10 + int64(s[k]-'0')
					}
				}
				if isPrime(temp) {
					proof = false
					break
				}
			}
			if !proof {
				break
			}
		}
		if proof {
			count++
			if count == 200 {
				fmt.Println(v)
				return
			}
		}
	}
}
