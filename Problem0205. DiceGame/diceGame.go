package main

import (
	"fmt"
)

func roll(d, s int) []float64 {
	dp := []float64{1}
	for i := 0; i < d; i++ {
		n := make([]float64, len(dp)+s)
		for j, v := range dp {
			for k := 1; k <= s; k++ {
				n[j+k] += v / float64(s)
			}
		}
		dp = n
	}
	return dp
}

func main() {
	p, c := roll(9, 4), roll(6, 6)
	res, acc := 0.0, 0.0
	for i, v := range p {
		if i > 0 && i-1 < len(c) {
			acc += c[i-1]
		}
		res += v * acc
	}
	fmt.Printf("%.7f\n", res)
}
