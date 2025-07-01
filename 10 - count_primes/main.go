package main

// This program counts the number of prime numbers less than a given integer n using the Sieve of Eratosthenes algorithm.

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// Even numbers are not prime except 2, as they are divisible by 2
	// Sieve of Eratosthenes efficiently marks non-prime numbers
	start := time.Now()

	primes := countPrimes(9)

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
	fmt.Printf("Total number of primes %d\n", primes)
}

// countPrimes returns the number of prime numbers less than n using the Sieve of Eratosthenes.
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	// Initialize a boolean slice where isPrime[i] indicates if i is prime
	isPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false

	// Mark non-prime numbers
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if isPrime[i] {
			for m_of_i := i * 2; m_of_i < n; m_of_i += i {
				isPrime[m_of_i] = false
			}
		}
	}
	primeCount := 0
	// Count primes
	for i := 0; i < n; i++ {
		if isPrime[i] {
			primeCount++
		}
	}
	return primeCount
}
