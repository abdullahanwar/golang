package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"prime"
)

func main() {
	t1 := time.Now()
	args := os.Args[1:]
	m, _ := strconv.Atoi(args[0])
	n, _ := strconv.Atoi(args[1])
	// batch, _ := strconv.Atoi(args[2])
	prime.BuildPrefix()
	// countParallel(m, n, batch)
	fmt.Printf("Total elapsed time: %d \n", (time.Now()).Sub(t1).Milliseconds())
	fmt.Printf("Expected number of primes:	%d", int(float64(n)/math.Log(float64(n))-float64(m)/math.Log(float64(m))))
}
