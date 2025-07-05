package main

import "fmt"

func main() {
	b, n := 0, 100
	for {
		n++
		x, inc, dec := n, true, true
		prev := x % 10
		x /= 10
		for x > 0 {
			d := x % 10
			if d < prev {
				inc = false
			}
			if d > prev {
				dec = false
			}
			if !inc && !dec {
				b++
				break
			}
			prev = d
			x /= 10
		}
		if b*100 == n*99 {
			fmt.Println("The answer is: ", n)
			return
		}
	}
}
