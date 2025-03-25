package main

import (
	"fmt"
)

func main() {
	n := 1000
	// c = 1000 - a - b
	// a2 + b2 = (1000 - a - b)^2
	// a2 + b2 = 1000000 + a2 + b2 - 2(1000a + 1000b - ab)
	// 2ab = 1000000 - 2000a - 2000b
	// b(-2a + 2000) = 1000000 - 2000a
	//b = (1000000 - 2000a) / (-2a + 2000)
	k := n / 3
	for a := 3; a < k; a++ {
		t1 := n*n - 2*n*a
		t2 := 2*n - 2*a
		if t2 == 0 || t1%t2 != 0 {
			continue
		}
		b := t1 / t2
		if b <= a {
			continue
		}
		c := n - b - a
		if a*a+b*b == c*c {
			fmt.Print(a, " + ", b, " = ", c, "\t")
			fmt.Println(a * b * c)
			fmt.Print("Found")
			return
		}
	}
	fmt.Print("No Solution Found")
}
