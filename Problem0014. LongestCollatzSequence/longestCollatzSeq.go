package main

import "fmt"

var memo map[int]int

func main() {
	memo = make(map[int]int)
	var resNum, resLen int = 0, 0
	n := 1000000
	for i := 1; i < n; i++ {
		l := collatzLength(i)
		if l > resLen {
			resLen = l
			resNum = i
		}
	}
	fmt.Println("Sequence Length: ", resLen, " and number was: ", resNum)
}

func collatzLength(a int) int {
	if a == 1 {
		return 1
	}
	if l, err := memo[a]; err {
		return l
	}
	var n int
	if a%2 == 0 {
		n = a / 2
	} else {
		n = 3*a + 1
	}
	length := 1 + collatzLength(n)
	memo[a] = length
	return length
}
