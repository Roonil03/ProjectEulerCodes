package main

import (
	"fmt"
	"sort"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type Pair struct {
	u, v int
}

const limit = 120000

func main() {
	l := 1
	for l*l <= limit {
		l++
	}
	pairs := make([]Pair, 0, 150000)
	for m := 2; m <= l; m++ {
		for n := 1; n < m; n++ {
			if gcd(m, n) != 1 || (m-n)%3 == 0 {
				continue
			}
			x := n*n + 2*m*n
			y := m*m - n*n
			if x > y {
				x, y = y, x
			}
			if x+y > limit {
				break
			}
			for k := 1; k*(x+y) <= limit; k++ {
				pairs = append(pairs, Pair{k * x, k * y})
			}
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].u != pairs[j].u {
			return pairs[i].u < pairs[j].u
		}
		return pairs[i].v < pairs[j].v
	})
	uniq := pairs[:0]
	for _, p := range pairs {
		if len(uniq) == 0 || uniq[len(uniq)-1] != p {
			uniq = append(uniq, p)
		}
	}
	adj := make([][]int, limit+1)
	for _, p := range uniq {
		adj[p.u] = append(adj[p.u], p.v)
	}
	seen := make([]bool, limit+1)
	for a := 1; a <= limit; a++ {
		na := adj[a]
		if len(na) == 0 {
			continue
		}
		for id, b := range na {
			if a+2*b > limit {
				break
			}
			nb := adj[b]
			i := max(sort.Search(len(na), func(t int) bool { return na[t] > b }), id+1)
			j := 0
			for i < len(na) && j < len(nb) {
				x, y := na[i], nb[j]
				if x == y {
					sum := a + b + x
					if sum <= limit {
						seen[sum] = true
					}
					i++
					j++
				} else if x < y {
					i++
				} else {
					j++
				}
			}
		}
	}
	res := 0
	for s := 0; s <= limit; s++ {
		if seen[s] {
			res += s
		}
	}
	fmt.Println(res)
}
