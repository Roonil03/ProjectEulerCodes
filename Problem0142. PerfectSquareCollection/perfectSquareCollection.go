package main

import (
	"fmt"
	"math"
)

func main() {
	res := int64(1<<63 - 1)
	maxS := 2000
	isSq := make([]bool, (maxS+1)*(maxS+1)+1)
	for i := 0; i <= maxS; i++ {
		isSq[i*i] = true
	}
	for a := 1; a <= maxS; a++ {
		a2 := int64(a * a)
		for b := a % 2; b < a; b += 2 {
			b2 := int64(b * b)
			x := (a2 + b2) >> 1
			y := (a2 - b2) >> 1
			if y <= 0 || x <= y {
				continue
			}
			sumXY := x + y
			if sumXY+3 >= res {
				continue
			}
			startC := int64(math.Sqrt(float64(y))) + 1
			for c := startC; ; c++ {
				c2 := c * c
				z := c2 - y
				if z <= 0 || y <= z {
					if z >= y {
						break
					}
					continue
				}
				total := sumXY + z
				if total >= res {
					break
				}
				xz1 := x + z
				xz2 := x - z
				yz := y - z
				if xz2 > 0 && xz1 <= int64(len(isSq)-1) && xz2 <= int64(len(isSq)-1) && yz <= int64(len(isSq)-1) {
					if isSq[xz1] && isSq[xz2] && isSq[yz] {
						res = total
					}
				}
			}
		}
	}
	fmt.Println("The answer is:", res)
}
