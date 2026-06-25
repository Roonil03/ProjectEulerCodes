package main

import "fmt"

func main() {
	d := []int{1, 1, 1}
	p := []int{1, 3, 9, 27, 81, 243, 729, 2187, 6561, 19683}
	for i := 1; i < 8; i++ {
		c := make([]int, p[i+1])
		for j := 0; j < p[i]; j++ {
			for k := range 3 {
				c[k+j*3] = d[j]
			}
		}
		for j := 1; j <= i; j++ {
			n := make([]int, p[i+1])
			for k := 0; k < p[i+1]; k++ {
				if c[k] == 0 {
					continue
				}
				a := (k / p[j-1]) % 3
				b := (k / p[j]) % 3
				for a1 := range 3 {
					w := 1
					if a != a1 {
						w++
					}
					if b != a && b != a1 {
						w++
					}
					w = 3 - w
					if w > 0 {
						n[k-b*p[j]+a1*p[j]] += c[k] * w
					}
				}
			}
			c = n
		}
		d = c
	}
	res := 0
	for _, v := range d {
		res += v
	}
	fmt.Println(res)
}
