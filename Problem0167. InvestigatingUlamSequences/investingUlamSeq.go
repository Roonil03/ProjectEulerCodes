package main

import "fmt"

func solve(n int) int64 {
	L := 2*n + 2
	var S0 uint32 = ((1 << (n + 2)) - 1) << n
	S := S0
	var P int64 = 0
	var W int64 = 0
	for {
		if (S & 1) == 1 {
			W++
		}
		newBit := ((S >> (L - 1)) & 1) ^ (S & 1)
		S = (S >> 1) | (newBit << (L - 1))
		P++
		if S == S0 {
			break
		}
	}
	R := int64(100000000000) - 2
	C := (R - 1) / W
	rem := R - C*W
	S = S0
	var ones int64 = 0
	var targetIdx int64 = 0
	for idx := int64(0); idx < P; idx++ {
		if (S & 1) == 1 {
			ones++
		}
		if ones == rem {
			targetIdx = C*P + idx
			break
		}
		newBit := ((S >> (L - 1)) & 1) ^ (S & 1)
		S = (S >> 1) | (newBit << (L - 1))
	}
	return 2*targetIdx + 1
}

func main() {
	var ans int64 = 0
	for n := 2; n <= 10; n++ {
		ans += solve(n)
	}
	fmt.Println(ans)
}
