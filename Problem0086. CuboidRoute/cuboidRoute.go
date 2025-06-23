package main

import (
	"fmt"
	"math"
)

func main() {
	n := 1000000
	count := 0
	for c := 1; ; c++ {
		for L := 2; L <= 2*c; L++ { // L = a+b
			if !isSquare(L*L + c*c) {
				continue
			}
			var add int
			if L <= c+1 {
				add = L / 2
			} else {
				add = (2*c - L + 2) / 2
			}
			count += add
		}
		if count > n {
			fmt.Print("The answer is: ", c)
			break
		}
	}
}

func isSquare(n int) bool {
	r := int(math.Sqrt(float64(n)))
	return r*r == n
}
