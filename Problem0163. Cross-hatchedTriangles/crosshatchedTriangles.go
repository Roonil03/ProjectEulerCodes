package main

import "fmt"

func main() {
	n := 36
	n2 := n * n
	n3 := n2 * n
	res := (1678*n3 + 3117*n2 + 88*n - 345*(n%2) - 320*(n%3) - 90*(n%4) - 288*((n3-n2+n)%5)) / 240
	fmt.Println(res)
}
