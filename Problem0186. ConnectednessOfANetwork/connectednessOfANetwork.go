package main

import "fmt"

func main() {
	par := make([]int, 1000000)
	size := make([]int, 1000000)
	for i := range par {
		par[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(i int) int {
		root := i
		for par[root] != root {
			root = par[root]
		}
		cur := i
		for cur != root {
			nxt := par[cur]
			par[cur] = root
			cur = nxt
		}
		return root
	}
	var hist [55]int
	for i := 1; i <= 55; i++ {
		iii := int64(i) * int64(i) * int64(i)
		val := (100003 - 200003*int64(i) + 300007*iii) % 1000000
		if val < 0 {
			val += 1000000
		}
		hist[i-1] = int(val)
	}
	res := 0
	k := 1
	for size[find(524287)] < 990000 {
		var caller, called int
		if k <= 55 {
			caller = hist[k-1]
		} else {
			id := (k - 1) % 55
			caller = (hist[id] + hist[(k-25)%55]) % 1000000
			hist[id] = caller
		}
		k++
		if k <= 55 {
			called = hist[k-1]
		} else {
			id := (k - 1) % 55
			called = (hist[id] + hist[(k-25)%55]) % 1000000
			hist[id] = called
		}
		k++
		if caller != called {
			res++
			r1, r2 := find(caller), find(called)
			if r1 != r2 {
				if size[r1] < size[r2] {
					r1, r2 = r2, r1
				}
				par[r2] = r1
				size[r1] += size[r2]
			}
		}
	}
	fmt.Print(res)
}
