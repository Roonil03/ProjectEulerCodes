package main

import "fmt"

func main() {
	const limit = 1e12
	b, n := 15, 21
	for n <= limit {
		nb, nn := 3*b+2*n-2, 4*b+3*n-3
		b, n = nb, nn
	}
	fmt.Println("The answer is:", b)
}
