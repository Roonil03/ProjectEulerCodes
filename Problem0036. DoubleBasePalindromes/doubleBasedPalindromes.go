package main

import (
	"fmt"
	"strconv"
)

func main() {
	lim := 1000000
	res := 0
	for i := 0; i < lim; i++ {
		if isPalin10(i) && isPalin2(i) {
			res += i
		}
	}
	fmt.Print("The answer is: ", res)
}

func isPalin10(n int) bool {
	og, rev := n, 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return og == rev
}

func isPalin2(n int) bool {
	bin := strconv.FormatInt(int64(n), 2)
	l := len(bin)
	for i := 0; i < l/2; i++ {
		if bin[i] != bin[l-1-i] {
			return false
		}
	}
	return true
}
