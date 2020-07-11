package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	// excluding negative number
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for f := 3; f*f <= n; f += 2 {
		if n%f == 0 {
			return false
		}
	}
	return true
}

// π(𝑁)≃ N/Log(N) - M/Log(M)
func countPrimes(m int, n int, cout chan<- int) {
	count := 0
	for i := m; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	cout <- count
	close(cout)
}

func countParallel(start, end, batchSize int) {
	count := make(chan int)
	for i := start; i < end; i += batchSize {
		j := i + batchSize
		if j < end {
			j = end
		}
		go countPrimes(i, j, count)
	}
	total := 0
	for c := range count {
		fmt.Println("Got: ", c)
		total += c
	}
	fmt.Printf("Total number of primes: %d \n", total)
}

func main() {
	t1 := time.Now()
	fmt.Println("Hello !!")
	countParallel(100, 10000000, 100000)
	fmt.Printf("Total elapsed time: ", (time.Now()).Sub(t1))
}
