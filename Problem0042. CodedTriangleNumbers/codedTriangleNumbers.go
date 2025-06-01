package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func isTriangular(t int) bool {
	discriminant := 1 + 8*t
	sqrtDiscriminant := math.Sqrt(float64(discriminant))
	intSqrt := int(sqrtDiscriminant + 0.5)
	if intSqrt*intSqrt != discriminant {
		return false
	}
	return (intSqrt-1)%2 == 0
}

func wordValue(word string) int {
	sum := 0
	for _, char := range word {
		if char >= 'A' && char <= 'Z' {
			sum += int(char - 'A' + 1)
		}
	}
	return sum
}

func main() {
	filename := "Problem0042. CodedTriangleNumbers/input.txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	content := strings.ReplaceAll(string(data), "\"", "")
	words := strings.Split(content, ",")
	count := 0
	for _, word := range words {
		word = strings.TrimSpace(word)
		value := wordValue(word)
		if isTriangular(value) {
			count++
		}
	}
	fmt.Println(count)
}
