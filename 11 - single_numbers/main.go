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
func findSingleNumber(input []int) int {
	final := 0
	for _, val := range input {
		// This operator will work at binary level so for instance:
		// 2 = 0010 if val was 2 which is 0010 the operation will return 0 = 0000
		final ^= val
	}
	return final
}
