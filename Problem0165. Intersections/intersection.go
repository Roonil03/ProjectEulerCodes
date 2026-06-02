package main

import "fmt"

type Point struct {
	x, y int64
}

type Segment struct {
	a, b Point
}

type Fraction struct {
	n, d int64
}

type Intersect struct {
	x, y Fraction
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func h2(n, d int64) Fraction {
	if d < 0 {
		n = -n
		d = -d
	}
	g := gcd(n, d)
	return Fraction{n / g, d / g}
}

func h1(a, b, c Point) int64 {
	return (b.x-a.x)*(c.y-a.y) - (b.y-a.y)*(c.x-a.x)
}

func main() {
	var s int64 = 290797
	var mod int64 = 50515093
	segs := make([]Segment, 5000)
	for i := range 5000 {
		var t [4]int64
		for j := range 4 {
			s = (s * s) % mod
			t[j] = s % 500
		}
		segs[i] = Segment{Point{t[0], t[1]}, Point{t[2], t[3]}}
	}
	pts := make(map[Intersect]struct{})
	for i := range 5000 {
		for j := i + 1; j < 5000; j++ {
			s1, s2 := segs[i], segs[j]
			cp1 := h1(s1.a, s1.b, s2.a)
			cp2 := h1(s1.a, s1.b, s2.b)
			cp3 := h1(s2.a, s2.b, s1.a)
			cp4 := h1(s2.a, s2.b, s1.b)
			if ((cp1 > 0 && cp2 < 0) || (cp1 < 0 && cp2 > 0)) && ((cp3 > 0 && cp4 < 0) || (cp3 < 0 && cp4 > 0)) {
				ax, ay := s1.a.x, s1.a.y
				bx, by := s1.b.x, s1.b.y
				cx, cy := s2.a.x, s2.a.y
				dx, dy := s2.b.x, s2.b.y
				denom := (ax-bx)*(cy-dy) - (ay-by)*(cx-dx)
				nx := (ax*by-ay*bx)*(cx-dx) - (ax-bx)*(cx*dy-cy*dx)
				ny := (ax*by-ay*bx)*(cy-dy) - (ay-by)*(cx*dy-cy*dx)
				pts[Intersect{h2(nx, denom), h2(ny, denom)}] = struct{}{}
			}
		}
	}
	fmt.Println(len(pts))
}
