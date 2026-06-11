package main

import "fmt"

func main() {
	var d [250001]int
	var c int
	for i := 1; i <= 250000; i++ {
		for j := i; j <= 250000; j += i {
			d[j]++
		}
		if l := d[i] / 2; l > 0 && l < 11 {
			c++
		}
	}
	fmt.Println(c)
}
