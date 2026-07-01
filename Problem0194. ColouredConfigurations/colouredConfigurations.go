package main

import "fmt"

func main() {
	a, b, cc, mo := 25, 75, 1984, 100000000
	t := make([]int, 8)
	for c := 2; c <= 7; c++ {
		v := 1
		for i := 2; i < c; i++ {
			v = v * (cc - i) / (i - 1)
		}
		t[c] = v % mo
	}
	f, g := 0, 0
	for c := 2; c <= 7; c++ {
		tc, exp := t[c], (1<<(c+1))-2
		for c2 := 2; c2 <= c; c2++ {
			for c3 := 2; c3 <= c; c3++ {
				if c2 == c3 {
					continue
				}
				for c7 := 1; c7 <= c; c7++ {
					if c2 == c7 {
						continue
					}
					for c4 := 1; c4 <= c; c4++ {
						if c3 == c4 {
							continue
						}
						for c5 := 1; c5 <= c; c5++ {
							if c5 == 2 || c4 == c5 || c5 == c7 {
								continue
							}
							if 6|(1<<c2)|(1<<c3)|(1<<c4)|(1<<c5)|(1<<c7) == exp {
								if c7 != 2 {
									f = (f + tc) % mo
								}
								g = (g + tc) % mo
							}
						}
					}
				}
			}
		}
	}
	cb := make([][]int, a+b+1)
	for i := 0; i <= a+b; i++ {
		cb[i] = make([]int, a+b+1)
		cb[i][0] = 1
		for j := 1; j <= i; j++ {
			cb[i][j] = (cb[i-1][j-1] + cb[i-1][j]) % mo
		}
	}
	res := cb[a+b][b]
	for range a {
		res = (res * f) % mo
	}
	for range b {
		res = (res * g) % mo
	}
	fmt.Printf("%08d\n", res*cc%mo*(cc-1)%mo)
}
