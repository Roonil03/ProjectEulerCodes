package main

import (
	"fmt"
	"math"
)

func main() {
	n := 3
	fmt.Print(palin(n))
}

func palin(n int) int {
	a := int(math.Pow(float64(10), float64(n)) - 1)
	b := 1 + a/10
	res := 0
	for i := a; i >= b; i-- {
		for j := i; j >= b; j-- {
			temp := i * j
			if temp < res {
				break
			}
			num, rev := temp, 0
			for num != 0 {
				rev = rev*10 + num%10
				num /= 10
			}
			if temp == rev && temp > res {
				res = temp
			}
		}
	}
	return res
}
