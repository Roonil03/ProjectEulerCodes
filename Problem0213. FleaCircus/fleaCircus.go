package main

import "fmt"

func main() {
	var mat [30][30]float64
	for r := range 30 {
		for c := range 30 {
			d := 4.0
			if r == 0 || r == 29 {
				d--
			}
			if c == 0 || c == 29 {
				d--
			}
			mat[r][c] = 1.0 / d
		}
	}
	a := 0.0
	for tr := range 15 {
		for tc := range 15 {
			var p [30][30]float64
			p[tr][tc] = 1.0
			for range 50 {
				var n [30][30]float64
				for r := range 30 {
					for c := range 30 {
						v := 0.0
						if r > 0 {
							v += p[r-1][c]
						}
						if r < 29 {
							v += p[r+1][c]
						}
						if c > 0 {
							v += p[r][c-1]
						}
						if c < 29 {
							v += p[r][c+1]
						}
						n[r][c] = v * mat[r][c]
					}
				}
				p = n
			}
			e := 1.0
			for r := range 30 {
				for c := range 30 {
					e *= 1.0 - p[r][c]
				}
			}
			a += e
		}
	}
	fmt.Printf("%.6f", a*4)
}
