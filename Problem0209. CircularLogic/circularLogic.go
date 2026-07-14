package main

import "fmt"

func main() {
	l := make([]uint64, 65)
	l[1], l[2] = 1, 3
	for i := 3; i <= 64; i++ {
		l[i] = l[i-1] + l[i-2]
	}
	vis := make([]bool, 64)
	var res uint64 = 1
	for i := range 64 {
		if !vis[i] {
			n := 0
			for c := i; !vis[c]; {
				vis[c] = true
				n++
				c = (c&31)<<1 | (c>>5 ^ (c>>4)&(c>>3)&1)
			}
			res *= l[n]
		}
	}
	fmt.Print(res)
}
