package main

import (
	"fmt"
	"math"
)

func main() {
	res := int64(1<<62 - 1)
	for n := int64(2); ; n++ {
		pN := pent(n)
		if pN-pent(n-1) > res {
			break
		}
		for k := int64(1); k < n; k++ {
			d := pN - pent(n-k)
			if d > res {
				break
			}
			if check(pN+pent(n-k)) && check(d) {
				res = d
			}
		}
	}
	fmt.Println("The answer is:", res)
}

func pent(n int64) int64 {
	return n * (3*n - 1) / 2
}

func check(x int64) bool {
	if x <= 0 {
		return false
	}
	n := (1 + math.Sqrt(float64(1+24*x))) / 6
	return n == float64(int64(n))
}
