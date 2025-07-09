package main

import (
	"fmt"
	"math/big"
)

func main() {
	const n = 15
	dp := make([]*big.Int, n+1)
	for i := range dp {
		dp[i] = big.NewInt(0)
	}
	dp[0].SetInt64(1)
	for k := 1; k <= n; k++ {
		next := make([]*big.Int, n+1)
		for i := range next {
			next[i] = big.NewInt(0)
		}
		for j := 0; j <= k; j++ {
			if j <= k-1 {
				temp := new(big.Int).Mul(dp[j], big.NewInt(int64(k)))
				next[j].Add(next[j], temp)
			}
			if j > 0 {
				next[j].Add(next[j], dp[j-1])
			}
		}
		dp = next
	}
	win := big.NewInt(0)
	for j := n/2 + 1; j <= n; j++ {
		win.Add(win, dp[j])
	}
	total := big.NewInt(1)
	for i := 2; i <= n+1; i++ {
		total.Mul(total, big.NewInt(int64(i)))
	}
	prize := new(big.Int).Div(total, win)
	fmt.Println("The answer is:", prize)
}
