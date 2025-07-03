package main

import (
	"fmt"
	"math"
)

func isPandigital9(x int) bool {
	var mask int
	for i := 0; i < 9; i++ {
		d := x % 10
		if d < 1 || d > 9 {
			return false
		}
		bit := 1 << d
		if mask&bit != 0 {
			return false
		}
		mask |= bit
		x /= 10
	}
	return mask == 0b1111111110
}

func main() {
	const mod = 1_000_000_000
	phi := (1 + math.Sqrt(5)) / 2
	log10Phi := math.Log10(phi)
	log10Sqrt5 := 0.5 * math.Log10(5)
	a, b := 1, 1
	for n := 3; ; n++ {
		c := (a + b) % mod
		a, b = b, c
		if !isPandigital9(c) {
			continue
		}
		x := float64(n)*log10Phi - log10Sqrt5
		frac := x - math.Floor(x)
		first9 := int(math.Pow10(8) * math.Pow(10, frac))
		if isPandigital9(first9) {
			fmt.Println("The answer is: ", n)
			return
		}
	}
}
