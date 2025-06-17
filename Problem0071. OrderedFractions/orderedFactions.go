package main

import "fmt"

func main() {
	const lim = 1000000
	res, temp := 0, 1
	for d := 1; d <= lim; d++ {
		if (3*d-1)%7 != 0 {
			continue
		}
		n := (3*d - 1) / 7
		if n*temp > res*d {
			res, temp = n, d
		}
	}
	fmt.Println("The answer is: ", res)
}
