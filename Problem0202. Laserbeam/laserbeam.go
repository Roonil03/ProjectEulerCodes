package main

import "fmt"

func main() {
	var n int64 = 12017639147
	s := (n + 3) / 2
	var p []int64
	temp := s
	for i := int64(2); i*i <= temp; i++ {
		if temp%i == 0 {
			p = append(p, i)
			for temp%i == 0 {
				temp /= i
			}
		}
	}
	if temp > 1 {
		p = append(p, temp)
	}
	c := (2 * s) % 3
	var g1 func(int, int64, int64) int64
	g1 = func(i int, d, mu int64) int64 {
		if i == len(p) {
			m := s / d
			if a := (c * d) % 3; m >= a {
				return mu * ((m-a)/3 + 1)
			}
			return 0
		}
		return g1(i+1, d, mu) + g1(i+1, d*p[i], -mu)
	}
	fmt.Print(g1(0, 1, 1))
}
