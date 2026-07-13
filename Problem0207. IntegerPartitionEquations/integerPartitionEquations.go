package main

import "fmt"

func main() {
	var n uint64 = 1
	for {
		x := 12345*n + 2
		if x < 1<<(n+1) {
			if x < 1<<n {
				x = 1 << n
			}
			fmt.Println(x * (x - 1))
			return
		}
		n++
	}
}
