package main

import (
	"fmt"
	"math"
)

func main() {
	r := uint64(1000000000)
	c := r / 4
	d := c*c/2 - 1
	m := uint64(math.Sqrt(float64(d)))
	for (m+1)*(m+1) <= d {
		m++
	}
	for m*m > d {
		m--
	}
	s, y := uint64(0), m
	for i := uint64(1); i <= m; i++ {
		for i*i+y*y > d {
			y--
		}
		s += y
	}
	fmt.Print(r*r + r*r/2 + 4*m + 4*s - c + 2)
}
