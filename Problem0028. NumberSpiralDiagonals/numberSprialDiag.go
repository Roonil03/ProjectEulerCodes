package main

import "fmt"

func main() {
	res := diagSum(1001)
	fmt.Println("The answer is:", res)
}

func diagSum(n int64) int64 {
	m := (n - 1) / 2
	t1 := 8 * m * (m + 1) * (2*m + 1) / 3
	t2 := 2 * m * (m + 1)
	t3 := 4 * m
	return 1 + t1 + t2 + t3
}
