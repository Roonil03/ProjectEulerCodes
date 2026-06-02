package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	ss := make([][]uint64, 19)
	ss[1] = []uint64{1<<32 | 1}
	seen := map[uint64]bool{1<<32 | 1: true}
	for i := 2; i <= 18; i++ {
		cur := make(map[uint64]bool)
		for j := 1; j <= i/2; j++ {
			for _, x := range ss[j] {
				xn, xd := int(x>>32), int(x&0xFFFFFFFF)
				for _, y := range ss[i-j] {
					yn, yd := int(y>>32), int(y&0xFFFFFFFF)
					np, dp := xn*yd+yn*xd, xd*yd
					ns, ds := xn*yn, np
					g1 := gcd(np, dp)
					f1 := uint64(np/g1)<<32 | uint64(dp/g1)
					if !seen[f1] {
						cur[f1] = true
					}
					g2 := gcd(ns, ds)
					f2 := uint64(ns/g2)<<32 | uint64(ds/g2)
					if !seen[f2] {
						cur[f2] = true
					}
				}
			}
		}
		for f := range cur {
			seen[f] = true
			ss[i] = append(ss[i], f)
		}
	}
	fmt.Println(len(seen))
}
