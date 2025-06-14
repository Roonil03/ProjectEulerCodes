package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var val = map[byte]int{
	'2': 2, '3': 3, '4': 4, '5': 5,
	'6': 6, '7': 7, '8': 8, '9': 9,
	'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}

var suitIdx = map[byte]int{'C': 0, 'D': 1, 'H': 2, 'S': 3}

type handScore struct {
	rank int
	kick [5]int
}

func better(a, b handScore) bool {
	if a.rank != b.rank {
		return a.rank > b.rank
	}
	for i := 0; i < 5; i++ {
		if a.kick[i] != b.kick[i] {
			return a.kick[i] > b.kick[i]
		}
	}
	return false
}

func score(h [5][2]int) handScore {
	var cnt [15]int
	var suits [4]int
	minV, maxV := 14, 2
	for _, c := range h {
		v, s := c[0], c[1]
		cnt[v]++
		suits[s]++
		if v < minV {
			minV = v
		}
		if v > maxV {
			maxV = v
		}
	}
	isFlush := suits[0] == 5 || suits[1] == 5 || suits[2] == 5 || suits[3] == 5
	isStraight := false
	for hi := maxV; hi >= 5; hi-- {
		ok := true
		for d := 0; d < 5; d++ {
			if cnt[hi-d] == 0 {
				ok = false
				break
			}
		}
		if ok {
			isStraight = true
			maxV = hi
			break
		}
	}
	if !isStraight &&
		cnt[14] == 1 && cnt[5] == 1 && cnt[4] == 1 && cnt[3] == 1 && cnt[2] == 1 {
		isStraight, maxV = true, 5
	}
	var four, three, pairs []int
	for v := 14; v >= 2; v-- {
		switch cnt[v] {
		case 4:
			four = append(four, v)
		case 3:
			three = append(three, v)
		case 2:
			pairs = append(pairs, v)
		}
	}
	var sc handScore
	switch {
	case isStraight && isFlush && maxV == 14:
		sc.rank, sc.kick[0] = 9, 14
	case isStraight && isFlush:
		sc.rank, sc.kick[0] = 8, maxV
	case len(four) == 1:
		sc.rank, sc.kick[0] = 7, four[0]
		fillSingles(&sc, cnt, map[int]bool{four[0]: true})
	case len(three) == 1 && len(pairs) == 1:
		sc.rank, sc.kick[0], sc.kick[1] = 6, three[0], pairs[0]
	case isFlush:
		sc.rank = 5
		fillSingles(&sc, cnt, nil)
	case isStraight:
		sc.rank, sc.kick[0] = 4, maxV
	case len(three) == 1:
		sc.rank, sc.kick[0] = 3, three[0]
		fillSingles(&sc, cnt, map[int]bool{three[0]: true})
	case len(pairs) == 2:
		sc.rank, sc.kick[0], sc.kick[1] = 2, pairs[0], pairs[1]
		fillSingles(&sc, cnt, map[int]bool{pairs[0]: true, pairs[1]: true})
	case len(pairs) == 1:
		sc.rank, sc.kick[0] = 1, pairs[0]
		fillSingles(&sc, cnt, map[int]bool{pairs[0]: true})
	default:
		fillSingles(&sc, cnt, nil)
	}
	return sc
}

func fillSingles(sc *handScore, cnt [15]int, skip map[int]bool) {
	i := 0
	for v := 14; v >= 2 && i < len(sc.kick); v-- {
		if cnt[v] == 1 && (skip == nil || !skip[v]) {
			for i < len(sc.kick) && sc.kick[i] != 0 {
				i++
			}
			if i < len(sc.kick) {
				sc.kick[i] = v
			}
		}
	}
}

func main() {
	filename := "Problem0054. PokerHands/input.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	p1wins := 0
	for sc.Scan() {
		line := sc.Text()
		var cards [10][2]int
		id := 0
		for i := 0; i < len(line); i += 3 {
			v := val[line[i]]
			s := suitIdx[line[i+1]]
			cards[id] = [2]int{v, s}
			id++
		}
		var h1, h2 [5][2]int
		copy(h1[:], cards[0:5])
		copy(h2[:], cards[5:10])
		if better(score(h1), score(h2)) {
			p1wins++
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("The answer is", p1wins)
}
