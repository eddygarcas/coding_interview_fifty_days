package main

// This program solves the "Container With Most Water" problem.
// It finds the maximum area of water that can be contained between two lines
// from the given heights array, using a two-pointer approach.

import "log"

func main() {
	// height represents the heights of the vertical lines
	var height = []int{1, 8, 6, 2, 5, 7}
	l := 0               // Left pointer
	r := len(height) - 1 // Right pointer
	log.Printf("height %v lengt %v", l, r)
	maxArea := 0 // Stores the maximum area found
	for l < r {
		log.Printf("l %v r %v", l, r)
		width := r - l // Distance between the two pointers
		if height[l] < height[r] {
			// Calculate area with the shorter line and move left pointer right
			maxArea = max(maxArea, width*height[l])
			l++
		} else {
			// Calculate area with the shorter line and move right pointer left
			maxArea = max(maxArea, width*height[r])
			r--
		}
	}
	log.Printf("maxArea %v", maxArea)
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
