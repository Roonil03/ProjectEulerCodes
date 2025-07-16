package main

import (
	"fmt"
	"math"
)

func main() {
	c := 0
	for i := 1; i <= 9; i++ {
		switch i % 4 {
		case 0, 2:
			c += 20 * int(math.Pow(30, float64(i/2-1)))
		case 3:
			c += 100 * int(math.Pow(500, float64(i/4)))
		case 1:
		}
	}
	fmt.Println("The answer is:", c)
}
