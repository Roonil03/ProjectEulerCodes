package main

import "fmt"

func h1(a, b int) (s int64) {
	var p [10]int64
	for i := a; i <= b; i++ {
		if i > 0 {
			p[i] = 1
		}
	}
	for range 40 {
		var n [10]int64
		for i := a; i <= b; i++ {
			s += p[i]
			if i > a {
				n[i-1] += p[i]
			}
			if i < b {
				n[i+1] += p[i]
			}
		}
		p = n
	}
	return
}

func main() {
	fmt.Print(h1(0, 9) - h1(1, 9) - h1(0, 8) + h1(1, 8))
}
