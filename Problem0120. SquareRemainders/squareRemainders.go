package main

import "fmt"

func main() {
	s := 0
	for a := 3; a <= 1000; a++ {
		if a%2 == 0 {
			s += a * (a - 2)
		} else {
			s += a * (a - 1)
		}
	}
	fmt.Println("The answer is:", s)
}
