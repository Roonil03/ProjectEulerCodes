package main

import (
	"fmt"
	"math"
)

func sieve(n int) []int {
	isP := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isP[i] = true
	}
	for p := 2; p*p <= n; p++ {
		if isP[p] {
			for q := p * p; q <= n; q += p {
				isP[q] = false
			}
		}
	}
	ps := []int{}
	for i := 2; i <= n; i++ {
		if isP[i] {
			ps = append(ps, i)
		}
	}
	return ps
}

func isPrime(n int, ps []int) bool {
	if n < 2 {
		return false
	}
	r := int(math.Sqrt(float64(n)))
	for _, p := range ps {
		if p > r {
			break
		}
		if n%p == 0 {
			return false
		}
	}
	return true
}

func combs(n, k int) [][]int {
	var res [][]int
	var dfs func(start int, cur []int)
	dfs = func(start int, cur []int) {
		if len(cur) == k {
			tmp := make([]int, k)
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		for i := start; i < n; i++ {
			dfs(i+1, append(cur, i))
		}
	}
	dfs(0, []int{})
	return res
}

func main() {
	const N = 10
	ps := sieve(int(math.Sqrt(math.Pow10(N))) + 1)
	total := 0
	for d := 0; d <= 9; d++ {
		for rep := N; rep >= 0; rep-- {
			sumD := 0
			masks := combs(N, N-rep)
			for _, m := range masks {
				max := int(math.Pow(9, float64(N-rep)))
				for x := 0; x < max; x++ {
					num := make([]int, N)
					for i := 0; i < N; i++ {
						num[i] = d
					}
					tmp := x
					//valid := true
					for i := 0; i < N-rep; i++ {
						v := tmp % 9
						if v >= d {
							v++
						}
						tmp /= 9
						pos := m[i]
						num[pos] = v
					}
					if num[0] == 0 {
						continue
					}
					n := 0
					for i := 0; i < N; i++ {
						n = n*10 + num[i]
					}
					if isPrime(n, ps) {
						sumD += n
					}
				}
			}
			if sumD > 0 {
				total += sumD
				break
			}
		}
	}
	fmt.Println("The answer is: ", total)
}
