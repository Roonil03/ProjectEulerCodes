package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	const L = 120000
	r := make([]int, L)
	for i := 1; i < L; i++ {
		r[i] = 1
	}
	for i := 2; i < L; i++ {
		if r[i] == 1 {
			for j := i; j < L; j += i {
				r[j] *= i
			}
		}
	}
	type p struct{ r, n int }
	s := make([]p, 0, L-1)
	for i := 1; i < L; i++ {
		s = append(s, p{r[i], i})
	}
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i].r > s[j].r {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	ans := 0
	for c := 2; c < L; c++ {
		for _, x := range s {
			rad := x.r * r[c]
			if rad >= c {
				break
			}
			a := x.n
			b := c - a
			if a < b && b < L && rad*r[b] < c && gcd(a, b) == 1 {
				ans += c
			}
		}
	}
	fmt.Println("The answer is:", ans)
}
