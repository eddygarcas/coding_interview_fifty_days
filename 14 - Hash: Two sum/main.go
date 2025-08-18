package main

import "fmt"

func main() {
	// Example input array and target sum
	nums := []int{2, 11, 7, 15}
	result := twoSum(nums, 26)
	fmt.Printf("Result: %d\n", result)
}

// twoSum finds two indices in nums whose values add up to target.
// It returns a slice containing the indices of the two numbers.
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
