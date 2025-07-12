package main

import "fmt"

func main() {
	const L = 1000000
	c := make([]int, L)
	for m := 1; m < L*2; m++ {
		for k := m/5 + 1; k*2 < m; k++ {
			t := (m - k) * (k*5 - m)
			if t >= L {
				break
			}
			c[t]++
		}
	}
	ans := 0
	for _, x := range c {
		if x == 10 {
			ans++
		}
	}
	fmt.Println("The answer is:", ans)
}
