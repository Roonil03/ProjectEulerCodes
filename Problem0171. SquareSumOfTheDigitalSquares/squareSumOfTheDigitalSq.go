package main

import "fmt"

func main() {
	var c, s [21][1621]int64
	c[0][0] = 1
	p := int64(1)
	for i := 0; i < 20; i++ {
		for j := 0; j <= 1620; j++ {
			if c[i][j] > 0 {
				for d := range int64(10) {
					n := j + int(d*d)
					c[i+1][n] = (c[i+1][n] + c[i][j]) % 1000000000
					s[i+1][n] = (s[i+1][n] + s[i][j] + d*p%1000000000*c[i][j]) % 1000000000
				}
			}
		}
		p = p * 10 % 1000000000
	}
	var a int64
	for k := 1; k*k <= 1620; k++ {
		a = (a + s[20][k*k]) % 1000000000
	}
	fmt.Printf("%09d\n", a)
}
