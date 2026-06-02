package main

import "fmt"

func main() {
	var res uint64
	var p16, p15, p14, p13 uint64 = 1, 1, 1, 1
	for l := 1; l <= 16; l++ {
		res += 15*p16 - 43*p15 + 41*p14 - 13*p13
		p16 *= 16
		p15 *= 15
		p14 *= 14
		p13 *= 13
	}
	fmt.Printf("%X\n", res)
}
