package main

import (
	"fmt"
	"strconv"
)

func isPal(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	const L = 100000000
	m := make(map[int]bool)
	for i := 1; i*i < L; i++ {
		s := i * i
		for j := i + 1; s < L; j++ {
			s += j * j
			if s >= L {
				break
			}
			if isPal(s) {
				m[s] = true
			}
		}
	}
	r := 0
	for k := range m {
		r += k
	}
	fmt.Println("The answer is:", r)
}
