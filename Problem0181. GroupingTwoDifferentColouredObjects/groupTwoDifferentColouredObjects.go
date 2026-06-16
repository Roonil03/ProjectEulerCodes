package main

import (
	"fmt"
	"math/big"
)

func main() {
	dp := make([]big.Int, 61*41)
	dp[0].SetInt64(1)
	for i := range 61 {
		for j := range 41 {
			if i == 0 && j == 0 {
				continue
			}
			for x := i; x <= 60; x++ {
				for y := j; y <= 40; y++ {
					dp[x*41+y].Add(&dp[x*41+y], &dp[(x-i)*41+(y-j)])
				}
			}
		}
	}
	fmt.Println(&dp[60*41+40])
}
