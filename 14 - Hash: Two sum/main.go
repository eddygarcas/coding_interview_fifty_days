package main

// This program solves the Two Sum problem using a hash map approach.
// Given an array of integers and a target sum, it finds two numbers that add up to the target.
// The solution uses a single pass through the array, storing complements in a hash map.
// Time Complexity: O(n) where n is length of input array
// Space Complexity: O(n) for storing the hash map

import "fmt"

func main() {
	// Example input array and target sum
	nums := []int{2, 11, 7, 15}
	result := twoSum(nums, 26)
	fmt.Printf("Result: %d\n", result)
}

// twoSum finds two indices in nums whose values add up to target.
// It uses a hash map to store complements (target - current number) as we iterate.
// For each number, we check if its complement exists in the map.
// If found, we've found two numbers that sum to target.
// Returns the indices of the two numbers that sum to target, or empty slice if not found.
func twoSum(nums []int, target int) []int {
	// Map to store the value needed to reach the target and its corresponding index
	findMap := make(map[int]int, len(nums))

	for index, num := range nums {
		// Check if current number is a complement of a previously seen number
		if val, ok := findMap[num]; ok {
			// If found, append current index and the stored index to result
			return []int{val, index}
		}
		// Store the complement (target - num) and the current index
		findMap[target-num] = index

	}
	return []int{}
}
