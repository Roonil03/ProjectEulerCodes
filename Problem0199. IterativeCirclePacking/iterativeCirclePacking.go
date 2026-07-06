package main

import (
	"fmt"
	"math"
)

type cc struct {
	a, b int64
}

func (c cc) f() float64 {
	return float64(c.a) + float64(c.b)*math.Sqrt(3)
}

func (c cc) f2() float64 {
	fa, fb := float64(c.a), float64(c.b)
	return fa*fa + 3*fb*fb + 2*fa*fb*math.Sqrt(3)
}

type gg struct {
	c1, c2, c3, co cc
}

func m(c1, c2, c3, co cc) gg {
	a := [3]cc{c1, c2, c3}
	if a[0].f() > a[1].f() {
		a[0], a[1] = a[1], a[0]
	}
	if a[1].f() > a[2].f() {
		a[1], a[2] = a[2], a[1]
	}
	if a[0].f() > a[1].f() {
		a[0], a[1] = a[1], a[0]
	}
	return gg{a[0], a[1], a[2], co}
}

func main() {
	o, in := cc{-3, 0}, cc{3, 2}
	c := 27 / in.f2()
	g := map[gg]int64{m(in, in, in, o): 1, m(o, in, in, in): 3}
	for range 10 {
		n := map[gg]int64{}
		for k, ct := range g {
			nw := cc{2*(k.c1.a+k.c2.a+k.c3.a) - k.co.a, 2*(k.c1.b+k.c2.b+k.c3.b) - k.co.b}
			c += float64(ct) * 9 / nw.f2()
			n[m(k.c1, k.c2, nw, k.c3)] += ct
			n[m(k.c2, k.c3, nw, k.c1)] += ct
			n[m(k.c3, k.c1, nw, k.c2)] += ct
		}
		g = n
	}
	fmt.Printf("%.8f\n", 1-c)
}
