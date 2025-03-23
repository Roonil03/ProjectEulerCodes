package main

import "fmt"

func main() {
	res := 2
	n := 20
	for i := 3; i < n+1; i++ {
		res = lcm(res, i)
	}
	fmt.Print(res)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
