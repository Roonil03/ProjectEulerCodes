package main

import "fmt"

func dfs(i int, v int, ps []int, lim int) int {
	if i == len(ps) {
		return 1
	}
	p, r := ps[i], 0
	for v <= lim {
		r += dfs(i+1, v, ps, lim)
		if lim/p < v {
			break
		}
		v *= p
	}
	return r
}

func main() {
	ps := []int{97, 89, 83, 79, 73, 71, 67, 61, 59, 53, 47, 43, 41, 37, 31, 29, 23, 19, 17, 13, 11, 7, 5, 3, 2}
	fmt.Println(dfs(0, 1, ps, 1000000000))
}
