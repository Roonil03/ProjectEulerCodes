package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "Problem0067. MaximumPathSumII/input.txt"
	triangle := loadTriangle(filename)
	res := maxPathSum(triangle)
	fmt.Println("The answer is:", res)
}

func loadTriangle(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var triangle [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		row := make([]int, len(fields))
		for i, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				log.Fatal(err)
			}
			row[i] = num
		}
		triangle = append(triangle, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return triangle
}

func maxPathSum(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
