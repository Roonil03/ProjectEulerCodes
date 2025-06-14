package main

import (
	"fmt"
	"sort"
	"strconv"
)

func digitFingerprint(n int64) string {
	digits := []rune(strconv.FormatInt(n, 10))
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] < digits[j]
	})
	return string(digits)
}

func main() {
	groups := make(map[string][]int64)
	for i := int64(1); ; i++ {
		cube := i * i * i
		fingerprint := digitFingerprint(cube)
		groups[fingerprint] = append(groups[fingerprint], cube)
		if len(groups[fingerprint]) == 5 {
			min := groups[fingerprint][0]
			for _, c := range groups[fingerprint] {
				if c < min {
					min = c
				}
			}
			fmt.Println("The answer is: ", min)
			return
		}
	}
}
