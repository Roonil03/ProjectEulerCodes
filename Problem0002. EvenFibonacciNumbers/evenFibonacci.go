package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2
	var res int64
	res = 2
	for a+b < 4000000 {
		c := a + b
		if c%2 == 0 {
			res += int64(c)
		}
		a = b
		b = c
	}
	fmt.Print(res)
}
