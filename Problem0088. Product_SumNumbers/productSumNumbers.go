package main

import "fmt"

const Kmax = 12000
const Limit = 2 * Kmax

var minProd [Kmax + 1]int

func dfs(start, P, S, d int) {
	k := P - S + d
	if k > 1 && k <= Kmax {
		if P < minProd[k] {
			minProd[k] = P
		}
	}
	for i := start; ; i++ {
		nextP := P * i
		if nextP > Limit {
			break
		}
		dfs(i, nextP, S+i, d+1)
	}
}

func main() {
	for k := 2; k <= Kmax; k++ {
		minProd[k] = 2 * k
	}
	dfs(2, 1, 0, 0)
	seen := make(map[int]struct{})
	sum := 0
	for k := 2; k <= Kmax; k++ {
		v := minProd[k]
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			sum += v
		}
	}
	fmt.Println("The answer is:", sum)
}
