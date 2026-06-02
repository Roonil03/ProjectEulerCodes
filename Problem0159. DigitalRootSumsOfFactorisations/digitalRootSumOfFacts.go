package main

import "fmt"

func main() {
	m := 1000000
	d := make([]int, m)
	s := 0
	for i := 1; i < m; i++ {
		d[i] = (i-1)%9 + 1
	}
	for i := 2; i < m; i++ {
		s += d[i]
		for j := 2; j <= i && i*j < m; j++ {
			if v := d[i] + d[j]; v > d[i*j] {
				d[i*j] = v
			}
		}
	}
	fmt.Println(s)
}
