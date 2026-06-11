package main

import (
	"fmt"
	"strconv"
	"strings"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	p, q := int64(123456789), int64(987654321)
	g := gcd(p, q)
	p, q = p/g, q/g
	var r []int64
	for p > 0 && q > 0 {
		if p > q {
			if p%q == 0 {
				r = append(r, p/q-1, 1)
				break
			}
			r = append(r, p/q)
			p %= q
		} else if q > p {
			if q%p == 0 {
				r = append(r, q/p)
				break
			}
			r = append(r, q/p)
			q %= p
		} else {
			r = append(r, 1)
			break
		}
	}
	var a []string
	for i := len(r) - 1; i >= 0; i-- {
		a = append(a, strconv.FormatInt(r[i], 10))
	}
	fmt.Println(strings.Join(a, ","))
}
