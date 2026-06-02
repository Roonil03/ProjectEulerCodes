package main

import "fmt"

func main() {
	const W, H = 9, 12
	dp := make([]int, 1<<(2*W))
	temp := make([]int, 1<<(2*W))
	dp[0] = 1
	for i := range W * H {
		for j := range temp {
			temp[j] = 0
		}
		r, c := i/W, i%W
		for id, v := range dp {
			if v == 0 {
				continue
			}
			if id&1 != 0 {
				temp[id>>1] += v
			} else {
				if c <= W-3 && id&6 == 0 {
					temp[(id>>1)|3] += v
				}
				if r <= H-3 && id&(1<<W) == 0 {
					temp[(id>>1)|(1<<(W-1))|(1<<(2*W-1))] += v
				}
				if c <= W-2 && r <= H-2 && id&((1<<1)|(1<<W)) == 0 {
					temp[(id>>1)|1|(1<<(W-1))] += v
				}
				if c <= W-2 && r <= H-2 && id&((1<<1)|(1<<(W+1))) == 0 {
					temp[(id>>1)|1|(1<<W)] += v
				}
				if c <= W-2 && r <= H-2 && id&((1<<W)|(1<<(W+1))) == 0 {
					temp[(id>>1)|(1<<(W-1))|(1<<W)] += v
				}
				if c >= 1 && r <= H-2 && id&((1<<(W-1))|(1<<W)) == 0 {
					temp[(id>>1)|(1<<(W-2))|(1<<(W-1))] += v
				}
			}
		}
		dp, temp = temp, dp
	}
	fmt.Println(dp[0])
}
