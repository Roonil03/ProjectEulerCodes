package main

// // import (
// // 	"fmt"
// // 	"math"
// // )

// // func sieve(limit int) []int {
// // 	isPrime := make([]bool, limit+1)
// // 	for i := 2; i <= limit; i++ {
// // 		isPrime[i] = true
// // 	}
// // 	for i := 2; i*i <= limit; i++ {
// // 		if isPrime[i] {
// // 			for j := i * i; j <= limit; j += i {
// // 				isPrime[j] = false
// // 			}
// // 		}
// // 	}
// // 	var primes []int
// // 	for i := 2; i <= limit; i++ {
// // 		if isPrime[i] {
// // 			primes = append(primes, i)
// // 		}
// // 	}
// // 	return primes
// // }

// // func segmentedSieve(start, end int) []int {
// // 	limit := int(math.Sqrt(float64(end))) + 1
// // 	smallPrimes := sieve(limit)
// // 	segment := make([]bool, end-start)
// // 	for i := range segment {
// // 		segment[i] = true
// // 	}
// // 	for _, p := range smallPrimes {
// // 		firstMultiple := ((start + p - 1) / p) * p
// // 		for multiple := firstMultiple; multiple < end; multiple += p {
// // 			segment[multiple-start] = false
// // 		}
// // 	}
// // 	var primes []int
// // 	for i := 0; i < len(segment); i++ {
// // 		if segment[i] {
// // 			primes = append(primes, start+i)
// // 		}
// // 	}
// // 	return primes
// // }

// // func R(p int) int {
// // 	product := 1
// // 	for x := 0; x < p; x++ {
// // 		value := x*x*x - 3*x + 4
// // 		value = ((value % p) + p) % p
// // 		if value == 0 {
// // 			return 0
// // 		}
// // 		product = (product * value) % p
// // 	}
// // 	return product
// // }

// // func main() {
// // 	start := 1000000000
// // 	end := 1100000000
// // 	primes := segmentedSieve(start, end)
// // 	sum := 0
// // 	for _, p := range primes {
// // 		//fmt.Println(p)
// // 		sum += R(p)
// // 	}
// // 	fmt.Println("The answer is: ", sum)
// // }

// package main

// // import (
// // 	"fmt"
// // 	"math"
// // 	// "runtime"
// // 	// "sync"
// // )

// // func sieve(limit int) []int {
// // 	isPrime := make([]bool, limit+1)
// // 	for i := 2; i <= limit; i++ {
// // 		isPrime[i] = true
// // 	}
// // 	for i := 2; i*i <= limit; i++ {
// // 		if isPrime[i] {
// // 			for j := i * i; j <= limit; j += i {
// // 				isPrime[j] = false
// // 			}
// // 		}
// // 	}
// // 	var primes []int
// // 	for i := 2; i <= limit; i++ {
// // 		if isPrime[i] {
// // 			primes = append(primes, i)
// // 		}
// // 	}
// // 	return primes
// // }

// // func segmentedSieve(start, end int) []int {
// // 	limit := int(math.Sqrt(float64(end))) + 1
// // 	smallPrimes := sieve(limit)
// // 	segment := make([]bool, end-start)
// // 	for i := range segment {
// // 		segment[i] = true
// // 	}
// // 	for _, p := range smallPrimes {
// // 		firstMultiple := ((start + p - 1) / p) * p
// // 		if firstMultiple < 2 {
// // 			firstMultiple = 2 * p
// // 		}
// // 		for multiple := firstMultiple; multiple < end; multiple += p {
// // 			segment[multiple-start] = false
// // 		}
// // 	}
// // 	var primes []int
// // 	startIdx := 0
// // 	if start <= 2 {
// // 		startIdx = 2 - start
// // 	} else if start%2 == 0 {
// // 		startIdx = 1
// // 	}
// // 	for i := startIdx; i < len(segment); i += 2 {
// // 		if segment[i] {
// // 			primes = append(primes, start+i)
// // 		}
// // 	}
// // 	return primes
// // }

