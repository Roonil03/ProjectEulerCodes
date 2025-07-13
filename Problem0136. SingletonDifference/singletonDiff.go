package main

import "fmt"

func main() {
	const L = 50000000
	c := make([]int, L)
	for m := 1; m < L*2; m++ {
		for k := m/5 + 1; k*2 < m; k++ {
			n := (m - k) * (k*5 - m)
			if n >= L {
				break
			}
			c[n]++
		}
	}
	ans := 0
	for _, x := range c {
		if x == 1 {
			ans++
		}
	}
	fmt.Println("The answer is:", ans)
}
