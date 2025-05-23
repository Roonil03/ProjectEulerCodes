package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	filename := "Problem0022. NamesScores/input.txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print("Error in opening file")
	}
	temp := strings.TrimSpace(string(data))
	names := strings.Split(temp, ",")
	for i := range names {
		names[i] = strings.TrimSpace(names[i])
	}
	sort.Strings(names)
	res := int64(0)
	for i, n := range names {
		pos := int64(i + 1)
		s := int64(pos) * value(n)
		//fmt.Println(i, ". ", n, ": ", s)
		res += s
	}
	fmt.Println("The answer is: ", res)
}

func value(name string) int64 {
	val := int64(0)
	for _, ch := range name {
		if ch >= 'A' && ch <= 'Z' {
			val += int64(ch - 'A' + 1)
		}
	}
	return val
}
