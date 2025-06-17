package main

import "fmt"

func main() {
	const N = 1000000
	phi := make([]int, N+1)
	for i := 0; i <= N; i++ {
		phi[i] = i
	}
	for i := 2; i <= N; i++ {
		if phi[i] == i {
			for j := i; j <= N; j += i {
				phi[j] = phi[j] / i * (i - 1)
			}
		}
	}
	tot := 0
	for d := 2; d <= N; d++ {
		tot += phi[d]
	}
	fmt.Println("The answer is: ", tot)
}
