package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(largestPF(600851475143))
}

func largestPF(n int64) int64 {
	if n <= 1 {
		return n
	}
	res := int64(2)
	for n%2 == 0 {
		n = n / 2
	}
	sqrtN := int64(math.Sqrt(float64(n)))
	for i := int64(3); i <= sqrtN && n > 1; i += 2 {
		for n%i == 0 {
			res = i
			n = n / i
		}
	}
	if n > 2 {
		res = n
	}
	return res
}
