package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Roonil03/Templates_GoLang/priorityqueue"
)

func main() {
	filepath := /*"Problem0079. PasscodeDerivation/*/ "input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	mcb := make(map[byte]map[byte]bool)
	ad := make(map[byte]bool)
	deg := make(map[byte]int)
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		if len(line) != 3 {
			continue
		}
		for i := 0; i < 3; i++ {
			dig := line[i]
			ad[dig] = true
			if mcb[dig] == nil {
				mcb[dig] = make(map[byte]bool)
			}
			if _, exists := deg[dig]; !exists {
				deg[dig] = 0
			}
		}
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 3; j++ {
				first := line[i]
				second := line[j]
				if !mcb[second][first] {
					mcb[second][first] = true
					deg[second]++
				}
			}
		}
	}
	pq := priorityqueue.NewMinHeap()
	q := make(map[byte]bool)
	for dig := range ad {
		if deg[dig] == 0 {
			pq.Push(int(dig))
			q[dig] = true
		}
	}
	var res []byte
	for pq.Len() > 0 {
		a := pq.Pop().(int)
		ch := byte(a)
		delete(q, ch)
		res = append(res, ch)
		for dig := range ad {
			if mcb[dig][ch] {
				deg[dig]--
				if deg[dig] == 0 && !q[dig] {
					pq.Push(int(dig))
					q[dig] = true
				}
			}
		}
	}
	if len(res) != len(ad) {
		fmt.Println("SMTH WRONG")
		return
	}
	fmt.Println("The answer is: ", string(res))
}
