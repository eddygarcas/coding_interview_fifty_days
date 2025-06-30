package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// Even are not prime except 2 as they actually divisible by 2
	// from 2 to input
	// if number is even skip
	// if number is divisible by 2 skip
	// Sieve of Eratosthenes
	start := time.Now()

	primes := countPrimes(9)

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
	fmt.Printf("Total number of primes %d\n", primes)
}

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	isPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if isPrime[i] {
			for m_of_i := i * 2; m_of_i < n; m_of_i += i {
				isPrime[m_of_i] = false
			}
		}
	}
	primeCount := 0
	for i := 0; i < n; i++ {
		if isPrime[i] {
			primeCount++
		}
	}
	return primeCount
}
