package main

import "fmt"

func main() {
	var c int
	for i := 1; i*i <= 250000; i++ {
		c += 250000/i - i
	}
	fmt.Println(c)
}
