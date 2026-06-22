package main

import (
	"fmt"
	"sort"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type pnt struct {
	x, y int
}

type ray struct {
	p pnt
	c int64
}

func main() {
	counts := make(map[pnt]int64)
	var n int64
	for x := -104; x <= 104; x++ {
		for y := -104; y <= 104; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if x*x+y*y < 11025 {
				g := gcd(abs(x), abs(y))
				counts[pnt{x / g, y / g}]++
				n++
			}
		}
	}
	var rays []ray
	for p, c := range counts {
		rays = append(rays, ray{p, c})
	}
	up := func(p pnt) bool {
		return p.y > 0 || (p.y == 0 && p.x > 0)
	}
	sort.Slice(rays, func(i, j int) bool {
		u1, u2 := up(rays[i].p), up(rays[j].p)
		if u1 != u2 {
			return u1
		}
		return rays[i].p.x*rays[j].p.y-rays[i].p.y*rays[j].p.x > 0
	})
	c2 := func(n int64) int64 {
		return n * (n - 1) / 2
	}
	c3 := func(n int64) int64 {
		return n * (n - 1) * (n - 2) / 6
	}
	var sl, so int64
	m := len(rays)
	j, s := 1, int64(0)
	for i := range m {
		if j <= i {
			j, s = i+1, 0
		}
		for {
			jm := j % m
			if jm == i {
				break
			}
			if rays[i].p.x*rays[jm].p.y-rays[i].p.y*rays[jm].p.x > 0 {
				s += rays[jm].c
				j++
			} else {
				break
			}
		}
		ci := rays[i].c
		sl += c3(ci) + c2(ci)*s + ci*c2(s)
		if j > i+1 {
			s -= rays[(i+1)%m].c
		}
	}
	for p, v := range counts {
		if up(p) {
			if cB, ok := counts[pnt{-p.x, -p.y}]; ok {
				so += v*cB*(n-v-cB) + c2(v)*cB + v*c2(cB)
			}
		}
	}
	fmt.Print(c3(n) - sl - so)
}
