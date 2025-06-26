package main

import "fmt"

const maxSum = 567
const limit = 10000000

func main() {
	var sq [10]int
	for d := 0; d < 10; d++ {
		sq[d] = d * d
	}
	next := func(n int) int {
		s := 0
		for n > 0 {
			s += sq[n%10]
			n /= 10
		}
		return s
	}
	ends89 := make([]bool, maxSum+1)
	for i := 1; i <= maxSum; i++ {
		n := i
		for n != 1 && n != 89 {
			n = next(n)
		}
		ends89[i] = (n == 89)
	}
	count := 0
	for x := 1; x < limit; x++ {
		if ends89[next(x)] {
			count++
		}
	}
	fmt.Println("The answer is:", count)
}
