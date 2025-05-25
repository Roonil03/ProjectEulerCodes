package main

import "fmt"

func main() {
	a, res := 0, 0
	num := 1000
	for i := 2; i < num; i++ {
		temp := helper(i)
		if temp > a {
			res = i
			a = temp
		}
	}
	fmt.Println("The number is: ", res, "\nThe length is: ", a)
}

func helper(n int) int {
	seen := make(map[int]int)
	rem, pos := 1, 0
	for rem != 0 {
		if p, ok := seen[rem]; ok {
			return pos - p
		}
		seen[rem] = pos
		rem = (rem * 10) % n
		pos++
	}
	return 0
}
