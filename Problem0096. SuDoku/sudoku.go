package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
)

const (
	N       = 9
	fullSet = (1 << 10) - 2
)

var (
	grid [N][N]int
	rm   [N]int
	cm   [N]int
	bm   [N]int
)

func solveBacktrack() bool {
	minCount, selR, selC := 10, -1, -1
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if grid[r][c] == 0 {
				m := fullSet &^ (rm[r] | cm[c] | bm[(r/3)*3+(c/3)])
				cnt := bits.OnesCount(uint(m))
				if cnt == 0 {
					return false
				}
				if cnt < minCount {
					minCount, selR, selC = cnt, r, c
					if cnt == 1 {
						break
					}
				}
			}
		}
	}
	if selR < 0 {
		return true
	}
	mask := fullSet &^ (rm[selR] | cm[selC] | bm[(selR/3)*3+(selC/3)])
	for d := 1; d <= 9; d++ {
		bit := 1 << d
		if mask&bit != 0 {
			grid[selR][selC] = d
			rm[selR] |= bit
			cm[selC] |= bit
			bm[(selR/3)*3+(selC/3)] |= bit
			if solveBacktrack() {
				return true
			}
			grid[selR][selC] = 0
			rm[selR] &^= bit
			cm[selC] &^= bit
			bm[(selR/3)*3+(selC/3)] &^= bit
		}
	}
	return false
}

func main() {
	filepath := "Problem0096. SuDoku/input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	res := 0
	for puzzle := 0; puzzle < 50; puzzle++ {
		sc.Scan()
		for i := range rm {
			rm[i], cm[i], bm[i] = 0, 0, 0
		}
		for r := 0; r < N; r++ {
			sc.Scan()
			line := sc.Text()
			for c := 0; c < N; c++ {
				d := int(line[c] - '0')
				grid[r][c] = d
				if d != 0 {
					bit := 1 << d
					rm[r] |= bit
					cm[c] |= bit
					bm[(r/3)*3+(c/3)] |= bit
				}
			}
		}
		if !solveBacktrack() {
			log.Fatalf("Puzzle %d has no solution", puzzle+1)
		}
		res += 100*grid[0][0] + 10*grid[0][1] + grid[0][2]
	}
	fmt.Println("The answer is:", res)
}
