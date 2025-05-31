package main

import (
	"fmt"
	"strconv"
)

func isPandigital(s string) bool {
	if len(s) != 9 {
		return false
	}
	seen := [10]bool{}
	for _, c := range s {
		digit := int(c - '0')
		if digit == 0 || seen[digit] {
			return false
		}
		seen[digit] = true
	}
	return true
}

func main() {
	maxPandigital := 0
	for x := 1; x <= 9999; x++ {
		concatenated := ""
		multiplier := 1
		for len(concatenated) < 9 {
			concatenated += strconv.Itoa(x * multiplier)
			multiplier++
		}
		if len(concatenated) == 9 && multiplier > 2 && isPandigital(concatenated) {
			if val, _ := strconv.Atoi(concatenated); val > maxPandigital {
				maxPandigital = val
			}
		}
	}
	fmt.Println(maxPandigital)
}
