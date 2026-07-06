package main

import "fmt"

func main() {
	const N = 100
	const K = 50
	sq := make([]int, N+1)
	for i := 1; i <= N; i++ {
		sq[i] = i * i
	}
	minS := make([]int, K+1)
	for j := 1; j <= K; j++ {
		for k := 1; k <= j; k++ {
			minS[j] += sq[k]
		}
	}
	maxS := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		maxS[i] = make([]int, K+1)
		for j := 1; j <= K; j++ {
			if j <= i {
				for k := i - j + 1; k <= i; k++ {
					maxS[i][j] += sq[k]
				}
			}
		}
	}
	maxSum := maxS[N][K]
	dp := make([][]int32, K+1)
	for i := range dp {
		dp[i] = make([]int32, maxSum+1)
	}
	dp[0][0] = 1
	for i := 1; i <= N; i++ {
		v := sq[i]
		lim := min(i, K)
		for j := lim; j >= 1; j-- {
			low := minS[j-1] + v
			high := maxS[i-1][j-1] + v
			for s := high; s >= low; s-- {
				if dp[j-1][s-v] > 0 {
					val := min(dp[j][s]+dp[j-1][s-v], 2)
					dp[j][s] = val
				}
			}
		}
	}
	var res int64
	for s := 0; s <= maxSum; s++ {
		if dp[K][s] == 1 {
			res += int64(s)
		}
	}
	fmt.Println(res)
}
