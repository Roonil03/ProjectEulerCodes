package main

import (
	"fmt"
)

const N = 7

var bS [N]int
var bM = 255

func cR2(p []int) bool {
	m := len(p)
	for k := 1; k <= m/2; k++ {
		s, t := 0, 0
		for i := 0; i <= k; i++ {
			s += p[i]
		}
		for i := 0; i < k; i++ {
			t += p[m-1-i]
		}
		if s <= t {
			return false
		}
	}
	return true
}

func cUS(c []int) bool {
	m := 1 << N
	seen := make(map[int]bool, m)
	for mask := 1; mask < m; mask++ {
		s := 0
		for i := 0; i < N; i++ {
			if mask&(1<<i) != 0 {
				s += c[i]
			}
		}
		if seen[s] {
			return false
		}
		seen[s] = true
	}
	return true
}

func dfs(p []int, st, sf int) {
	if len(p) == N {
		if cUS(p) && sf <= bM {
			bM = sf
			copy(bS[:], p)
		}
		return
	}
	need := N - len(p)
	minTail := need * (2*st + need - 1) / 2
	if sf+minTail >= bM {
		return
	}
	for x := st; ; x++ {
		if sf+x+(need-1)*(x+1) >= bM {
			break
		}
		p = append(p, x)
		if cR2(p) {
			dfs(p, x+1, sf+x)
		}
		p = p[:len(p)-1]
	}
}

func main() {
	dfs([]int{}, 20, 0)
	fmt.Print("The answer is: ")
	for _, v := range bS {
		fmt.Printf("%d", v)
	}
	fmt.Println()
}
