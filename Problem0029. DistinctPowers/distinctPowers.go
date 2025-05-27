package main

import (
	"fmt"
	"math"
)

type key struct {
	base, expo int
}

func main() {
	mem := make(map[key]bool)
	for a := 2; a <= 100; a++ {
		base, exp := helper(a)
		for b := 2; b <= 100; b++ {
			a := exp * b
			k := key{base: base, expo: a}
			mem[k] = true
		}
	}
	fmt.Println(len(mem))
}

func helper(a int) (int, int) {
	if a < 2 {
		return a, 1
	}
	maxE := int(math.Log2(float64(a)))
	for e := maxE; e >= 2; e-- {
		base := int(math.Round(math.Pow(float64(a), 1.0/float64(e))))
		product := 1
		for i := 0; i < e; i++ {
			product *= base
			if product > a {
				break
			}
		}
		if product == a {
			return base, e
		}
	}
	return a, 1
}
