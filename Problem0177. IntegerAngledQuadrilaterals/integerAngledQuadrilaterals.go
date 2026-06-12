package main

import (
	"fmt"
	"math"
)

func h1(a1, a2, a3, a4, a5, a6, a7, a8 uint64) uint64 {
	b1 := [8]uint64{a1, a8, a7, a6, a5, a4, a3, a2}
	b2 := [8]uint64{a2, a3, a4, a5, a6, a7, a8, a1}
	m := ^uint64(0)
	for i := 0; i < 8; i += 2 {
		if v := (b1[i] << 56) | (b1[(i+1)%8] << 48) | (b1[(i+2)%8] << 40) | (b1[(i+3)%8] << 32) | (b1[(i+4)%8] << 24) | (b1[(i+5)%8] << 16) | (b1[(i+6)%8] << 8) | b1[(i+7)%8]; v < m {
			m = v
		}
		if v := (b2[i] << 56) | (b2[(i+1)%8] << 48) | (b2[(i+2)%8] << 40) | (b2[(i+3)%8] << 32) | (b2[(i+4)%8] << 24) | (b2[(i+5)%8] << 16) | (b2[(i+6)%8] << 8) | b2[(i+7)%8]; v < m {
			m = v
		}
	}
	return m
}

func main() {
	s, c := make([]float64, 180), make([]float64, 180)
	for i := range 180 {
		s[i] = math.Sin(float64(i) * math.Pi / 180)
		c[i] = math.Cos(float64(i) * math.Pi / 180)
	}
	var v [180][180]float64
	for i := 1; i < 180; i++ {
		for j := 1; j < 180-i; j++ {
			v[i][j] = s[j] / s[i+j]
		}
	}
	res := make(map[uint64]struct{})
	for x := 1; x < 180; x++ {
		for y := 1; y < 180-x; y++ {
			ab := v[x][y]
			for z := 1; z < 180-x; z++ {
				a := x + z
				sa, ca := s[a], c[a]
				for w := 1; w < 180-z; w++ {
					if y+w >= 180 {
						continue
					}
					a8 := math.Atan2(v[z][w]*sa, ab-v[z][w]*ca) * 180 / math.Pi
					k := int(math.Round(a8))
					if math.Abs(a8-float64(k)) <= 1e-9 {
						a3, a7, a4 := 180-a-k, 180-x-y-k, x-w+k
						if a3 > 0 && a4 > 0 && a7 > 0 && k > 0 {
							res[h1(uint64(x), uint64(z), uint64(a3), uint64(a4), uint64(w), uint64(y), uint64(a7), uint64(k))] = struct{}{}
						}
					}
				}
			}
		}
	}
	fmt.Println(len(res))
}
