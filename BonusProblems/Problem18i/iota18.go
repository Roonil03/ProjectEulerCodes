// package main

// import (
// 	"fmt"
// 	"math"
// )

// func sieve(limit int) []int {
// 	isPrime := make([]bool, limit+1)
// 	for i := 2; i <= limit; i++ {
// 		isPrime[i] = true
// 	}
// 	for i := 2; i*i <= limit; i++ {
// 		if isPrime[i] {
// 			for j := i * i; j <= limit; j += i {
// 				isPrime[j] = false
// 			}
// 		}
// 	}
// 	var primes []int
// 	for i := 2; i <= limit; i++ {
// 		if isPrime[i] {
// 			primes = append(primes, i)
// 		}
// 	}
// 	return primes
// }

// func segmentedSieve(start, end int) []int {
// 	limit := int(math.Sqrt(float64(end))) + 1
// 	smallPrimes := sieve(limit)
// 	segment := make([]bool, end-start)
// 	for i := range segment {
// 		segment[i] = true
// 	}
// 	for _, p := range smallPrimes {
// 		firstMultiple := ((start + p - 1) / p) * p
// 		for multiple := firstMultiple; multiple < end; multiple += p {
// 			segment[multiple-start] = false
// 		}
// 	}
// 	var primes []int
// 	for i := 0; i < len(segment); i++ {
// 		if segment[i] {
// 			primes = append(primes, start+i)
// 		}
// 	}
// 	return primes
// }

// func R(p int) int {
// 	product := 1
// 	for x := 0; x < p; x++ {
// 		value := x*x*x - 3*x + 4
// 		value = ((value % p) + p) % p
// 		if value == 0 {
// 			return 0
// 		}
// 		product = (product * value) % p
// 	}
// 	return product
// }

// func main() {
// 	start := 1000000000
// 	end := 1100000000
// 	primes := segmentedSieve(start, end)
// 	sum := 0
// 	for _, p := range primes {
// 		//fmt.Println(p)
// 		sum += R(p)
// 	}
// 	fmt.Println("The answer is: ", sum)
// }

package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func sieve(limit int) []int {
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func segmentedSieve(start, end int) []int {
	limit := int(math.Sqrt(float64(end))) + 1
	smallPrimes := sieve(limit)
	segment := make([]bool, end-start)
	for i := range segment {
		segment[i] = true
	}
	for _, p := range smallPrimes {
		firstMultiple := ((start + p - 1) / p) * p
		if firstMultiple < 2 {
			firstMultiple = 2 * p
		}
		for multiple := firstMultiple; multiple < end; multiple += p {
			segment[multiple-start] = false
		}
	}
	var primes []int
	startIdx := 0
	if start <= 2 {
		startIdx = 2 - start
	} else if start%2 == 0 {
		startIdx = 1
	}
	for i := startIdx; i < len(segment); i += 2 {
		if segment[i] {
			primes = append(primes, start+i)
		}
	}
	return primes
}

func powMod(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return result
}

func hasRoot(p int) bool {
	for x := 0; x < p && x < 1000; x++ {
		val := ((x*x*x-3*x+4)%p + p) % p
		if val == 0 {
			return true
		}
	}
	if p <= 1000 {
		return false
	}
	for i := 0; i < min(100, p/2); i++ {
		x := i
		val := ((x*x*x-3*x+4)%p + p) % p
		if val == 0 {
			return true
		}
		x = p - 1 - i
		val = ((x*x*x-3*x+4)%p + p) % p
		if val == 0 {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func R(p int) int {
	if hasRoot(p) {
		return 0
	}
	product := 1
	for x := 0; x < p; x++ {
		val := x*x*x - 3*x + 4
		val = ((val % p) + p) % p
		product = (product * val) % p
		if product == 0 {
			return 0
		}
	}
	return product
}

func computeRSum(primes []int, start, end int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := start; i < end; i++ {
		sum += R(primes[i])
	}
	result <- sum
}

func main() {
	start := 1000000000
	end := 1100000000
	fmt.Println("Generating primes...")
	primes := segmentedSieve(start, end)
	fmt.Printf("Found %d primes\n", len(primes))
	numWorkers := runtime.NumCPU()
	chunkSize := len(primes) / numWorkers
	if chunkSize == 0 {
		chunkSize = 1
	}
	result := make(chan int, numWorkers)
	var wg sync.WaitGroup
	fmt.Println("Computing R(p) values...")
	for i := 0; i < len(primes); i += chunkSize {
		endIdx := i + chunkSize
		if endIdx > len(primes) {
			endIdx = len(primes)
		}

		wg.Add(1)
		go computeRSum(primes, i, endIdx, result, &wg)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	totalSum := 0
	for sum := range result {
		totalSum += sum
	}
	fmt.Println("The answer is:", totalSum)
}
