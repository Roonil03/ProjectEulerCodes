package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	mem := make(map[int]bool)
	for a := 1; a <= 9; a++ {
		for b := 1000; b <= 9999; b++ {
			if uniq(b) {
				c := a * b
				s := fmt.Sprintf("%d%d%d", a, b, c)
				if panDigi(s) {
					mem[c] = true
				}
			}
		}
	}
	for a := 10; a <= 99; a++ {
		if uniq(a) {
			for b := 100; b <= 999; b++ {
				if uniq(b) {
					c := a * b
					s := fmt.Sprintf("%d%d%d", a, b, c)
					if panDigi(s) {
						mem[c] = true
					}
				}
			}
		}
	}
	sum := 0
	for p := range mem {
		sum += p
	}
	fmt.Println(sum)
}

func uniq(n int) bool {
	s := strconv.Itoa(n)
	if strings.ContainsRune(s, '0') {
		return false
	}
	seen := make(map[rune]bool)
	for _, c := range s {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

func panDigi(s string) bool {
	if len(s) != 9 || strings.ContainsRune(s, '0') {
		return false
	}
	seen := make(map[rune]bool)
	for _, c := range s {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return len(seen) == 9
}
