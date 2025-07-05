package main

import "fmt"

func main() {
	n := 50
	d := make([]int, n+1)
	d[0] = 1
	for i := 1; i <= n; i++ {
		d[i] = d[i-1]
		for j := 3; j <= i; j++ {
			if i-j == 0 {
				d[i]++
			} else {
				d[i] += d[i-j-1]
			}
		}
	}
	fmt.Println("The answer is: ", d[n])
}
