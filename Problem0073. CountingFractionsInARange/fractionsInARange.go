package main

import "fmt"

const lim = 12000

func main() {
	fmt.Println("The answer is: ", count(1, 3, 1, 2))
}

func count(ln, ld, rn, rd int) int {
	md := ld + rd
	if md > lim {
		return 0
	}
	mn := ln + rn
	return 1 + count(ln, ld, mn, md) + count(mn, md, rn, rd)
}
