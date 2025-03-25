package main

import (
	"fmt"
	"math"
)

func main() {
	n := 500
	res := triangleNumber(n)
	fmt.Print("Answer: ", res)
}

func triangleNumber(n int) int {
	a := 0
	i := 1
	for {
		a += i
		div := count(a)
		if div > n {
			return a
		}
		i++
	}
}

func count(num int) int {
	if num == 1 {
		return 1
	}
	count := 2
	sq := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq; i++ {
		if num%i == 0 {
			if i == num/i {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}
