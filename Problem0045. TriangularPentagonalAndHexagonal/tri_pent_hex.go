package main

import (
	"fmt"
	"math"
)

func main() {
	//n(n+1)/2 = a(2a-1)
	//n(n+1) = 2a(2a - 1)

	var n int64
	for n = 144; ; n++ {
		h := n * (2*n - 1)
		if check(h) {
			fmt.Println("The answer is ", h)
			break
		}
	}
}

func check(n int64) bool {
	d := 1 + 24*n
	s := int64(math.Sqrt(float64(d)) + 0.5)
	if s*s != d {
		return false
	}
	return (1+s)%6 == 0
}
