package main

import "fmt"

const W = 32

var ara []uint32

func f(p int, m uint32) {
	if p == W {
		ara = append(ara, m)
		return
	}
	for _, x := range []int{2, 3} {
		if p+x <= W {
			n := m
			if p+x < W {
				n |= 1 << uint(p+x)
			}
			f(p+x, n)
		}
	}
}

func main() {
	f(0, 0)
	g := make([][]int, len(ara))
	for i := range ara {
		for j := i + 1; j < len(ara); j++ {
			if ara[i]&ara[j] == 0 {
				g[i] = append(g[i], j)
				g[j] = append(g[j], i)
			}
		}
	}
	a := make([]int64, len(ara))
	for i := range a {
		a[i] = 1
	}
	for h := 1; h < 10; h++ {
		b := make([]int64, len(ara))
		for i := range a {
			for _, j := range g[i] {
				b[j] += a[i]
			}
		}
		a = b
	}
	var z int64
	for _, v := range a {
		z += v
	}
	fmt.Println(z)
}
