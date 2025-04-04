package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "Problem0018. MaximumPathSumI/input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error with file")
		return
	}
	defer file.Close()
	var triangle [][]int
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		row := make([]int, len(fields))
		for i, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error with file")
				return
			}
			row[i] = num
		}
		triangle = append(triangle, row)
	}
	if err := sc.Err(); err != nil {
		fmt.Println("Error with file")
		return
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	fmt.Println("Maximum path sum:", triangle[0][0])
}
