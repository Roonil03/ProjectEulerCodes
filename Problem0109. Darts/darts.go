package main

import (
	"fmt"
	"sort"
)

func main() {
	pts := []int{0, 25, 50}
	for i := 1; i <= 20; i++ {
		pts = append(pts, i, 2*i, 3*i)
	}
	sort.Ints(pts)
	dbl := []int{50}
	for i := 1; i <= 20; i++ {
		dbl = append(dbl, 2*i)
	}
	sort.Ints(dbl)
	m := len(pts)
	res := 0
	for t := 1; t < 100; t++ {
		for _, z := range dbl {
			if z > t {
				continue
			}
			rem := t - z
			if rem == 0 {
				res++
				continue
			}
			for i := 0; i < m; i++ {
				a := pts[i]
				if a > rem {
					break
				}
				for j := i; j < m; j++ {
					if a+pts[j] == rem {
						res++
					}
				}
			}
		}
	}
	fmt.Print("The answer is: ", res)
}
