package main

// This program finds the element that appears only once in an array where every other element appears twice.
// It uses the XOR bitwise operation to achieve this in linear time and constant space.

import (
	"fmt"
	"time"
)

func main() {
	serie := []int{2, 2, 1, 1, 4}

	// Good approach would be using a map: map[int]int
	// Implement using binary numbers
	start := time.Now()
	singleNumber := findSingleNumber(serie)
	fmt.Printf("The single number is: %d\n", singleNumber)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
}

// findSingleNumber returns the element that appears only once using XOR.
// XOR properties used:
// 1. a ^ a = 0 (XOR of a number with itself is 0)
// 2. a ^ 0 = a (XOR of a number with 0 returns the number)
// 3. a ^ b ^ a = b (XOR is associative and pairs cancel out)
//
// Example with binary representation:
// input = [2, 2, 1, 1, 4]
// 2:    0010
// 2:    0010  -> 0010 ^ 0010 = 0000
// 1:    0001
// 1:    0001  -> 0000 ^ 0001 ^ 0001 = 0000
// 4:    0100  -> 0000 ^ 0100 = 0100 (4)
func findSingleNumber(input []int) int {
	final := 0
	for _, val := range input {
		final ^= val
	}
	return final
}
