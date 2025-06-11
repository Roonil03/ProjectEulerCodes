package main

import "fmt"

func main() {
	const limit = 1000000
	count := 0
	for n := 1; n <= 100; n++ {
		c := 1
		for r := 1; r <= n/2; r++ {
			c = c * (n - r + 1) / r
			if c > limit {
				if n%2 == 0 && r == n/2 {
					count++
				} else {
					count += n - 2*r + 1
				}
				break
			}
		}
	}
	fmt.Println("The answer is: ", count)
}
