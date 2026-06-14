package main

import "fmt"

func main() {
	const max = 10000000
	d := make([]uint16, max+1)
	for i := 2; i <= max/2; i++ {
		for j := 2 * i; j <= max; j += i {
			d[j]++
		}
	}
	res := 0
	for i := 2; i < max; i++ {
		if d[i] == d[i+1] {
			res++
		}
	}
	fmt.Println(res)
}
