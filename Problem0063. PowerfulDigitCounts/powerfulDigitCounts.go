package main

import (
	"fmt"
	"math"
)

func countDigits(n int64) int {
	if n == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(float64(n)))) + 1
}

func power(base, exp int) int64 {
	result := int64(1)
	b := int64(base)
	for i := 0; i < exp; i++ {
		result *= b
	}
	return result
}

func main() {
	res := 0
	for base := 1; base <= 9; base++ {
		for exp := 1; exp <= 21; exp++ {
			num := power(base, exp)
			digits := countDigits(num)
			if digits == exp {
				res++
			} else if digits < exp {
				break
			}
		}
	}
	fmt.Println("Number of powerful digit counts:", res)
}
