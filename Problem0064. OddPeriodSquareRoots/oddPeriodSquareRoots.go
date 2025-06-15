package main

import (
	"fmt"
	"math"
)

func isqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func periodLength(n int) int {
	a0 := isqrt(n)
	if a0*a0 == n {
		return 0
	}
	seen := make(map[[2]int]bool)
	m := a0
	d := n - a0*a0
	res := 0
	for !seen[[2]int{m, d}] {
		seen[[2]int{m, d}] = true
		a := (a0 + m) / d
		m = a*d - m
		d = (n - m*m) / d
		res++
	}
	return res
}

func main() {
	res := 0
	for n := 2; n <= 10000; n++ {
		if periodLength(n)%2 == 1 {
			res++
		}
	}
	fmt.Println("The answer is:", res)
}
