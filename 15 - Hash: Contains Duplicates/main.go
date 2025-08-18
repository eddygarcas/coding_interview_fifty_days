package main

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
