package main

import (
	"fmt"
)

func poly(t, n int) int {
	return ((t-2)*n*n + (4-t)*n) / 2
}

func main() {
	const (
		low  = 1000
		high = 10000
	)
	types := []int{3, 4, 5, 6, 7, 8}
	byPrefix := make(map[int]map[int][]int)
	numsOfType := make(map[int][]int)
	for _, t := range types {
		byPrefix[t] = make(map[int][]int)
		for n := 1; ; n++ {
			v := poly(t, n)
			if v >= high {
				break
			}
			if v >= low && v%100 >= 10 { // 4-digit and no xx00/xx01 … xx09 tails
				numsOfType[t] = append(numsOfType[t], v)
				p := v / 100
				byPrefix[t][p] = append(byPrefix[t][p], v)
			}
		}
	}
	best := 1 << 30
	var permute func([]int, int)
	permute = func(a []int, idx int) {
		if idx == len(a) {
			tryOrder(a, byPrefix, numsOfType, &best)
			return
		}
		for i := idx; i < len(a); i++ {
			a[idx], a[i] = a[i], a[idx]
			permute(a, idx+1)
			a[idx], a[i] = a[i], a[idx]
		}
	}
	permute(types, 0)
	fmt.Println("The answer is: ", best)
}

func tryOrder(order []int, byPref map[int]map[int][]int,
	nums map[int][]int, best *int) {
	var dfs func(pos int, first, lastPref, sum int, used []int)
	dfs = func(pos int, first, lastPref, sum int, used []int) {
		if sum >= *best {
			return
		}
		if pos == len(order) {
			if lastPref == first/100 {
				*best = sum
			}
			return
		}
		t := order[pos]
		if pos == 0 {
			for _, v := range nums[t] {
				dfs(pos+1, v, v%100, sum+v, []int{v})
			}
			return
		}
		for _, v := range byPref[t][lastPref] {
			dfs(pos+1, first, v%100, sum+v, append(used, v))
		}
	}
	dfs(0, 0, 0, 0, nil)
}
