package main

import (
	"fmt"
	"math"
)

func main() {
	s := 2.0 * 1053779.0 * math.Sqrt(3.0)
	var res int64
	x := int64(s/3.0) + 1
	for m := int64(2); m <= x; m++ {
		n2 := int64(s / (3.0 * float64(m)))
		n1 := m
		d := 9.0*float64(m*m) - 8.0*s
		if d < 0 {
			n1 = 1
		} else {
			n1 = m - int64((3.0*float64(m)-math.Sqrt(d))/4.0)
		}
		for i := int64(1); i < m; i++ {
			if i > n2 && i < n1 {
				i = n1 - 1
				continue
			}
			if (m-i)%3 != 0 {
				a, b := m, i
				for b != 0 {
					a, b = b, a%b
				}
				if a == 1 {
					res += int64(s/float64((m-i)*(m+2*i))) + int64(s/float64(3*m*i))
				}
			}
		}
	}
	fmt.Print(res)
}
