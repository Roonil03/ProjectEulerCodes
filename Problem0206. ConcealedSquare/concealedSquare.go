package main

import "fmt"

func h1(n uint64) bool {
	for d := uint64(9); d > 0; d-- {
		if n%10 != d {
			return false
		}
		n /= 100
	}
	return true
}

func main() {
	for i := uint64(101010100); ; i += 10 {
		if h1((i + 3) * (i + 3)) {
			fmt.Println((i + 3) * 10)
			return
		}
		if h1((i + 7) * (i + 7)) {
			fmt.Println((i + 7) * 10)
			return
		}
	}
}
