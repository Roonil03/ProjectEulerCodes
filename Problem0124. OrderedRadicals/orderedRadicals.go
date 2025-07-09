package main

import (
	"fmt"
	"sort"
)

func main() {
	const N = 100000
	const K = 10000
	r := make([]int, N+1)
	for i := 1; i <= N; i++ {
		r[i] = 1
	}
	for p := 2; p <= N; p++ {
		if r[p] == 1 {
			for m := p; m <= N; m += p {
				r[m] *= p
			}
		}
	}
	pairs := make([][2]int, N)
	for i := 1; i <= N; i++ {
		pairs[i-1] = [2]int{r[i], i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] < pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})
	fmt.Println("The answer is:", pairs[K-1][1])
}
