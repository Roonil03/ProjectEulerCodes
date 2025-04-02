package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("Problem0013. LargeSum/input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sum := big.NewInt(0)
	count := 0
	n := 100
	for sc.Scan() && count < n {
		str := strings.TrimSpace(sc.Text())
		if len(str) != 50 {
			fmt.Println("Warning!")
			continue
		}
		num := new(big.Int)
		num.SetString(str, 10)
		sum.Add(sum, num)
		count++
	}
	if err := sc.Err(); err != nil {
		fmt.Println("Error in reading file ", err)
	}
	fmt.Println("Sum is: ", ((sum.String())[:10]))
}
