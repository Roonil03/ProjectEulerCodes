package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := "Problem0102. TriangleContainment/input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	count := 0
	for sc.Scan() {
		var ax, ay, bx, by, cx, cy int
		fmt.Sscanf(sc.Text(), "%d,%d,%d,%d,%d,%d", &ax, &ay, &bx, &by, &cx, &cy)
		s1 := (ax*by - ay*bx) > 0
		s2 := (bx*cy - by*cx) > 0
		s3 := (cx*ay - cy*ax) > 0
		if s1 == s2 && s2 == s3 {
			count++
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("The answer is:", count)
}
