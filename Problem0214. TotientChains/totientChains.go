package main

import (
	"bufio"
	"fmt"
	"os"
)

const lim = 40000000

var s = []uint16(make([]uint16, lim))
var d = []uint8(make([]uint8, lim))

func phi(n int) int {
	r := n
	for n > 1 {
		p := int(s[n])
		if p == 0 {
			p = n
		}
		r -= r / p
		for n%p == 0 {
			n /= p
		}
	}
	return r
}

func h1(n int) uint8 {
	if d[n] > 0 {
		return d[n]
	}
	d[n] = h1(phi(n)) + 1
	return d[n]
}

func main() {
	for i := 2; i*i < lim; i++ {
		if s[i] == 0 {
			for j := i * i; j < lim; j += i {
				if s[j] == 0 {
					s[j] = uint16(i)
				}
			}
		}
	}
	d[1] = 1
	res := int64(0)
	for p := 2; p < lim; p++ {
		if s[p] == 0 && h1(p-1)+1 == 25 {
			res += int64(p)
		}
	}
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(w, res)
	w.Flush()
}
