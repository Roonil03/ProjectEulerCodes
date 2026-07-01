package main

import "fmt"

func main() {
	const M = 5200000
	isp := make([]uint64, (M+64)/64)
	for i := range isp {
		isp[i] = ^uint64(0)
	}
	isp[0] &= ^uint64(3)
	var P []int64
	for i := int64(2); i <= M; i++ {
		if (isp[i>>6] & (uint64(1) << (i & 63))) != 0 {
			P = append(P, i)
			for j := i * i; j <= M; j += i {
				isp[j>>6] &= ^(uint64(1) << (j & 63))
			}
		}
	}
	solve := func(n int64) int64 {
		st := (n-3)*(n-2)/2 + 1
		en := (n + 2) * (n + 3) / 2
		sz := en - st + 1
		is := make([]uint64, (sz+63)/64)
		for i := range is {
			is[i] = ^uint64(0)
		}
		for _, p := range P {
			if p*p > en {
				break
			}
			f := max(((st+p-1)/p)*p, p*p)
			for j := f; j <= en; j += p {
				ix := j - st
				is[ix>>6] &= ^(uint64(1) << (ix & 63))
			}
		}
		cN := func(r, c int64) (int, int64, int64) {
			count := 0
			var lr, lc int64
			for dr := int64(-1); dr <= 1; dr++ {
				for dc := int64(-1); dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					}
					nr, nc := r+dr, c+dc
					if nc >= 1 && nc <= nr {
						ix := nr*(nr-1)/2 + nc - st
						if (is[ix>>6] & (uint64(1) << (ix & 63))) != 0 {
							count++
							lr, lc = nr, nc
						}
					}
				}
			}
			return count, lr, lc
		}
		var sum int64 = 0
		for c := int64(1); c <= n; c++ {
			ix := n*(n-1)/2 + c - st
			if (is[ix>>6] & (uint64(1) << (ix & 63))) != 0 {
				count, lr, lc := cN(n, c)
				if count >= 2 {
					sum += n*(n-1)/2 + c
				} else if count == 1 {
					count2, _, _ := cN(lr, lc)
					if count2 >= 2 {
						sum += n*(n-1)/2 + c
					}
				}
			}
		}
		return sum
	}
	fmt.Print(solve(5678027) + solve(7208785))
}
