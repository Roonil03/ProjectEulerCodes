package main

import "fmt"

func main() {
	const N = 50
	ks := []int{2, 3, 4}
	res := 0
	for _, k := range ks {
		p := make([]int, N+1)
		p[0] = 1
		for i := 1; i <= N; i++ {
			if i < k {
				p[i] = 1
			} else {
				p[i] = p[i-1] + p[i-k]
			}
		}
		res += p[N] - 1
	}
	fmt.Println("The answer is:", res)
}
