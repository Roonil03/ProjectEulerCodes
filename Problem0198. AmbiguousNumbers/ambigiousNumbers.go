package main

import "fmt"

func main() {
	var dfs func(a, b, c, d int64) int64
	var lim int64 = 50000000
	dfs = func(a, b, c, d int64) int64 {
		if b*d > lim || 100*a >= b {
			return 0
		}
		if 100*(a*d+b*c) < 2*b*d {
			temp := (lim - b*d) / (b * b)
			res := 1 + temp + dfs(a+c, b+d, c, d)
			for k := int64(1); k <= temp; k++ {
				if ((k+1)*b+d)*(k*b+d) > lim {
					break
				}
				res += dfs((k+1)*a+c, (k+1)*b+d, k*a+c, k*b+d)
			}
			return res
		}
		k := (100*(a*d+b*c)-2*b*d)/(2*b*(b-100*a)) + 1
		if b*(k*b+d) <= lim {
			return dfs(a, b, k*a+c, k*b+d)
		}
		return 0
	}
	fmt.Println(dfs(0, 1, 1, 1))
}
