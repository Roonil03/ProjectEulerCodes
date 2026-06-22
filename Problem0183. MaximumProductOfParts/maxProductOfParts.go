package main

import (
	"fmt"
	"math"
)

func main() {
	s, k, v := 0, 1, 4.0
	for n := 5; n <= 10000; n++ {
		for float64(n) > v {
			k++
			v = math.Exp(float64(k+1)*math.Log(float64(k+1)) - float64(k)*math.Log(float64(k)))
		}
		x, y := n, k
		for y != 0 {
			x, y = y, x%y
		}
		d := k / x
		for d%2 == 0 {
			d /= 2
		}
		for d%5 == 0 {
			d /= 5
		}
		if d == 1 {
			s -= n
		} else {
			s += n
		}
	}
	fmt.Print(s)
}
