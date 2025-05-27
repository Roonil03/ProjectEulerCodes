package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var validNums []int
	for n := 11; n <= 99; n++ {
		tens := n / 10
		units := n % 10
		if tens != 0 && units != 0 {
			validNums = append(validNums, n)
		}
	}
	var numerators, denominators []int
	for i := 0; i < len(validNums); i++ {
		num := validNums[i]
		for j := i + 1; j < len(validNums); j++ {
			den := validNums[j]
			numTens := num / 10
			numUnits := num % 10
			denTens := den / 10
			denUnits := den % 10
			digitsNum := []int{numTens, numUnits}
			digitsDen := []int{denTens, denUnits}
			for iDigit := 0; iDigit < 2; iDigit++ {
				for jDigit := 0; jDigit < 2; jDigit++ {
					if digitsNum[iDigit] == digitsDen[jDigit] && digitsNum[iDigit] != 0 {
						newNum := digitsNum[1-iDigit]
						newDen := digitsDen[1-jDigit]
						if newDen == 0 {
							continue
						}
						if num*newDen == den*newNum {
							numerators = append(numerators, num)
							denominators = append(denominators, den)
						}
					}
				}
			}
		}
	}
	productNum, productDen := 1, 1
	for i := range numerators {
		productNum *= numerators[i]
		productDen *= denominators[i]
	}
	commonGCD := gcd(productNum, productDen)
	productDen /= commonGCD

	fmt.Println("The denominator of the product in lowest terms:", productDen)
}
