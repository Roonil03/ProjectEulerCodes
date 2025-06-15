package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	precision := uint(3500)
	thirteen := new(big.Float).SetPrec(precision).SetInt64(13)
	result := sqrtNewton(thirteen, precision)
	sum := sumFractionalDigits(result, 1000)
	fmt.Printf("S(13, 1000) = %d\n", sum)
}

func sqrtNewton(x *big.Float, precision uint) *big.Float {
	guess := new(big.Float).SetPrec(precision).SetFloat64(3.6)
	epsilon := new(big.Float).SetPrec(precision)
	epsilon.SetString("1e-1010")
	two := new(big.Float).SetPrec(precision).SetInt64(2)
	temp1 := new(big.Float).SetPrec(precision)
	temp2 := new(big.Float).SetPrec(precision)
	for i := 0; i < 50; i++ {
		temp1.Quo(x, guess)
		temp2.Add(guess, temp1)
		temp2.Quo(temp2, two)
		temp1.Sub(temp2, guess)
		temp1.Abs(temp1)
		if temp1.Cmp(epsilon) < 0 {
			break
		}

		guess.Set(temp2)
	}
	return guess
}

func sumFractionalDigits(sqrt *big.Float, numDigits int) int {
	str := sqrt.Text('f', numDigits+10)
	dotIndex := strings.Index(str, ".")
	if dotIndex == -1 {
		return 0
	}
	fractional := str[dotIndex+1:]
	sum := 0
	for i := 0; i < numDigits && i < len(fractional); i++ {
		if fractional[i] >= '0' && fractional[i] <= '9' {
			sum += int(fractional[i] - '0')
		}
	}
	return sum
}
