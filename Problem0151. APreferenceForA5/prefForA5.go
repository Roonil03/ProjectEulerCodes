package main

import "fmt"

func main() {
	d := make([]float64, 1024)
	d[657] = 1
	res := 0.0
	for b := 2; b < 16; b++ {
		n := make([]float64, 1024)
		for s, p := range d {
			if p == 0 {
				continue
			}
			c8, c4, c2, c1 := s>>9, (s>>7)&3, (s>>4)&7, s&15
			t := c8 + c4 + c2 + c1
			if t == 1 {
				res += p
			}
			ft := float64(t)
			if c8 > 0 {
				n[((c8-1)<<9)|((c4+1)<<7)|((c2+1)<<4)|(c1+1)] += p * float64(c8) / ft
			}
			if c4 > 0 {
				n[(c8<<9)|((c4-1)<<7)|((c2+1)<<4)|(c1+1)] += p * float64(c4) / ft
			}
			if c2 > 0 {
				n[(c8<<9)|(c4<<7)|((c2-1)<<4)|(c1+1)] += p * float64(c2) / ft
			}
			if c1 > 0 {
				n[(c8<<9)|(c4<<7)|(c2<<4)|(c1-1)] += p * float64(c1) / ft
			}
		}
		d = n
	}
	fmt.Printf("%.6f\n", res)
}