// // func powMod(base, exp, mod int) int {
// // 	result := 1
// // 	base %= mod
// // 	for exp > 0 {
// // 		if exp&1 == 1 {
// // 			result = (result * base) % mod
// // 		}
// // 		base = (base * base) % mod
// // 		exp >>= 1
// // 	}
// // 	return result
// // }

// // func hasRoot(p int) bool {
// // 	for x := 0; x < p && x < 1000; x++ {
// // 		val := ((x*x*x-3*x+4)%p + p) % p
// // 		if val == 0 {
// // 			return true
// // 		}
// // 	}
// // 	if p <= 1000 {
// // 		return false
// // 	}
// // 	for i := 0; i < min(100, p/2); i++ {
// // 		x := i
// // 		val := ((x*x*x-3*x+4)%p + p) % p
// // 		if val == 0 {
// // 			return true
// // 		}
// // 		x = p - 1 - i
// // 		val = ((x*x*x-3*x+4)%p + p) % p
// // 		if val == 0 {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// // func min(a, b int) int {
// // 	if a < b {
// // 		return a
// // 	}
// // 	return b
// // }

// // func R(p int) int {
// // 	if hasRoot(p) {
// // 		return 0
// // 	}
// // 	product := 1
// // 	for x := 0; x < p; x++ {
// // 		val := x*x*x - 3*x + 4
// // 		val = ((val % p) + p) % p
// // 		product = (product * val) % p
// // 		if product == 0 {
// // 			return 0
// // 		}
// // 	}
// // 	return product
// // }

// // func computeRSum(primes []int, start, end int, result chan int, wg *sync.WaitGroup) {
// // 	defer wg.Done()
// // 	sum := 0
// // 	for i := start; i < end; i++ {
// // 		sum += R(primes[i])
// // 	}
// // 	result <- sum
// // }

// // func main() {
// // 	start := 1000000000
// // 	end := 1100000000
// // 	fmt.Println("Generating primes...")
// // 	primes := segmentedSieve(start, end)
// // 	fmt.Printf("Found %d primes\n", len(primes))
// // 	numWorkers := runtime.NumCPU()
// // 	chunkSize := len(primes) / numWorkers
// // 	if chunkSize == 0 {
// // 		chunkSize = 1
// // 	}
// // 	result := make(chan int, numWorkers)
// // 	var wg sync.WaitGroup
// // 	fmt.Println("Computing R(p) values...")
// // 	for i := 0; i < len(primes); i += chunkSize {
// // 		endIdx := i + chunkSize
// // 		if endIdx > len(primes) {
// // 			endIdx = len(primes)
// // 		}

// // 		wg.Add(1)
// // 		go computeRSum(primes, i, endIdx, result, &wg)
// // 	}
// // 	go func() {
// // 		wg.Wait()
// // 		close(result)
// // 	}()
// // 	totalSum := 0
// // 	for sum := range result {
// // 		totalSum += sum
// // 	}
// // 	fmt.Println("The answer is:", totalSum)
// // }

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func segmentedSieve(start, end int64) []int64 {
	fmt.Printf("Starting segmented sieve for range [%d, %d]...\n", start, end)
	startTime := time.Now()
	limit := int64(math.Sqrt(float64(end))) + 1
	fmt.Printf("Finding base primes up to %d...\n", limit)
	isPrime := make([]bool, limit+1)
	for i := int64(2); i <= limit; i++ {
		isPrime[i] = true
	}
	for i := int64(2); i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}
	var basePrimes []int64
	for i := int64(2); i <= limit; i++ {
		if isPrime[i] {
			basePrimes = append(basePrimes, i)
		}
	}
	fmt.Printf("Found %d base primes\n", len(basePrimes))
	segmentSize := int64(1000000)
	var primes []int64
	for segStart := start; segStart <= end; segStart += segmentSize {
		segEnd := segStart + segmentSize - 1
		if segEnd > end {
			segEnd = end
		}
		segmentLength := segEnd - segStart + 1
		segment := make([]bool, segmentLength)
		for i := range segment {
			segment[i] = true
		}
		for _, p := range basePrimes {
			start_multiple := ((segStart + p - 1) / p) * p
			if start_multiple < p*p {
				start_multiple = p * p
			}
			for multiple := start_multiple; multiple <= segEnd; multiple += p {
				segment[multiple-segStart] = false
			}
		}
		for i := int64(0); i < segmentLength; i++ {
			candidate := segStart + i
			if candidate >= 2 && segment[i] {
				primes = append(primes, candidate)
			}
		}
		if len(primes)%10000 == 0 {
			fmt.Printf("Progress: found %d primes so far...\n", len(primes))
		}
	}
	elapsed := time.Since(startTime)
	fmt.Printf("Segmented sieve completed in %v\n", elapsed)
	fmt.Printf("Found %d primes in range [%d, %d]\n", len(primes), start, end)

	return primes
}

