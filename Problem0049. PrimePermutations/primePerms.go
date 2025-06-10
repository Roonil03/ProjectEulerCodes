package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	const m1, m2 = 1000, 9999
	isPrime := sieve(m2)
	grps := make(map[string][]int)
	for n := m1; n <= m2; n++ {
		if isPrime[n] {
			k := digKey(n)
			grps[k] = append(grps[k], n)
		}
	}
	for _, p := range grps {
		if len(p) < 3 {
			continue
		}
		sort.Ints(p)
		for i := 0; i < len(p)-2; i++ {
			for j := i + 1; j < len(p)-1; j++ {
				p1, p2 := p[i], p[j]
				p3 := 2*p2 - p1
				if p3 > p2 && p3 <= m2 && isPrime[p3] && digKey(p3) == digKey(p1) && p1 != 1487 {
					fmt.Printf("The answer is : %d\n", p1*1e8+p2*1e4+p3)
					return
				}
			}
		}
	}
}

func sieve(max int) []bool {
	isPrime := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= max; i++ {
		if isPrime[i] {
			for j := i * i; j <= max; j += i {
				isPrime[j] = false
			}
		}
	}
	return isPrime
}

func digKey(n int) string {
	s := strconv.Itoa(n)
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
