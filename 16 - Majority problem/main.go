package main

import (
	"fmt"
	"time"
)

// This program finds the majority element in an array using two approaches:
// 1. A hashmap-based counting method
// 2. The Boyer-Moore voting algorithm
//
// A majority element is defined as an element that appears more than n/2 times
// in the array, where n is the array length.

func main() {
	// Example input array containing a majority element (3)
	nums := []int{2, 1, 3, 1, 1, 3, 3, 3, 3}

	// Test hashmap-based approach
	start := time.Now()
	result := majorityElement(nums)
	fmt.Printf("result: %v\n", result)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)

	// Test Boyer-Moore approach
	start = time.Now()
	result = majorityBoyer(nums)
	fmt.Printf("result: %v\n", result)
	elapsed = time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
}

// majorityElement finds the element that appears more than n/2 times in the nums slice.
// It uses a hashmap to count occurrences of each element.
// Time complexity: O(n), Space complexity: O(n)
// Returns the majority element if it exists, otherwise returns 0.
func majorityElement(nums []int) int {
	// Map to store the count of each element
	findMap := make(map[int]int, len(nums))
	for _, num := range nums {
		// Increment the count for num
		if val, ok := findMap[num]; ok {
			findMap[num] = val + 1
			// If count exceeds n/2, return num as the majority element
			if findMap[num] > len(nums)/2 {
				return num
			}
		} else {
			findMap[num] = 1
		}
	}
	// No majority element found
	return 0
}

// majorityBoyer implements the Boyer-Moore voting algorithm to find the majority element.
// The algorithm works by maintaining a candidate and a count:
// 1. Initialize candidate as first element and count as 1
// 2. For each subsequent element:
//   - If count becomes 0, pick current element as new candidate
//   - If element matches candidate, increment count
//   - If element differs from candidate, decrement count
//
// 3. Final candidate is the majority element
//
// Time complexity: O(n), Space complexity: O(1)
// Note: This algorithm only works when a majority element is guaranteed to exist.
// For arrays like [1,1,2,2,3], it may give incorrect results.
func majorityBoyer(nums []int) int {
	count := 1
	candidate := nums[0]
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
