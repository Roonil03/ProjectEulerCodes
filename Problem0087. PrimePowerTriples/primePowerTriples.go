package main

import (
	"fmt"
	"sort"
)

func main() {
	const limit = 50_000_000
	primes := sieve(7100)
	var squares, cubes, fourths []int
	for _, p := range primes {
		sq := p * p
		if sq >= limit {
			break
		}
		squares = append(squares, sq)
		cu := sq * p
		if cu < limit {
			cubes = append(cubes, cu)
		}
		fo := cu * p
		if fo < limit {
			fourths = append(fourths, fo)
		}
	}
	sums := make([]int, 0, len(squares)*len(cubes)*len(fourths))
	for _, a := range squares {
		for _, b := range cubes {
			ab := a + b
			if ab >= limit {
				break
			}
			for _, c := range fourths {
				sum := ab + c
				if sum >= limit {
					break
				}
				sums = append(sums, sum)
			}
		}
	}
	sort.Ints(sums)
	res := 0
	prev := -1
	for _, v := range sums {
		if v != prev {
			res++
			prev = v
		}
	}
	fmt.Println("The answer is: ", res)

}

func sieve(n int) []int {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for k := p * p; k <= n; k += p {
				isPrime[k] = false
			}
		}
	}
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
