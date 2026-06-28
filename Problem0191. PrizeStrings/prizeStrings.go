package main

import "fmt"

func main() {
	n := 30
	a := make([]int, n+1)
	a[0], a[1], a[2] = 1, 2, 4
	for i := 3; i <= n; i++ {
		a[i] = a[i-1] + a[i-2] + a[i-3]
	}
	res := a[n]
	for i := 1; i <= n; i++ {
		res += a[i-1] * a[n-i]
	}
	fmt.Println(res)
}
