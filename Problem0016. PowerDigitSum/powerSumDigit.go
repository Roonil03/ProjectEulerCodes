package main

import (
	"fmt"
	"math/big"
)

func main() {
	res := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)
	r := res.String()
	sum := 0
	for _, i := range r {
		sum += int(i - '0')
	}
	fmt.Println("The result is: ", sum)
}
