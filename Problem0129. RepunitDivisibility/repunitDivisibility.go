package main

import "fmt"

func main() {
	const limit = 1000000
	for n := limit + 1; ; n += 2 {
		if n%5 == 0 {
			continue
		}
		rem := 1 % n
		k := 1
		for ; k <= limit; k++ {
			if rem == 0 {
				break
			}
			rem = (rem*10 + 1) % n
		}
		if k > limit {
			fmt.Println("The answer is:", n)
			return
		}
	}
}
