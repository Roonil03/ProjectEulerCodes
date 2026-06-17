package main

import "fmt"

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var res int64
	for e := 3; e < 3671136; e += 4 {
		if gcd(e, 3671136) == 1 && gcd(e-1, 1008) == 2 && gcd(e-1, 3642) == 2 {
			res += int64(e)
		}
	}
	fmt.Println(res)
}
