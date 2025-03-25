package main

import (
	"fmt"
	"math"
)

func main() {
	n := 2000000
	res := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			res += i
		}
	}
	fmt.Print("Sum: ", res)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sq := int(math.Sqrt(float64(n)))
	for i := 3; i <= sq; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