func modPow(base, exp, mod int64) int64 {
	result := int64(1)
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		exp = exp >> 1
		base = (base * base) % mod
	}
	return result
}

func computeR(p int64) int64 {
	product := int64(1)
	for x := int64(0); x < p; x++ {
		x_mod := x % p
		x2_mod := (x_mod * x_mod) % p
		x3_mod := (x2_mod * x_mod) % p
		term := (x3_mod - 3*x_mod + 4 + p*p) % p
		product = (product * term) % p
		if product == 0 {
			return 0
		}
	}

	return product
}

func computeWorker(primes []int64, start, end int, results chan<- int64,
	progress chan<- int, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	localSum := int64(0)
	processed := 0
	fmt.Printf("Worker %d: starting with %d primes (indices %d to %d)\n",
		workerID, end-start, start, end-1)
	for i := start; i < end && i < len(primes); i++ {
		p := primes[i]
		r := computeR(p)
		localSum += r
		processed++
		if processed%100 == 0 {
			select {
			case progress <- processed:
			default:
			}
		}
	}
	fmt.Printf("Worker %d completed: processed %d primes, sum = %d\n",
		workerID, processed, localSum)
	results <- localSum
}

func progressMonitor(progress chan int, totalPrimes int, done chan bool) {
	totalProcessed := 0
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case count := <-progress:
			totalProcessed += count
			percentage := float64(totalProcessed) / float64(totalPrimes) * 100
			fmt.Printf("Progress: %d/%d primes processed (%.2f%%)\n",
				totalProcessed, totalPrimes, percentage)

		case <-ticker.C:
			percentage := float64(totalProcessed) / float64(totalPrimes) * 100
			fmt.Printf("Status: %d/%d primes processed (%.2f%%)\n",
				totalProcessed, totalPrimes, percentage)

		case <-done:
			return
		}
	}
}

