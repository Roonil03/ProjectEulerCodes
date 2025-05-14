package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := 100
	fmt.Println("The result is:", factSum(n))
}

func factSum(n int) int {
	fact := big.NewInt(1)
	for i := 2; i <= n; i++ {
		fact.Mul(fact, big.NewInt(int64(i)))
	}
	str := fact.String()
	sum := 0
	for _, ch := range str {
		sum += int(ch - '0')
	}
	return sum
}
