package main

import (
	"fmt"
	"math/bits"
)

var squares = [][2]int{
	{0, 1}, {0, 4}, {0, 9}, {1, 6}, {2, 5},
	{3, 6}, {4, 9}, {6, 4}, {8, 1},
}

func main() {
	var combos []int
	for mask := 0; mask < 1<<10; mask++ {
		if bits.OnesCount(uint(mask)) == 6 {
			combos = append(combos, mask)
		}
	}
	pres := make([][10]bool, len(combos))
	for i, mask := range combos {
		for d := 0; d < 10; d++ {
			if mask&(1<<d) != 0 {
				pres[i][d] = true
			}
		}
	}
	contains := func(i, d int) bool {
		if d == 6 || d == 9 {
			return pres[i][6] || pres[i][9]
		}
		return pres[i][d]
	}
	count := 0
	n := len(combos)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			valid := true
			for _, p := range squares {
				a, b := p[0], p[1]
				if !((contains(i, a) && contains(j, b)) || (contains(i, b) && contains(j, a))) {
					valid = false
					break
				}
			}
			if valid {
				count++
			}
		}
	}
	fmt.Println("The answer is:", count)
}
