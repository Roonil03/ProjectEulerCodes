package main

import "fmt"

func main() {
	const a = 14
	var dp [a + 1][a + 1][a + 1][a + 1][a + 1][5]uint64
	dp[0][0][0][0][0][0] = 1
	for c0 := 0; c0 <= a; c0++ {
		for c1 := 0; c1 <= a; c1++ {
			for c2 := 0; c2 <= a; c2++ {
				for c3 := 0; c3 <= a; c3++ {
					for c4 := 0; c4 <= a; c4++ {
						for d := 0; d < 5; d++ {
							v := dp[c0][c1][c2][c3][c4][d]
							if v == 0 {
								continue
							}
							if d == 0 && c0 < a {
								dp[c0+1][c1][c2][c3][c4][1] += v
							} else if d == 1 && c1 < a {
								dp[c0][c1+1][c2][c3][c4][2] += v
							} else if d == 2 && c2 < a {
								dp[c0][c1][c2+1][c3][c4][3] += v
							} else if d == 3 && c3 < a {
								dp[c0][c1][c2][c3+1][c4][4] += v
							} else if d == 4 && c4 < a {
								dp[c0][c1][c2][c3][c4+1][0] += v
							}
							tR := (d + 4) % 5
							if tR == 0 && c0 < a {
								dp[c0+1][c1][c2][c3][c4][0] += v
							} else if tR == 1 && c1 < a {
								dp[c0][c1+1][c2][c3][c4][1] += v
							} else if tR == 2 && c2 < a {
								dp[c0][c1][c2+1][c3][c4][2] += v
							} else if tR == 3 && c3 < a {
								dp[c0][c1][c2][c3+1][c4][3] += v
							} else if tR == 4 && c4 < a {
								dp[c0][c1][c2][c3][c4+1][4] += v
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(dp[a][a][a][a][a][0])
}
