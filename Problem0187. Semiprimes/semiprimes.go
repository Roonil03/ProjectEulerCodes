package main

import (
	"fmt"
)

func main() {
	const lim = 100000000
	pc := make([]int, lim/2)
	primes := make([]int, 0)
	count := 0
	for i := 2; i < lim/2; i++ {
		if pc[i] == 0 {
			primes = append(primes, i)
		}
		for _, v := range primes {
			if i*v >= lim/2 {
				break
			}
			pc[i*v] = pc[i] + 1
			if i%v == 0 {
				break
			}
		}
	}
	for _, v := range primes {
		if v*v >= lim {
			break
		}
		count++
	}
	for i := 0; i < len(primes); i++ {
		p1 := primes[i]
		if p1*p1 >= lim {
			break
		}
		for j := i + 1; j < len(primes); j++ {
			p2 := primes[j]
			if p1*p2 >= lim {
				break
			}
			count++
		}
	}
	fmt.Println(count)
}
