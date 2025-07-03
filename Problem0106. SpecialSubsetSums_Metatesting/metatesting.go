package main

import "fmt"

func main() {
	n, ans := 12, 0
	var ss [7][]int
	for m := 1; m < 1<<n; m++ {
		c := 0
		for x := m; x > 0; x &= x - 1 {
			c++
		}
		if c <= n/2 {
			ss[c] = append(ss[c], m)
		}
	}
	for k := 1; k <= n/2; k++ {
		s := ss[k]
		for i := 0; i < len(s); i++ {
			for j := i + 1; j < len(s); j++ {
				a, b := s[i], s[j]
				if a&b != 0 {
					continue
				}
				ok1, ok2 := true, true
				ai, bi := a, b
				for p := 0; p < k; p++ {
					var xi, yi int
					for xi = 0; ai&(1<<xi) == 0; xi++ {
					}
					for yi = 0; bi&(1<<yi) == 0; yi++ {
					}
					if xi >= yi {
						ok1 = false
					}
					if yi >= xi {
						ok2 = false
					}
					ai &= ^(1 << xi)
					bi &= ^(1 << yi)
				}
				if !(ok1 || ok2) {
					ans++
				}
			}
		}
	}
	fmt.Println("The answer is: ", ans)
}
