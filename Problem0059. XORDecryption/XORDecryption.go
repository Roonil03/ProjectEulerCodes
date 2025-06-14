package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readCipher(name string) []byte {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var data []byte
	sc := bufio.NewScanner(f)
	sc.Scan()
	for _, tok := range strings.Split(sc.Text(), ",") {
		n, _ := strconv.Atoi(tok)
		data = append(data, byte(n))
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func score(plain []byte) int {
	s := 0
	for _, b := range plain {
		if b == ' ' {
			s++
		}
	}
	return s
}

func main() {
	filename := "Problem0059. XORDecryption/input.txt"
	cipher := readCipher(filename)
	bestScore := -1
	var key [3]byte
	plain := make([]byte, len(cipher))
	for k0 := byte('a'); k0 <= 'z'; k0++ {
		for k1 := byte('a'); k1 <= 'z'; k1++ {
			for k2 := byte('a'); k2 <= 'z'; k2++ {
				ok := true
				for i, c := range cipher {
					p := c ^ []byte{k0, k1, k2}[i%3]
					if p < 32 || p > 126 {
						ok = false
						break
					}
					plain[i] = p
				}
				if !ok {
					continue
				}
				if sc := score(plain); sc > bestScore {
					bestScore = sc
					key = [3]byte{k0, k1, k2}
				}
			}
		}
	}
	sum := 0
	for i, c := range cipher {
		p := c ^ key[i%3]
		sum += int(p)
	}
	fmt.Println("The answer is: ", sum)
}
