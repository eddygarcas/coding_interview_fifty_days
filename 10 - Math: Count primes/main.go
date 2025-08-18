package main

// This program implements the Sieve of Eratosthenes algorithm to count prime numbers less than n.
//
// The Sieve of Eratosthenes is an ancient and efficient algorithm for finding prime numbers:
// 1. Create a boolean array of n numbers, initially all marked as prime (true)
// 2. Starting from 2 (the first prime), for each prime number p:
//    - Mark all multiples of p as non-prime (false)
//    - These multiples start from p*p since smaller multiples were already marked
// 3. Continue this process up to sqrt(n), as larger numbers' multiples are already marked
// 4. The remaining true values in the array represent prime numbers
//
// Time Complexity: O(n log log n)
// Space Complexity: O(n)

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	primes := countPrimes(9)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
	fmt.Printf("Total number of primes %d\n", primes)
}

// countPrimes implements the Sieve of Eratosthenes to find all primes less than n.
// Returns the count of prime numbers found.
// For n < 2, returns 0 as there are no prime numbers less than 2.
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	// Initialize slice where isPrime[i] represents if i is prime
	isPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false // 0 and 1 are not prime

	// Mark non-prime numbers using Sieve of Eratosthenes
	// We only need to check up to sqrt(n) because:
	// If a number n is not prime, it can be factored into two factors a and b
	// If both a and b were greater than sqrt(n), then a*b would be greater than n
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if isPrime[i] {
			// Mark all multiples of i as non-prime
			for multiple := i * 2; multiple < n; multiple += i {
				isPrime[multiple] = false
			}
		}
	}

	// Count remaining prime numbers
	primeCount := 0
	for i := 0; i < n; i++ {
		if isPrime[i] {
			primeCount++
		}
	}
	return primeCount
}
