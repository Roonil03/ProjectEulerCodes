package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filepath := "Problem0098. AnagramicSquares/input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Scan()
	line := sc.Text()
	raw := strings.Split(line, ",")
	ags := make(map[string][]string)
	for _, w := range raw {
		w = strings.Trim(w, `"`)
		sig := sortString(w)
		ags[sig] = append(ags[sig], w)
	}
	var groups [][]string
	for _, grp := range ags {
		if len(grp) > 1 {
			groups = append(groups, grp)
		}
	}
	maxLen := 0
	for _, grp := range groups {
		if l := len(grp[0]); l > maxLen {
			maxLen = l
		}
	}
	sqrlen := make(map[int][]string)
	squareSet := make(map[string]map[string]bool)
	for L := 1; L <= maxLen; L++ {
		lo := int64(1)
		for i := 1; i < L; i++ {
			lo *= 10
		}
		hi := lo*10 - 1
		start := int64(0)
		for i := 1; int64(i)*int64(i) < lo; i++ {
			start = int64(i)
		}
		sqs := []string{}
		set := map[string]bool{}
		for n := start; ; n++ {
			sq := n * n
			if sq > hi {
				break
			}
			s := strconv.FormatInt(sq, 10)
			if len(s) == L {
				sqs = append(sqs, s)
				set[s] = true
			}
		}
		sqrlen[L] = sqs
		squareSet[strconv.Itoa(L)] = set
	}
	res := 0
	for _, grp := range groups {
		L := len(grp[0])
		for _, word := range grp {
			for _, sq := range sqrlen[L] {
				mapping := make(map[rune]byte)
				used := make(map[byte]bool)
				valid := true
				for i, r := range word {
					d := sq[i]
					if m, ok := mapping[r]; ok {
						if m != d {
							valid = false
							break
						}
					} else {
						if used[d] {
							valid = false
							break
						}
						if i == 0 && d == '0' {
							valid = false
							break
						}
						mapping[r] = d
						used[d] = true
					}
				}
				if !valid {
					continue
				}
				for _, other := range grp {
					if other == word {
						continue
					}
					sb := make([]byte, L)
					for i, r := range other {
						sb[i] = mapping[r]
					}
					s2 := string(sb)
					if !squareSet[strconv.Itoa(L)][s2] {
						valid = false
						break
					}
					if v, _ := strconv.Atoi(s2); v > res {
						res = v
					}
				}
			}
		}
	}
	fmt.Println("The answer is:", res)
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}
