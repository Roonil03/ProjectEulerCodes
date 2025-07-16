package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 1.4, -9.6
	dx, dy := 1.4, -19.7
	hits := 1
	for {
		nx := 8 * x
		ny := 2 * y
		nLen := math.Sqrt(nx*nx + ny*ny)
		nx /= nLen
		ny /= nLen
		dot := dx*nx + dy*ny
		dx -= 2 * dot * nx
		dy -= 2 * dot * ny
		a := 4*dx*dx + dy*dy
		b := 2 * (4*x*dx + y*dy)
		c := 4*x*x + y*y - 100
		discriminant := b*b - 4*a*c
		t1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		t2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		var t float64
		if t1 > 1e-9 {
			t = t1
		} else {
			t = t2
		}
		x += t * dx
		y += t * dy
		if y > 0 && x >= -0.01 && x <= 0.01 {
			break
		}
		hits++
	}
	fmt.Println("The answer is:", hits)
}
