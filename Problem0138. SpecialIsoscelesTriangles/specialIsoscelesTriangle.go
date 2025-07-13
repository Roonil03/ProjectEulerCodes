package main

import "fmt"

func main() {
	const n = 12
	p, c := 1, 17
	s := c
	for i := 2; i <= n; i++ {
		x := 18*c - p
		s += x
		p, c = c, x
	}
	fmt.Println("the answer is:", s)
}
