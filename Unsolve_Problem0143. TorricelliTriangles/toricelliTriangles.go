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
	seen := make([]bool, L+1)
	for m := 2; ; m++ {
		for n := 1; n < m; n++ {
			if gcd(m, n) != 1 || (m-n)%2 == 0 || (m-n)%3 == 0 {
				continue
			}
			a := 3*m*m - 2*m*n - n*n
			b := 3*m*m + 2*m*n - n*n
			if a <= 0 || a%4 != 0 || b%4 != 0 {
				continue
			}
			p := m * n
			q := a / 4
			r := b / 4
			s := p + q + r
			if s > L {
				break
			}
			for k := 1; k*s <= L; k++ {
				seen[k*s] = true
			}
		}
		if 2*m+1 > L {
			break
		}
	}
	ans := 0
	for s, ok := range seen {
		if ok {
			ans += s
		}
	}
	fmt.Println(ans)
}
