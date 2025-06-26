package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var romanValues = map[rune]int{
	'I': 1, 'V': 5, 'X': 10,
	'L': 50, 'C': 100, 'D': 500,
	'M': 1000,
}
var placeStrings = [][]string{
	{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	{"", "M", "MM", "MMM", "MMMM"},
}

func main() {
	filepath := "Problem0089. RomanNumerals/input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	res := 0
	for sc.Scan() {
		orig := sc.Text()
		n := parseRoman(orig)
		minimal := toMinimalRoman(n)
		res += len(orig) - len(minimal)
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("The answer is:", res)
}

func toMinimalRoman(n int) string {
	digits := [4]int{
		n % 10,
		(n / 10) % 10,
		(n / 100) % 10,
		(n / 1000),
	}
	var m string
	for d := 3; d >= 0; d-- {
		m += placeStrings[d][digits[d]]
	}
	return m
}

func parseRoman(s string) int {
	t, prev := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		val := romanValues[rune(s[i])]
		if val < prev {
			t -= val
		} else {
			t += val
		}
		prev = val
	}
	return t
}
