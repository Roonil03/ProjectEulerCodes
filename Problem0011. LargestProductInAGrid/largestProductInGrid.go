package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "Problem0011. LargestProductInAGrid/input.txt"
	grid := readGrid(filename)
	if grid == nil {
		fmt.Println("Error in reading grid")
		return
	}
	res := maxProduct(grid)
	fmt.Println("Result: ", res)
}

func readGrid(filename string) [][]int {
	file, er := os.Open(filename)
	if er != nil {
		return nil
	}
	defer file.Close()
	var grid [][]int
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		numbers := strings.Fields(line)
		row := make([]int, len(numbers))
		for i, num := range numbers {
			val, err := strconv.Atoi(num)
			if err != nil {
				return nil
			}
			row[i] = val
		}
		grid = append(grid, row)
	}
	return grid
}

func maxProduct(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	res := 0
	d := [][2]int{
		{0, 1},  // right
		{1, 0},  // down
		{1, 1},  // diagonal right
		{1, -1}, // diagonal left
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range d {
				p := calc(grid, i, j, dir, rows, cols)
				if p > res {
					res = p
				}
			}
		}
	}
	return res
}

func calc(grid [][]int, row, col int, dir [2]int, rows, cols int) int {
	r1 := row + 3*dir[0]
	c1 := col + 3*dir[1]
	if r1 >= rows || r1 < 0 || c1 >= cols || c1 < 0 {
		return 0
	}
	res := 1
	for i := 0; i < 4; i++ {
		res *= grid[row+i*dir[0]][col+i*dir[1]]
	}
	return res
}
