package main

import (
	"fmt"
	"sort"
	"time"
)

/*
*
Missing Number problem
Given an array containing n distinct numbers taken from 0,1,2,3,...,n.
find the number that is missing from the array.
*/
func main() {
	// input example [5,3,2,4,0]
	// Sort it [0,2,3,4,5] = position: 5 == 5 && position: 0 == 0
	// left and right
	input := []int{5, 3, 2, 4, 0, 6, 1, 8, 9}
	missing := -1

	l := 0          // Left pointer for checking positions
	r := len(input) // Right pointer for checking positions
	start := time.Now()
	sort.Ints(input) // Sort the array for position checking
	// Check from both ends for the missing number
	for l < r {
		if l != input[l] {
			missing = l
			break
		} else if r != input[r-1] {
			missing = r
			break
		} else {
			l++
			r--
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
	fmt.Printf("The missing number is %d\n", missing)

	// Alternative approach using Gauss' formula
	// Gauss' formula allows you to quickly calculate the sum of the first n natural numbers:
	//     sum = n * (n + 1) / 2.
	// Pre-conditions for using Gauss' formula in the missing number problem:
	//   1. The sequence must contain all integers from 0 to n, with exactly one number missing.
	//   2. There should be no duplicates or extra numbers outside the range 0 to n.
	// This can be used to find a missing number in a sequence from 0 to n by comparing the
	// expected sum to the actual sum.
	currentSum := 0
	input = []int{5, 3, 2, 4, 0, 6, 1, 8, 9}
	// Approach using Gauss
	start = time.Now()
	for _, num := range input {
		currentSum += num // Sum all elements in the array
	}
	n := len(input)                // Number of elements in the array
	intendedSum := n * (n + 1) / 2 // Expected sum from 0 to n
	elapsed = time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
	fmt.Printf("The missing number is %d\n", intendedSum-currentSum)
}
