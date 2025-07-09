package main

import "fmt"

func main() {
	const lim = 20000
	const tgt = 1000
	cnt := make([]int, lim+1)
	for a := 1; a <= lim; a++ {
		if layer(a, 1, 1, 1) > lim {
			break
		}
		for b := 1; b <= a; b++ {
			if layer(a, b, 1, 1) > lim {
				break
			}
			for c := 1; c <= b; c++ {
				for k := 1; ; k++ {
					v := layer(a, b, c, k)
					if v > lim {
						break
					}
					cnt[v]++
				}
			}
		}
	}
	for i := 1; i <= lim; i++ {
		if cnt[i] == tgt {
			fmt.Println("The answer is:", i)
			return
		}
	}
}

func layer(a, b, c, k int) int {
	return 2*(a*b+b*c+c*a) + 4*(k-1)*(a+b+c) + (4*k*k - 12*k + 8)
}
