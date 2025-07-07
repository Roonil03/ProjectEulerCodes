package main

import "fmt"

func main() {
	var used [9]bool
	gen(0, 0, &used)
	all := (1 << 9) - 1
	fmt.Println("The answer is:", dfs(all))
}

var ps []int

func init() {
	const limit = 100000
	bs := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		if !bs[i] {
			ps = append(ps, i)
			for j := i * i; j <= limit; j += i {
				bs[j] = true
			}
		}
	}
}

func isP(n int) bool {
	if n < 2 {
		return false
	}
	for _, p := range ps {
		if p*p > n {
			break
		}
		if n%p == 0 {
			return n == p
		}
	}
	return true
}

var cnt [1 << 9]int

func gen(m, cur int, used *[9]bool) {
	if cur > 0 && isP(cur) {
		cnt[m]++
	}
	for d := 0; d < 9; d++ {
		if used[d] {
			continue
		}
		used[d] = true
		gen(m|(1<<d), cur*10+(d+1), used)
		used[d] = false
	}
}

var memo [1 << 9]int

func dfs(m int) int {
	if m == 0 {
		return 1
	}
	if memo[m] > 0 {
		return memo[m]
	}
	total := 0
	lo := m & -m
	for s := m; s > 0; s = (s - 1) & m {
		if s&lo == 0 {
			continue
		}
		if cnt[s] > 0 {
			total += cnt[s] * dfs(m^s)
		}
	}
	memo[m] = total
	return total
}
