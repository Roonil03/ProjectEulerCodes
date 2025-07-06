package main

import "fmt"

func main() {
	m, lim := 50, 1000000
	f := make([]int, 1, 2000)
	f[0] = 1
	pref := make([]int, 1, 2000)
	pref[0] = 1
	for n := 1; ; n++ {
		var cnt int
		if n < m {
			cnt = 1
		} else {
			cnt = f[n-1] + pref[n-m]
		}
		f = append(f, cnt)
		pref = append(pref, pref[n-1]+cnt)
		if cnt > lim {
			fmt.Println("The answer is:", n+1)
			return
		}
	}
}
