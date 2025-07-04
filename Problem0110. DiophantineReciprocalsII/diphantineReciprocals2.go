package main

import "fmt"

var p = []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}

const tgt = 8000000

var best = ^uint64(0)

func dfs(i, prev int, n, d uint64) {
	if d >= tgt {
		if n < best {
			best = n
		}
		return
	}
	if i >= len(p) {
		return
	}
	for e := 1; e <= prev; e++ {
		n *= p[i]
		if n >= best {
			return
		}
		d2 := d * (2*uint64(e) + 1)
		dfs(i+1, e, n, d2)
	}
}

func main() {
	dfs(0, 60, 1, 1)
	fmt.Println("The answer is:", best)
}
