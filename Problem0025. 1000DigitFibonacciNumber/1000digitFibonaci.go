package main

import (
	"fmt"
	"math/big"
)

func main() {
	f1, f2 := big.NewInt(1), big.NewInt(1)
	id := 2
	for len(f2.String()) < 1000 {
		temp := new(big.Int).Add(f1, f2)
		f1, f2 = f2, temp
		id++
	}
	fmt.Println("The number is: ", f2)
	fmt.Println("The answer is: ", id)
}