func main() {
	fmt.Println("=== Large Scale Prime Computation with R(p) ===")
	overallStart := time.Now()
	const (
		RANGE_START = 1000000000
		RANGE_END   = 1100000000
		NUM_THREADS = 1
	)
	fmt.Printf("Range: [%d, %d]\n", RANGE_START, RANGE_END)
	fmt.Printf("Threads: %d\n", NUM_THREADS)
	primes := segmentedSieve(RANGE_START, RANGE_END)
	if len(primes) == 0 {
		fmt.Println("ERROR: No primes found in the specified range!")
		return
	}
	fmt.Printf("\n=== PRIME GENERATION COMPLETE ===\n")
	fmt.Printf("Total primes found: %d\n", len(primes))
	fmt.Println("Close Answer:", len(primes)*4)
	fmt.Printf("First 10 primes: %v\n", primes[:min(10, len(primes))])
	if len(primes) > 10 {
		fmt.Printf("Last 10 primes: %v\n", primes[len(primes)-10:])
	}
	runtime.GOMAXPROCS(NUM_THREADS)
	primesPerThread := len(primes) / NUM_THREADS
	remainder := len(primes) % NUM_THREADS
	fmt.Printf("\n=== STARTING R(p) COMPUTATION ===\n")
	fmt.Printf("Primes per thread: %d (remainder: %d)\n", primesPerThread, remainder)
	results := make(chan int64, NUM_THREADS)
	progress := make(chan int, 100)
	done := make(chan bool)
	var wg sync.WaitGroup
	go progressMonitor(progress, len(primes), done)
	computationStart := time.Now()
	for i := 0; i < NUM_THREADS; i++ {
		startIdx := i * primesPerThread
		endIdx := (i + 1) * primesPerThread

		if i < remainder {
			startIdx += i
			endIdx += i + 1
		} else {
			startIdx += remainder
			endIdx += remainder
		}
		if startIdx < len(primes) {
			wg.Add(1)
			go computeWorker(primes, startIdx, endIdx, results, progress, &wg, i+1)
		}
	}
	go func() {
		wg.Wait()
		close(results)
		done <- true
	}()
	totalSum := int64(0)
	threadsCompleted := 0
	for sum := range results {
		threadsCompleted++
		totalSum += sum
		fmt.Printf("Thread %d result received: %d\n", threadsCompleted, sum)
	}
	computationTime := time.Since(computationStart)
	totalTime := time.Since(overallStart)
	fmt.Printf("FINAL RESULTS\n")
	fmt.Printf("Range processed: [%d, %d]\n", RANGE_START, RANGE_END)
	fmt.Printf("Total primes: %d\n", len(primes))
	fmt.Printf("Threads used: %d\n", NUM_THREADS)
	fmt.Printf("Sum of R(p): %d\n", totalSum)
	fmt.Printf("Sieve + computation time: %v\n", totalTime)
	fmt.Printf("R(p) computation time only: %v\n", computationTime)
	fmt.Printf("Average time per prime: %v\n", time.Duration(int64(computationTime)/int64(len(primes))))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// // #cgo CFLAGS: -I"C:/Program Files/NVIDIA GPU Computing Toolkit/CUDA/v12.0/include"
// // #cgo LDFLAGS: -L"C:/Program Files/NVIDIA GPU Computing Toolkit/CUDA/v12.0/lib/x64" -lcuda -lcudart
// // #include <stdio.h>
// // #include <stdlib.h>
// // #include <cuda_runtime.h>
// // #include <cuda.h>
// //
// // // CUDA kernel for computing R(p) - runs on GPU
// // __global__ void computeR_kernel(long long *primes, long long *results, int start, int count) {
// //     int idx = blockIdx.x * blockDim.x + threadIdx.x;
// //     if (idx >= count) return;
// //
// //     long long p = primes[start + idx];
// //     long long product = 1;
// //
// //     // Compute product of (x^3 - 3x + 4) for x = 0 to p-1
// //     for (long long x = 0; x < p; x++) {
// //         long long x_mod = x % p;
// //         long long x2_mod = (x_mod * x_mod) % p;
// //         long long x3_mod = (x2_mod * x_mod) % p;
// //
// //         // Compute (x^3 - 3x + 4) mod p
// //         long long term = (x3_mod - 3*x_mod + 4 + p*p) % p;
// //         product = (product * term) % p;
// //
// //         // Early termination if product becomes 0
// //         if (product == 0) {
// //             break;
// //         }
// //     }
// //
// //     results[idx] = product;
// // }
// //
// // // Wrapper function to launch CUDA kernel
// // int launch_cuda_computation(long long *h_primes, long long *h_results, int total_primes, int chunk_size) {
// //     long long *d_primes, *d_results;
// //     size_t primes_size = total_primes * sizeof(long long);
// //     size_t results_size = chunk_size * sizeof(long long);
// //
// //     // Allocate GPU memory
// //     if (cudaMalloc(&d_primes, primes_size) != cudaSuccess) return -1;
// //     if (cudaMalloc(&d_results, results_size) != cudaSuccess) return -1;
// //
// //     // Copy primes to GPU
// //     if (cudaMemcpy(d_primes, h_primes, primes_size, cudaMemcpyHostToDevice) != cudaSuccess) return -1;
// //
// //     // Process primes in chunks to manage GPU memory
// //     for (int start = 0; start < total_primes; start += chunk_size) {
// //         int current_chunk = (start + chunk_size > total_primes) ? (total_primes - start) : chunk_size;
// //
// //         // Configure kernel launch parameters
// //         int blockSize = 256;
// //         int gridSize = (current_chunk + blockSize - 1) / blockSize;
// //
// //         // Launch kernel
// //         computeR_kernel<<<gridSize, blockSize>>>(d_primes, d_results, start, current_chunk);
// //
// //         // Wait for kernel to complete
// //         if (cudaDeviceSynchronize() != cudaSuccess) return -1;
// //
// //         // Copy results back to host
// //         if (cudaMemcpy(&h_results[start], d_results, current_chunk * sizeof(long long), cudaMemcpyDeviceToHost) != cudaSuccess) return -1;
// //     }
// //
// //     // Cleanup
// //     cudaFree(d_primes);
// //     cudaFree(d_results);
// //
// //     return 0;
// // }
// //
// // // GPU device information
// // int get_gpu_info() {
// //     int deviceCount;
// //     if (cudaGetDeviceCount(&deviceCount) != cudaSuccess) return -1;
// //
// //     printf("Found %d CUDA devices:\n", deviceCount);
// //     for (int i = 0; i < deviceCount; i++) {
// //         cudaDeviceProp prop;
// //         cudaGetDeviceProperties(&prop, i);
// //         printf("Device %d: %s\n", i, prop.name);
// //         printf("  Compute Capability: %d.%d\n", prop.major, prop.minor);
// //         printf("  Global Memory: %lu MB\n", prop.totalGlobalMem / (1024*1024));
// //         printf("  Multiprocessors: %d\n", prop.multiProcessorCount);
// //         printf("  Max threads per block: %d\n", prop.maxThreadsPerBlock);
// //     }
// //     return deviceCount;
// // }
// import "C"
// import (
//     "fmt"
//     "math"
//     "runtime"
//     "sync"
//     "time"
//     "unsafe"
// )

// func segmentedSieveGPU(start, end int64) []int64 {
//     fmt.Printf("Starting GPU-accelerated segmented sieve for range [%d, %d]...\n", start, end)
//     startTime := time.Now()

//     // Check GPU availability
//     deviceCount := int(C.get_gpu_info())
//     if deviceCount <= 0 {
//         fmt.Println("No CUDA devices found, falling back to CPU sieve")
//         return segmentedSieveCPU(start, end)
//     }

//     // Find base primes using CPU (up to sqrt(end))
//     limit := int64(math.Sqrt(float64(end))) + 1
//     fmt.Printf("Finding base primes up to %d using CPU...\n", limit)

//     basePrimes := simpleSieve(limit)
//     fmt.Printf("Found %d base primes\n", len(basePrimes))

//     // Use CPU sieve for now - GPU sieve implementation would be more complex
//     // and the main GPU acceleration benefit is in R(p) computation
//     primes := segmentedSieveCPU(start, end)

//     elapsed := time.Since(startTime)
//     fmt.Printf("Sieve completed in %v, found %d primes\n", elapsed, len(primes))
//     return primes
// }

// // CPU fallback for segmented sieve
// func segmentedSieveCPU(start, end int64) []int64 {
//     limit := int64(math.Sqrt(float64(end))) + 1
//     basePrimes := simpleSieve(limit)

//     segmentSize := int64(1000000)
//     var primes []int64

//     for segStart := start; segStart <= end; segStart += segmentSize {
//         segEnd := segStart + segmentSize - 1
//         if segEnd > end {
//             segEnd = end
//         }

//         segmentLength := segEnd - segStart + 1
//         segment := make([]bool, segmentLength)
//         for i := range segment {
//             segment[i] = true
//         }

//         for _, p := range basePrimes {
//             startMultiple := ((segStart + p - 1) / p) * p
//             if startMultiple < p*p {
//                 startMultiple = p * p
//             }

//             for multiple := startMultiple; multiple <= segEnd; multiple += p {
//                 segment[multiple-segStart] = false
//             }
//         }

//         for i := int64(0); i < segmentLength; i++ {
//             candidate := segStart + i
//             if candidate >= 2 && segment[i] {
//                 primes = append(primes, candidate)
//             }
//         }
//     }

//     return primes
// }

// // Simple sieve for small numbers
// func simpleSieve(n int64) []int64 {
//     isPrime := make([]bool, n+1)
//     for i := int64(2); i <= n; i++ {
//         isPrime[i] = true
//     }

//     for i := int64(2); i*i <= n; i++ {
//         if isPrime[i] {
//             for j := i * i; j <= n; j += i {
//                 isPrime[j] = false
//             }
//         }
//     }

//     var primes []int64
//     for i := int64(2); i <= n; i++ {
//         if isPrime[i] {
//             primes = append(primes, i)
//         }
//     }
//     return primes
// }

// // GPU-accelerated R(p) computation
// func computeRGPU(primes []int64, chunkSize int) (int64, error) {
//     fmt.Printf("Starting GPU-accelerated R(p) computation for %d primes...\n", len(primes))
//     startTime := time.Now()

//     // Convert Go slices to C arrays
//     primesPtr := (*C.longlong)(unsafe.Pointer(&primes[0]))
//     results := make([]int64, len(primes))
//     resultsPtr := (*C.longlong)(unsafe.Pointer(&results[0]))

//     // Launch CUDA computation
//     ret := C.launch_cuda_computation(primesPtr, resultsPtr, C.int(len(primes)), C.int(chunkSize))
//     if ret != 0 {
//         return 0, fmt.Errorf("CUDA computation failed with code %d", ret)
//     }

//     // Sum all results
//     totalSum := int64(0)
//     for _, r := range results {
//         totalSum += r
//     }

//     elapsed := time.Since(startTime)
//     fmt.Printf("GPU computation completed in %v\n", elapsed)
//     return totalSum, nil
// }

// func hybridWorker(primes []int64, start, end int, useGPU bool, results chan<- int64,
//                   progress chan<- int, wg *sync.WaitGroup, workerID int) {
//     defer wg.Done()

//     localSum := int64(0)
//     processed := 0

//     if useGPU && end-start > 1000 { // Use GPU for large chunks
//         fmt.Printf("GPU Worker %d: processing %d primes (indices %d to %d)\n",
//             workerID, end-start, start, end-1)

//         chunk := primes[start:end]
//         gpuSum, err := computeRGPU(chunk, end-start)
//         if err != nil {
//             fmt.Printf("GPU Worker %d failed, falling back to CPU: %v\n", workerID, err)
//             for i := start; i < end && i < len(primes); i++ {
//                 p := primes[i]
//                 r := computeRCPU(p)
//                 localSum += r
//                 processed++
//             }
//         } else {
//             localSum = gpuSum
//             processed = end - start
//         }
//     } else {
//         fmt.Printf("CPU Worker %d: processing %d primes (indices %d to %d)\n",
//             workerID, end-start, start, end-1)

//         for i := start; i < end && i < len(primes); i++ {
//             p := primes[i]
//             r := computeRCPU(p)
//             localSum += r
//             processed++

//             if processed%100 == 0 {
//                 select {
//                 case progress <- processed:
//                 default:
//                 }
//             }
//         }
//     }

//     fmt.Printf("Worker %d completed: processed %d primes, sum = %d\n",
//         workerID, processed, localSum)
//     results <- localSum
// }

// func computeRCPU(p int64) int64 {
//     product := int64(1)
//     for x := int64(0); x < p; x++ {
//         xMod := x % p
//         x2Mod := (xMod * xMod) % p
//         x3Mod := (x2Mod * xMod) % p

//         term := (x3Mod - 3*xMod + 4 + p*p) % p
//         product = (product * term) % p

//         if product == 0 {
//             return 0
//         }
//     }
//     return product
// }

// func progressMonitor(progress chan int, totalPrimes int, done chan bool) {
//     totalProcessed := 0
//     ticker := time.NewTicker(10 * time.Second)
//     defer ticker.Stop()

//     for {
//         select {
//         case count := <-progress:
//             totalProcessed += count
//             percentage := float64(totalProcessed) / float64(totalPrimes) * 100
//             fmt.Printf("Progress: %d/%d primes processed (%.2f%%)\n",
//                 totalProcessed, totalPrimes, percentage)

//         case <-ticker.C:
//             percentage := float64(totalProcessed) / float64(totalPrimes) * 100
//             fmt.Printf("Status: %d/%d primes processed (%.2f%%)\n",
//                 totalProcessed, totalPrimes, percentage)

//         case <-done:
//             return
//         }
//     }
// }

// func main() {
//     fmt.Println("=== GPU-Accelerated Prime Computation with CUDA ===")
//     overallStart := time.Now()
//     const (
//         RANGE_START     = 1000000000
//         RANGE_END       = 1100000000
//         NUM_WORKERS     = 5
//         GPU_CHUNK_SIZE  = 10000
//     )

//     fmt.Printf("Range: [%d, %d]\n", RANGE_START, RANGE_END)
//     fmt.Printf("Workers: %d\n", NUM_WORKERS)

//     primes := segmentedSieveGPU(RANGE_START, RANGE_END)

//     if len(primes) == 0 {
//         fmt.Println("ERROR: No primes found!")
//         return
//     }

//     fmt.Printf("\n=== PRIME GENERATION COMPLETE ===\n")
//     fmt.Printf("Total primes: %d\n", len(primes))
//     fmt.Printf("First 10: %v\n", primes[:min(10, len(primes))])
//     if len(primes) > 10 {
//         fmt.Printf("Last 10: %v\n", primes[len(primes)-10:])
//     }
//     runtime.GOMAXPROCS(NUM_WORKERS)

//     primesPerWorker := len(primes) / NUM_WORKERS
//     remainder := len(primes) % NUM_WORKERS

//     fmt.Printf("\n=== STARTING HYBRID GPU/CPU R(p) COMPUTATION ===\n")
//     fmt.Printf("Primes per worker: %d (remainder: %d)\n", primesPerWorker, remainder)

//     results := make(chan int64, NUM_WORKERS)
//     progress := make(chan int, 100)
//     done := make(chan bool)
//     var wg sync.WaitGroup

//     go progressMonitor(progress, len(primes), done)

//     computationStart := time.Now()
//     for i := 0; i < NUM_WORKERS; i++ {
//         startIdx := i * primesPerWorker
//         endIdx := (i + 1) * primesPerWorker

//         if i < remainder {
//             startIdx += i
//             endIdx += i + 1
//         } else {
//             startIdx += remainder
//             endIdx += remainder
//         }

//         if startIdx < len(primes) {
//             wg.Add(1)
//             useGPU := (endIdx - startIdx) >= 1000
//             go hybridWorker(primes, startIdx, endIdx, useGPU, results, progress, &wg, i+1)
//         }
//     }

//     go func() {
//         wg.Wait()
//         close(results)
//         done <- true
//     }()

//     totalSum := int64(0)
//     workersCompleted := 0

//     for sum := range results {
//         workersCompleted++
//         totalSum += sum
//         fmt.Printf("Worker %d completed with sum: %d\n", workersCompleted, sum)
//     }

//     computationTime := time.Since(computationStart)
//     totalTime := time.Since(overallStart)

//     fmt.Printf("\n" + "="*60 + "\n")
//     fmt.Printf("FINAL RESULTS - GPU ACCELERATED\n")
//     fmt.Printf("="*60 + "\n")
//     fmt.Printf("Range: [%d, %d]\n", RANGE_START, RANGE_END)
//     fmt.Printf("Total primes: %d\n", len(primes))
//     fmt.Printf("Workers: %d (Hybrid GPU/CPU)\n", NUM_WORKERS)
//     fmt.Printf("Sum of R(p): %d\n", totalSum)
//     fmt.Printf("Total time: %v\n", totalTime)
//     fmt.Printf("R(p) computation time: %v\n", computationTime)
//     fmt.Printf("Average per prime: %v\n", time.Duration(int64(computationTime)/int64(len(primes))))
//     fmt.Printf("="*60 + "\n")
// }

// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }
