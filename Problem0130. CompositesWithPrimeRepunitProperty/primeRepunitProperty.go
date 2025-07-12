package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}
	r := int(math.Sqrt(float64(n)))
	for d := 3; d <= r; d += 2 {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func A(n int) int {
	x, k := 1, 1
	for x != 0 {
		x = (x*10 + 1) % n
		k++
	}
	return k
}

func main() {
	found, sum, n := 0, 0, 91
	for found < 25 {
		if n%2 == 0 || n%5 == 0 || isPrime(n) {
			n++
			continue
		}
		k := A(n)
		if (n-1)%k == 0 {
			sum += n
			found++
		}
		n++
	}
	fmt.Println("The answer is:", sum)
}
