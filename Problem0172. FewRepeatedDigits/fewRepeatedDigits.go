package main

import "fmt"

func main() {
	var c [19][19]int64
	for i := range 19 {
		c[i][0] = 1
		for j := 1; j <= i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
	var d [11][19]int64
	d[0][0] = 1
	for i := 1; i <= 10; i++ {
		for j := 0; j <= 18; j++ {
			for k := 0; k <= 3 && k <= j; k++ {
				d[i][j] += d[i-1][j-k] * c[j][k]
			}
		}
	}
	fmt.Println(d[10][18] / 10 * 9)
}
