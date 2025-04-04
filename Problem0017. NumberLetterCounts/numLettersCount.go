package main

import (
	"fmt"
	"strings"
)

func main() {
	total := 0
	n := 1000
	for i := 1; i <= n; i++ {
		word := numberToWords(i)
		word = strings.ReplaceAll(word, " ", "")
		word = strings.ReplaceAll(word, "-", "")
		total += len(word)
	}
	fmt.Println("Total letters:", total)
}

func numberToWords(n int) string {
	if n == 0 {
		return "zero"
	}
	ones := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	if n < 20 {
		return ones[n]
	} else if n < 100 {
		if n%10 == 0 {
			return tens[n/10]
		}
		return tens[n/10] + "-" + ones[n%10]
	} else if n < 1000 {
		if n%100 == 0 {
			return ones[n/100] + " hundred"
		}
		return ones[n/100] + " hundred and " + numberToWords(n%100)
	} else if n == 1000 {
		return "one thousand"
	}
	return "number out of range"
}
