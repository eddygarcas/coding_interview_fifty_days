package main

// This program checks if an array contains any duplicate values using a hash map.
// It demonstrates a simple hash-based approach to track and identify duplicates:
// 1. Create a map to store seen numbers
// 2. Iterate through array once, checking if each number exists in map
// 3. Return true immediately when a duplicate is found
// 4. Otherwise mark number as seen and continue
//
// Time Complexity: O(n) - single pass through the array
// Space Complexity: O(n) - hash map may store up to n elements

import "fmt"

func main() {
	// Example input array containing integers
	nums := []int{2, 7, 11, 15, 2}
	// Call duplicates to check if the array contains any duplicate values
	result := duplicates(nums)
	fmt.Printf("Result: %v\n", result)
}

// duplicates checks if there are any duplicate values in the nums slice.
// Returns true if a duplicate is found, otherwise returns false.
func duplicates(nums []int) bool {
	// Map to track seen numbers
	findMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		// If num is already in the map, a duplicate exists
		if _, ok := findMap[num]; ok {
			return ok
		}
		// Mark num as seen
		findMap[num] = true
	}
	// No duplicates found
	return false
}
