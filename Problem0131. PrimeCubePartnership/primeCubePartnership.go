package main

import "fmt"

func main() {
	const L = 1000000
	isP := make([]bool, L)
	for i := 2; i < L; i++ {
		isP[i] = true
	}
	for i := 2; i*i < L; i++ {
		if isP[i] {
			for j := i * i; j < L; j += i {
				isP[j] = false
			}
		}
	}
	res := 0
	for a := 1; ; a++ {
		p := 3*a*a + 3*a + 1
		if p >= L {
			break
		}
		if isP[p] {
			res++
		}
	}
	fmt.Println("The answer is:", res)
}
