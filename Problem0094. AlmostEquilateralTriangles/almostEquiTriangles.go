package main

import "fmt"

func main() {
	const Limit = 1000000000
	s0, s1, m := 1, 1, 1
	sum := 0
	for {
		s2 := 4*s1 - s0 + 2*m
		s0, s1 = s1, s2
		m = -m
		p := 3*s1 - m
		if p > Limit {
			break
		}
		sum += p
	}
	fmt.Println("The answer is:", sum)
}
