package main

import (
	"fmt"
)

func powmod(a, e, m int) int {
	r := 1 % m
	for e > 0 {
		if e&1 == 1 {
			r = (r * a) % m
		}
		a = (a * a) % m
		e >>= 1
	}
	return r
}

func main() {
	const K = 40
	const L = 2000000
	bs := make([]bool, L+1)
	ps := []int{}
	for i := 2; i <= L; i++ {
		if !bs[i] {
			ps = append(ps, i)
			for j := i * i; j <= L; j += i {
				bs[j] = true
			}
		}
	}
	sum, cnt := 0, 0
	for _, p := range ps {
		if p == 2 || p == 5 {
			continue
		}
		if powmod(10, 1_000_000_000, 9*p) == 1 {
			sum += p
			cnt++
			if cnt == K {
				break
			}
		}
	}
	fmt.Println("The answer is:", sum)
}
