package main

import "fmt"

func main() {
	const n = 1000
	var t uint64
	var tri [n][]int64
	var ps [n][]int64
	k := 0
	for i := 0; i < n; i++ {
		tri[i] = make([]int64, i+1)
		ps[i] = make([]int64, i+2)
		for j := 0; j <= i; j++ {
			t = (615949*t + 797807) & ((1 << 20) - 1)
			tri[i][j] = int64(t) - (1 << 19)
			ps[i][j+1] = ps[i][j] + tri[i][j]
			k++
		}
	}
	var mn int64
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			var s int64
			for h := 1; h <= n-i; h++ {
				if j+h-1 > i+h-1 {
					break
				}
				s += ps[i+h-1][j+h] - ps[i+h-1][j]
				if s < mn {
					mn = s
				}
			}
		}
	}
	fmt.Println("The answer is:", mn)
}
