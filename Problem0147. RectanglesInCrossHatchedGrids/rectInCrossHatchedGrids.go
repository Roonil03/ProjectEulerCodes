package main

import "fmt"

func main() {
	const W, H = 47, 43
	var total int64
	for w := 1; w <= W; w++ {
		uw := int64(w * (w + 1) / 2)
		for h := 1; h <= H; h++ {
			uh := int64(h * (h + 1) / 2)
			total += uw * uh
		}
	}
	for w := 1; w <= W; w++ {
		for h := 1; h <= H; h++ {
			m, n := w, h
			if m < n {
				m, n = n, m
			}
			total += int64(n) * ((2*int64(m)-int64(n))*(4*int64(n)*int64(n)-1) - 3) / 6
		}
	}
	fmt.Println("The answer is:", total)
}
