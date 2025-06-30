package main

import (
	"fmt"
	"time"
)

func main() {
	// Example input array
	nums := []int{2, 1, 3, 1, 1, 3, 3, 3, 3}
	// Call majorityElement to find the majority element in the array
	start := time.Now()
	result := majorityElement(nums)
	fmt.Printf("result: %v\n", result)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)

	start = time.Now()
	result = majorityBoyer(nums)
	fmt.Printf("result: %v\n", result)
	elapsed = time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
}

// majorityElement finds the element that appears more than n/2 times in the nums slice.
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

// Here will implement the Boyer-Moore majority vote algorithm
// This algorithm only works if you are guaranteed tho have one element that occurs more than n/2 time
// in this case [1,1,2,2,3] this algorithm won't work.
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
