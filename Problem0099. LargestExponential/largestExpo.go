package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := "Problem0099. LargestExponential/input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	maxVal := -1.0
	res := 0
	lineNum := 0
	for sc.Scan() {
		lineNum++
		parts := strings.Split(sc.Text(), ",")
		if len(parts) != 2 {
			log.Fatalf("invalid format on line %d", lineNum)
		}
		base, err1 := strconv.ParseFloat(parts[0], 64)
		exp, err2 := strconv.ParseFloat(parts[1], 64)
		if err1 != nil || err2 != nil {
			log.Fatalf("parse error on line %d", lineNum)
		}
		v := exp * math.Log(base)
		if v > maxVal {
			maxVal = v
			res = lineNum
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("The answer is:", res)
}
