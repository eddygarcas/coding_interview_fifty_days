package main

// This program finds the first and last positions of a target value in a sorted array.
// It demonstrates both a linear scan and a binary search approach for efficiency.

import (
	"fmt"
	"time"
)

type point struct {
	x, y int
}

// ToArray converts a point to a slice of two integers.
func (p point) ToArray() []int {
	return []int{p.x, p.y}
}

func main() {
	// e is the sorted input array
	nums := []int{10, 11, 11, 11, 12, 13, 18, 19, 19, 19, 21, 22, 23, 24, 25, 26, 26}
	target := 19
	result := point{x: -1, y: -1}

	// Linear scan to find first and last occurrence
	// Less efficient for large arrays
	start := time.Now()
	for index, element := range nums {
		if element == target {
			if result.x == -1 {
				result = point{x: index, y: index}
			} else {
				result.y = index
			}
		}
	}
	fmt.Printf("First and last position is %v\n", result.ToArray())
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)

	// Binary search approach for better efficiency
	// Finds first and last positions separately
	start = time.Now()
	result.x = findFirst(nums, target)
	result.y = findLast(nums, target)
	fmt.Printf("First and last position is %v\n", result.ToArray())
	elapsed = time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
}

// findFirst returns the index of the first occurrence of target in e using binary search.
// Step-by-step explanation:
//  1. Initialize two pointers: left at the start (0) and right at the end (len(e)-1) of the array.
//  2. While the search range is valid (right >= left):
//     a. Compute mid as the middle index between left and right.
//     b. If e[mid] == target:
//     - Check if mid is the first occurrence:
//     * If mid == 0 (start of array), or the previous element e[mid-1] != target, then mid is the first occurrence, so return mid.
//     * Otherwise, move the search to the left half (right = mid - 1) to find an earlier occurrence.
//     c. If e[mid] > target, move the search to the left half (right = mid - 1).
//     d. If e[mid] < target, move the search to the right half (left = mid + 1).
//  3. If the loop ends without finding the target, return -1.
func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		mid := (right + left) / 2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// findLast returns the index of the last occurrence of target in e using binary search.
func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		mid := (right + left) / 2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
