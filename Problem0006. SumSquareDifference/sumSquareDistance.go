package main

import (
	"fmt"
	"math"
)

func main() {
	n := int64(100)
	a := int64(math.Pow(float64(int64(n*(n+int64(1))/int64(2))), float64(2)))
	b := int64(n * (n + int64(1)) * (int64(2)*n + int64(1)) / int64(6))
	fmt.Print(b, " ", a, " ", a-b)
}
