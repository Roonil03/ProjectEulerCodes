package main

import "fmt"

func main() {
	var n uint64 = 1000000000
	var p [16]uint64
	p[0] = 1
	for i := 1; i < len(p); i++ {
		p[i] = p[i-1] * 28
	}
	var d [16]uint64
	var k int
	for t := n; t > 0; t /= 7 {
		k++
		d[k-1] = t % 7
	}
	var s uint64
	pre := uint64(1)
	for i := k - 1; i >= 0; i-- {
		s += pre * (d[i] * (d[i] + 1) / 2) * p[i]
		pre *= d[i] + 1
	}
	fmt.Println("The answer is:", s)
}
