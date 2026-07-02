package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return math.Floor(math.Exp2(30.403243784-x*x)) * 1e-9
}

func main() {
	u := -1.0
	for {
		v := f(u)
		if u == f(v) {
			fmt.Printf("%.9f\n", u+v)
			return
		}
		u = v
	}
}
