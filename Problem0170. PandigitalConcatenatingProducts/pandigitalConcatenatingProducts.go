package main

import "fmt"

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func isPan(m, a, b int64) bool {
	var mask, d int
	for _, v := range []int64{m, a, b} {
		for v > 0 {
			mask |= 1 << (v % 10)
			v /= 10
			d++
		}
	}
	return d == 10 && mask == 1023
}

func main() {
	p := []int64{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for {
		var n int64
		for _, v := range p {
			n = n*10 + v
		}
		for i := 1; i < 9; i++ {
			pw := int64(1)
			for j := 0; j < 10-i; j++ {
				pw *= 10
			}
			p1, p2 := n/pw, n%pw
			g := gcd(p1, p2)
			for m := int64(2); m*m <= g; m++ {
				if g%m == 0 {
					if isPan(m, p1/m, p2/m) || isPan(g/m, p1/(g/m), p2/(g/m)) {
						fmt.Println(n)
						return
					}
				}
			}
			if g > 1 && isPan(g, p1/g, p2/g) {
				fmt.Println(n)
				return
			}
		}
		i := 8
		for i >= 0 && p[i] <= p[i+1] {
			i--
		}
		if i < 0 {
			break
		}
		j := 9
		for p[j] >= p[i] {
			j--
		}
		p[i], p[j] = p[j], p[i]
		for l, r := i+1, 9; l < r; l, r = l+1, r-1 {
			p[l], p[r] = p[r], p[l]
		}
	}
}
