package main

import (
	"fmt"
	"math"
)

func main() {
	tar := 2000000
	zx := 2000
	tn := make([]int, zx)
	for i := 0; i < zx; i++ {
		n := i + 1
		tn[i] = n * (n + 1) / 2
	}
	r := 0
	for i := 0; i < zx; i++ {
		if tn[i] >= tar {
			r = i
			break
		}
	}
	l := 0
	res := 0
	mn := math.MaxInt32
	for l <= r {
		count := tn[l] * tn[r]
		diff := int(math.Abs(float64(count - tar)))
		area := (l + 1) * (r + 1)
		if diff < mn || (diff == mn && area > res) {
			res = area
			mn = diff
		}
		if count > tar {
			r--
		} else {
			l++
		}
	}
	fmt.Println("The answer is:", res)
}
